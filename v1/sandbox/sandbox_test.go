package sandbox

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/coder/websocket"
)

type fakeSandbox struct {
	server *httptest.Server

	mu    sync.Mutex
	files map[string][]byte

	connections   atomic.Int32
	tokenRequests atomic.Int32
	flakyDrops    atomic.Int32
	moveRequests  atomic.Int32
}

func newFakeSandbox(t *testing.T) *fakeSandbox {
	fake := &fakeSandbox{files: map[string][]byte{}}

	mux := http.NewServeMux()
	mux.HandleFunc("/files", fake.socket)
	mux.HandleFunc("/exec", fake.socket)
	mux.HandleFunc("/app/", fake.http)

	fake.server = httptest.NewServer(mux)
	t.Cleanup(fake.server.Close)
	return fake
}

func (f *fakeSandbox) access(path string, tokens ...string) *accessCache {
	if len(tokens) == 0 {
		tokens = []string{"good"}
	}
	var mu sync.Mutex

	return newAccessCache(func(context.Context) (*Grant, error) {
		f.tokenRequests.Add(1)

		mu.Lock()
		token := tokens[0]
		if len(tokens) > 1 {
			tokens = tokens[1:]
		}
		mu.Unlock()

		url := strings.Replace(f.server.URL, "http", "ws", 1) + path
		if path == "/port" {
			url = f.server.URL + "/app"
		}
		return &Grant{URL: url, Token: token, ExpiresAt: time.Now().Add(time.Hour)}, nil
	})
}

func (f *fakeSandbox) socket(w http.ResponseWriter, r *http.Request) {
	token := ""
	for _, protocol := range strings.Split(r.Header.Get("Sec-WebSocket-Protocol"), ",") {
		protocol = strings.TrimSpace(protocol)
		if strings.HasPrefix(protocol, "metorial-sandbox-token.") {
			token = strings.TrimPrefix(protocol, "metorial-sandbox-token.")
		}
	}
	if token != "good" {
		http.Error(w, "token_invalid", http.StatusUnauthorized)
		return
	}

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{Subprotocols: []string{"metorial-sandbox-v1"}})
	if err != nil {
		return
	}
	conn.SetReadLimit(maxMessageSize)
	f.connections.Add(1)

	if r.URL.Path == "/files" {
		f.serveFiles(r.Context(), conn)
	} else {
		f.serveExec(r.Context(), conn)
	}
}

func sendJSON(ctx context.Context, conn *websocket.Conn, value map[string]any) {
	data, _ := json.Marshal(value)
	_ = conn.Write(ctx, websocket.MessageText, data)
}

func (f *fakeSandbox) serveFiles(ctx context.Context, conn *websocket.Conn) {
	uploads := map[string]*struct {
		path string
		data bytes.Buffer
	}{}

	for {
		messageType, data, err := conn.Read(ctx)
		if err != nil {
			return
		}

		if messageType == websocket.MessageBinary {
			id, _, payload, _ := decodeFrame(data, false)
			if upload := uploads[id]; upload != nil {
				upload.data.Write(payload)
			}
			continue
		}

		var message map[string]any
		_ = json.Unmarshal(data, &message)
		id, _ := message["id"].(string)
		reply := func(fields map[string]any) {
			fields["id"] = id
			sendJSON(ctx, conn, fields)
		}

		if message["type"] == "end" {
			upload := uploads[id]
			if upload == nil {
				continue
			}
			delete(uploads, id)
			sum := sha256.Sum256(upload.data.Bytes())
			if message["sha256"] != hex.EncodeToString(sum[:]) {
				reply(map[string]any{"type": "error", "error": map[string]any{"code": "checksum_mismatch"}})
				continue
			}
			f.mu.Lock()
			f.files[upload.path] = append([]byte(nil), upload.data.Bytes()...)
			f.mu.Unlock()
			reply(map[string]any{"type": "result", "size": upload.data.Len()})
			continue
		}

		path, _ := message["path"].(string)
		f.mu.Lock()
		content, exists := f.files[path]
		f.mu.Unlock()

		switch message["op"] {
		case "write":
			if overwrite, _ := message["overwrite"].(bool); !overwrite && exists {
				reply(map[string]any{"type": "error", "error": map[string]any{"code": "already_exists", "path": path}})
				continue
			}
			uploads[id] = &struct {
				path string
				data bytes.Buffer
			}{path: path}
		case "read":
			if !exists {
				reply(map[string]any{"type": "error", "error": map[string]any{"code": "not_found", "path": path}})
				continue
			}
			reply(map[string]any{"type": "start", "size": len(content), "mtime": 0, "offset": 0})
			for start := 0; start < len(content); start += 5 {
				end := min(start+5, len(content))
				_ = conn.Write(ctx, websocket.MessageBinary, encodeFrame(id, content[start:end]))
			}
			reply(map[string]any{"type": "result"})
		case "stat":
			if !exists {
				reply(map[string]any{"type": "error", "error": map[string]any{"code": "not_found"}})
				continue
			}
			reply(map[string]any{"type": "result", "entry": map[string]any{
				"path": path, "type": "file", "size": len(content), "mode": 420, "mtime": 1000,
			}})
		case "list":
			if path == "/flaky" && f.flakyDrops.Load() > 0 {
				f.flakyDrops.Add(-1)
				reply(map[string]any{"type": "entries", "entries": []any{map[string]any{"name": "partial"}}})
				conn.CloseNow()
				return
			}
			var entries []any
			f.mu.Lock()
			for name, data := range f.files {
				if strings.HasPrefix(name, path+"/") {
					entries = append(entries, map[string]any{"name": name[len(path)+1:], "type": "file", "size": len(data)})
				}
			}
			f.mu.Unlock()
			split := min(1, len(entries))
			reply(map[string]any{"type": "entries", "entries": entries[:split]})
			reply(map[string]any{"type": "entries", "entries": entries[split:]})
			reply(map[string]any{"type": "result", "truncated": false})
		case "move":
			f.moveRequests.Add(1)
			conn.CloseNow()
			return
		case "delete":
			f.mu.Lock()
			delete(f.files, path)
			f.mu.Unlock()
			reply(map[string]any{"type": "result"})
		}
	}
}

func (f *fakeSandbox) serveExec(ctx context.Context, conn *websocket.Conn) {
	stdin := map[string]*bytes.Buffer{}

	output := func(id string, stream byte, text string) {
		_ = conn.Write(ctx, websocket.MessageBinary, encodeFrame(id, []byte(text), stream))
	}

	for {
		messageType, data, err := conn.Read(ctx)
		if err != nil {
			return
		}

		if messageType == websocket.MessageBinary {
			id, _, payload, _ := decodeFrame(data, false)
			if buffer := stdin[id]; buffer != nil {
				buffer.Write(payload)
			}
			continue
		}

		var message map[string]any
		_ = json.Unmarshal(data, &message)
		id, _ := message["id"].(string)
		send := func(fields map[string]any) {
			fields["id"] = id
			sendJSON(ctx, conn, fields)
		}

		switch message["type"] {
		case "stdin_eof":
			output(id, 1, strings.ToUpper(stdin[id].String()))
			delete(stdin, id)
			send(map[string]any{"type": "exit", "code": 0, "signal": nil, "timedOut": false})
			continue
		case "signal":
			send(map[string]any{"type": "exit", "code": nil, "signal": message["signal"], "timedOut": false})
			continue
		}

		command, _ := message["command"].([]any)
		program, _ := command[0].(string)
		var args []string
		for _, arg := range command[1:] {
			args = append(args, arg.(string))
		}

		send(map[string]any{"type": "started", "pid": 42})

		switch {
		case program == "missing":
			send(map[string]any{"type": "error", "error": map[string]any{"code": "spawn_failed", "message": "not found"}})
		case program == "drop":
			conn.CloseNow()
			return
		case message["stdin"] == true || program == "sleep":
			stdin[id] = &bytes.Buffer{}
		default:
			cwd, _ := message["cwd"].(string)
			output(id, 1, "ran "+strings.Join(args, " "))
			output(id, 2, "cwd="+cwd)
			send(map[string]any{"type": "exit", "code": 3, "signal": nil, "timedOut": false})
		}
	}
}

func (f *fakeSandbox) http(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(tokenHeader) != "good" {
		http.Error(w, "token_invalid", http.StatusUnauthorized)
		return
	}

	body, _ := io.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"method": r.Method,
		"url":    r.URL.RequestURI(),
		"body":   string(body),
		"test":   r.Header.Get("X-Test"),
	})
}

func testContext(t *testing.T) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	t.Cleanup(cancel)
	return ctx
}

func must[T any](value T, err error) func(t *testing.T) T {
	return func(t *testing.T) T {
		t.Helper()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		return value
	}
}

func TestFilesRoundTrip(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	files := must(connectFiles(ctx, fake.access("/files"), nil))(t)
	defer files.Close()

	big := bytes.Repeat([]byte("x"), 200*1024)
	must(files.Write(ctx, "/work/a.txt", []byte("hello world"), nil))(t)
	must(files.Write(ctx, "/work/b.bin", big, nil))(t)

	if text := must(files.ReadText(ctx, "/work/a.txt", nil))(t); text != "hello world" {
		t.Fatalf("read %q", text)
	}
	if data := must(files.Read(ctx, "/work/b.bin", nil))(t); !bytes.Equal(data, big) {
		t.Fatalf("read %d bytes", len(data))
	}

	entry := must(files.Stat(ctx, "/work/a.txt", nil))(t)
	if entry.Name != "a.txt" || entry.Path != "/work/a.txt" || entry.Size != 11 || entry.ModifiedAt.UnixMilli() != 1000 {
		t.Fatalf("stat %+v", entry)
	}

	list := must(files.List(ctx, "/work", nil))(t)
	var paths []string
	for _, e := range list.Entries {
		paths = append(paths, e.Path)
	}
	sort.Strings(paths)
	if strings.Join(paths, ",") != "/work/a.txt,/work/b.bin" || list.Truncated {
		t.Fatalf("list %v", paths)
	}

	if err := files.Delete(ctx, "/work/a.txt", nil); err != nil {
		t.Fatal(err)
	}
	if must(files.Exists(ctx, "/work/a.txt"))(t) {
		t.Fatal("file still exists")
	}
}

func TestFilesErrors(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	files := must(connectFiles(ctx, fake.access("/files"), nil))(t)
	defer files.Close()

	_, err := files.Read(ctx, "/nope", nil)
	var accessErr *AccessError
	if !errors.As(err, &accessErr) || accessErr.Code != "not_found" || accessErr.Path != "/nope" {
		t.Fatalf("got %v", err)
	}

	must(files.Write(ctx, "/exists", []byte("a"), nil))(t)
	if _, err := files.Write(ctx, "/exists", []byte("b"), &WriteOptions{KeepExisting: true}); ErrorCode(err) != "already_exists" {
		t.Fatalf("got %v", err)
	}
}

func TestFilesRefreshesRejectedToken(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	files := must(connectFiles(ctx, fake.access("/files", "stale", "good"), nil))(t)
	defer files.Close()

	if fake.tokenRequests.Load() != 2 {
		t.Fatalf("token requests %d", fake.tokenRequests.Load())
	}
	must(files.Write(ctx, "/refresh", []byte("ok"), nil))(t)
}

func TestFilesResendsPendingList(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	files := must(connectFiles(ctx, fake.access("/files"), nil))(t)
	defer files.Close()

	must(files.Write(ctx, "/flaky/one", []byte("1"), nil))(t)
	fake.flakyDrops.Store(2)

	list := must(files.List(ctx, "/flaky", nil))(t)
	if fake.flakyDrops.Load() != 0 || len(list.Entries) != 1 || list.Entries[0].Path != "/flaky/one" {
		t.Fatalf("list %+v drops %d", list, fake.flakyDrops.Load())
	}
}

func TestFilesDoesNotResendMove(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	files := must(connectFiles(ctx, fake.access("/files"), nil))(t)
	defer files.Close()

	if err := files.Move(ctx, "/a", "/b", nil); ErrorCode(err) != "disconnected" {
		t.Fatalf("got %v", err)
	}
	if fake.moveRequests.Load() != 1 {
		t.Fatalf("move requests %d", fake.moveRequests.Load())
	}
}

func TestFilesClosed(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	files := must(connectFiles(ctx, fake.access("/files"), nil))(t)
	files.Close()

	if _, err := files.Stat(ctx, "/x", nil); ErrorCode(err) != "closed" {
		t.Fatalf("got %v", err)
	}
}

func TestShellCommands(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	shell := must(connectShell(ctx, fake.access("/exec"), nil))(t)
	defer shell.Close()

	result := must(shell.Exec(ctx, []string{"echo", "hi", "there"}, &CommandOptions{Cwd: "/tmp"}))(t)
	if *result.ExitCode != 3 || result.StdoutString() != "ran hi there" || result.StderrString() != "cwd=/tmp" {
		t.Fatalf("result %+v", result)
	}

	viaShell := must(shell.Command(ctx, "ls -la", nil))(t)
	if viaShell.StdoutString() != "ran -c ls -la" {
		t.Fatalf("stdout %q", viaShell.StdoutString())
	}

	piped := must(shell.Exec(ctx, []string{"cat"}, &CommandOptions{Stdin: []byte("shout")}))(t)
	if piped.StdoutString() != "SHOUT" {
		t.Fatalf("stdout %q", piped.StdoutString())
	}
}

func TestShellSpawnSignal(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	shell := must(connectShell(ctx, fake.access("/exec"), nil))(t)
	defer shell.Close()

	process := must(shell.Spawn(ctx, []string{"sleep", "100"}, nil))(t)
	if pid := must(process.Started(ctx))(t); pid != 42 {
		t.Fatalf("pid %d", pid)
	}

	if err := process.Signal("SIGTERM"); err != nil {
		t.Fatal(err)
	}
	exit := must(process.Wait(ctx))(t)
	if exit.ExitCode != nil || exit.Signal != "TERM" {
		t.Fatalf("exit %+v", exit)
	}
}

func TestShellErrorsAndReconnect(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	shell := must(connectShell(ctx, fake.access("/exec"), nil))(t)
	defer shell.Close()

	if _, err := shell.Exec(ctx, []string{"missing"}, nil); ErrorCode(err) != "spawn_failed" {
		t.Fatalf("got %v", err)
	}

	before := fake.connections.Load()
	_, err := shell.Exec(ctx, []string{"drop"}, nil)
	var commandErr *CommandError
	if !errors.As(err, &commandErr) || commandErr.Code != "connection_interrupted" {
		t.Fatalf("got %v", err)
	}

	time.Sleep(600 * time.Millisecond)
	if delta := fake.connections.Load() - before; delta != 1 {
		t.Fatalf("reconnections %d", delta)
	}

	result := must(shell.Exec(ctx, []string{"echo", "again"}, nil))(t)
	if result.StdoutString() != "ran again" {
		t.Fatalf("stdout %q", result.StdoutString())
	}
}

func TestPortRequests(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	port := must(connectPort(ctx, 3000, fake.access("/port"), nil))(t)

	resp := must(port.Get(ctx, "/testing?x=1", http.Header{"X-Test": {"yes"}}))(t)
	var body map[string]any
	_ = json.NewDecoder(resp.Body).Decode(&body)
	resp.Body.Close()
	if body["method"] != "GET" || body["url"] != "/app/testing?x=1" || body["test"] != "yes" {
		t.Fatalf("body %v", body)
	}

	posted := must(port.HTTPClient().Post("/items", "application/json", strings.NewReader(`{"name":"a"}`)))(t)
	_ = json.NewDecoder(posted.Body).Decode(&body)
	posted.Body.Close()
	if body["body"] != `{"name":"a"}` {
		t.Fatalf("body %v", body)
	}

	if _, err := port.Get(ctx, "https://example.com/", nil); err == nil {
		t.Fatal("expected cross-origin error")
	}
	if port.URL() != fake.server.URL+"/app" {
		t.Fatalf("url %s", port.URL())
	}
}

func TestPortRetries401(t *testing.T) {
	fake, ctx := newFakeSandbox(t), testContext(t)
	port := must(connectPort(ctx, 3000, fake.access("/port", "stale", "good"), nil))(t)

	resp := must(port.Post(ctx, "/", "text/plain", strings.NewReader("hello")))(t)
	var body map[string]any
	_ = json.NewDecoder(resp.Body).Decode(&body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK || body["body"] != "hello" {
		t.Fatalf("status %d body %v", resp.StatusCode, body)
	}

	port.Close()
	if _, err := port.Get(ctx, "/", nil); ErrorCode(err) != "closed" {
		t.Fatalf("got %v", err)
	}
}
