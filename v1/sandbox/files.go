package sandbox

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/coder/websocket"
)

const (
	fileChunkSize    = 64 * 1024
	defaultListLimit = 10000
	maxFileAttempts  = 3
	resendDelay      = 200 * time.Millisecond
)

// FileEntry describes a file, directory, or symlink inside the machine.
type FileEntry struct {
	Name       string
	Path       string
	Type       string
	Size       int64
	Mode       uint32
	ModifiedAt time.Time
	Target     string
}

// FileList is the result of Files.List. Truncated is set when the listing hit its limit.
type FileList struct {
	Entries   []FileEntry
	Truncated bool
}

type ListOptions struct {
	Recursive bool
	MaxDepth  int
	Limit     int
}

type StatOptions struct {
	FollowSymlinks bool
}

type ReadOptions struct {
	Offset int64
	Length int64
}

type WriteOptions struct {
	// KeepExisting fails the write with "already_exists" instead of replacing the file.
	KeepExisting bool
	Parents      bool
	// Mode is the file's permission bits; zero keeps the server default (0644).
	Mode uint32
}

type MkdirOptions struct {
	Parents bool
	Mode    uint32
}

type MoveOptions struct {
	Overwrite bool
}

type DeleteOptions struct {
	Recursive bool
	MissingOK bool
}

type fileResult struct {
	fields map[string]any
	err    *AccessError
}

type pendingFileOp struct {
	onEntries func(entries []any)
	onData    func(data []byte)
	done      chan fileResult
	once      sync.Once
}

func (p *pendingFileOp) settle(result fileResult) {
	p.once.Do(func() { p.done <- result })
}

// Files is a connection to a machine's file service. It is safe for concurrent use.
type Files struct {
	socket *muxSocket

	mu      sync.Mutex
	pending map[string]*pendingFileOp
}

func connectFiles(ctx context.Context, access *accessCache, onNotice func(Notice)) (*Files, error) {
	files := &Files{pending: map[string]*pendingFileOp{}}
	files.socket = newMuxSocket(access, muxHandlers{
		onText:       files.handleText,
		onBinary:     files.handleBinary,
		onDisconnect: files.handleDisconnect,
		onNotice:     onNotice,
	})

	if err := files.socket.ensureOpen(ctx); err != nil {
		files.socket.close()
		return nil, err
	}
	return files, nil
}

func (f *Files) List(ctx context.Context, path string, opts *ListOptions) (*FileList, error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	return withFileRetry(ctx, f, true, func() (*FileList, error) {
		var mu sync.Mutex
		list := &FileList{Entries: []FileEntry{}}

		fields := map[string]any{"path": path, "limit": defaultListLimit}
		if opts.Limit > 0 {
			fields["limit"] = opts.Limit
		}
		if opts.Recursive {
			fields["recursive"] = true
		}
		if opts.MaxDepth > 0 {
			fields["maxDepth"] = opts.MaxDepth
		}

		result, err := f.run(ctx, "list", fields, &pendingFileOp{onEntries: func(entries []any) {
			mu.Lock()
			defer mu.Unlock()
			for _, raw := range entries {
				list.Entries = append(list.Entries, toFileEntry(raw, path))
			}
		}}, nil)
		if err != nil {
			return nil, err
		}

		list.Truncated, _ = result["truncated"].(bool)
		return list, nil
	})
}

func (f *Files) Stat(ctx context.Context, path string, opts *StatOptions) (*FileEntry, error) {
	follow := opts != nil && opts.FollowSymlinks

	return withFileRetry(ctx, f, true, func() (*FileEntry, error) {
		result, err := f.run(ctx, "stat", map[string]any{"path": path, "followSymlinks": follow}, nil, nil)
		if err != nil {
			return nil, err
		}
		entry := toFileEntry(result["entry"], "")
		return &entry, nil
	})
}

func (f *Files) Exists(ctx context.Context, path string) (bool, error) {
	_, err := f.Stat(ctx, path, nil)
	if err == nil {
		return true, nil
	}
	if ErrorCode(err) == "not_found" {
		return false, nil
	}
	return false, err
}

func (f *Files) Read(ctx context.Context, path string, opts *ReadOptions) ([]byte, error) {
	return withFileRetry(ctx, f, true, func() ([]byte, error) {
		var buf bytes.Buffer
		var mu sync.Mutex

		fields := map[string]any{"path": path}
		if opts != nil && opts.Offset > 0 {
			fields["offset"] = opts.Offset
		}
		if opts != nil && opts.Length > 0 {
			fields["length"] = opts.Length
		}

		_, err := f.run(ctx, "read", fields, &pendingFileOp{onData: func(data []byte) {
			mu.Lock()
			buf.Write(data)
			mu.Unlock()
		}}, nil)
		if err != nil {
			return nil, err
		}

		mu.Lock()
		defer mu.Unlock()
		return buf.Bytes(), nil
	})
}

func (f *Files) ReadText(ctx context.Context, path string, opts *ReadOptions) (string, error) {
	data, err := f.Read(ctx, path, opts)
	return string(data), err
}

// Write replaces the file at path with data and returns the number of bytes written.
func (f *Files) Write(ctx context.Context, path string, data []byte, opts *WriteOptions) (int64, error) {
	if opts == nil {
		opts = &WriteOptions{}
	}

	sum := sha256.Sum256(data)
	checksum := hex.EncodeToString(sum[:])

	fields := map[string]any{"path": path, "size": len(data), "overwrite": !opts.KeepExisting}
	if opts.Parents {
		fields["parents"] = true
	}
	if opts.Mode != 0 {
		fields["mode"] = opts.Mode
	}

	upload := func(id string, active func() bool) {
		for start := 0; start < len(data); start += fileChunkSize {
			if !active() {
				return
			}
			end := min(start+fileChunkSize, len(data))
			f.socket.send(ctx, websocket.MessageBinary, encodeFrame(id, data[start:end]))
		}
		if active() {
			f.socket.sendJSON(ctx, map[string]any{"id": id, "type": "end", "sha256": checksum})
		}
	}

	return withFileRetry(ctx, f, !opts.KeepExisting, func() (int64, error) {
		result, err := f.run(ctx, "write", fields, nil, upload)
		if err != nil {
			return 0, err
		}
		if size, ok := result["size"].(float64); ok {
			return int64(size), nil
		}
		return int64(len(data)), nil
	})
}

func (f *Files) Mkdir(ctx context.Context, path string, opts *MkdirOptions) error {
	if opts == nil {
		opts = &MkdirOptions{}
	}

	fields := map[string]any{"path": path, "parents": opts.Parents}
	if opts.Mode != 0 {
		fields["mode"] = opts.Mode
	}

	_, err := withFileRetry(ctx, f, opts.Parents, func() (map[string]any, error) {
		return f.run(ctx, "mkdir", fields, nil, nil)
	})
	return err
}

func (f *Files) Move(ctx context.Context, from, to string, opts *MoveOptions) error {
	overwrite := opts != nil && opts.Overwrite

	_, err := withFileRetry(ctx, f, false, func() (map[string]any, error) {
		return f.run(ctx, "move", map[string]any{"from": from, "to": to, "overwrite": overwrite}, nil, nil)
	})
	return err
}

func (f *Files) Delete(ctx context.Context, path string, opts *DeleteOptions) error {
	if opts == nil {
		opts = &DeleteOptions{}
	}

	fields := map[string]any{"path": path, "recursive": opts.Recursive}
	if opts.MissingOK {
		fields["missingOk"] = true
	}

	_, err := withFileRetry(ctx, f, opts.MissingOK, func() (map[string]any, error) {
		return f.run(ctx, "delete", fields, nil, nil)
	})
	return err
}

// Close ends the connection. Pending and later calls fail with code "closed".
func (f *Files) Close() error {
	f.socket.close()
	return nil
}

func (f *Files) take(id string) *pendingFileOp {
	f.mu.Lock()
	defer f.mu.Unlock()

	op := f.pending[id]
	delete(f.pending, id)
	return op
}

func (f *Files) lookup(id string) *pendingFileOp {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.pending[id]
}

func (f *Files) handleText(message map[string]any) {
	id, _ := message["id"].(string)
	if id == "" {
		return
	}

	switch message["type"] {
	case "entries":
		if op := f.lookup(id); op != nil && op.onEntries != nil {
			entries, _ := message["entries"].([]any)
			op.onEntries(entries)
		}
	case "result":
		if op := f.take(id); op != nil {
			delete(message, "id")
			delete(message, "type")
			op.settle(fileResult{fields: message})
		}
	case "error":
		if op := f.take(id); op != nil {
			op.settle(fileResult{err: serverError(message["error"])})
		}
	}
}

func (f *Files) handleBinary(data []byte) {
	id, _, payload, ok := decodeFrame(data, false)
	if !ok {
		return
	}
	if op := f.lookup(id); op != nil && op.onData != nil {
		op.onData(append([]byte(nil), payload...))
	}
}

func (f *Files) handleDisconnect(err *AccessError) {
	f.mu.Lock()
	pending := f.pending
	f.pending = map[string]*pendingFileOp{}
	f.mu.Unlock()

	for _, op := range pending {
		op.settle(fileResult{err: err})
	}
}

func (f *Files) run(
	ctx context.Context,
	op string,
	fields map[string]any,
	hooks *pendingFileOp,
	upload func(id string, active func() bool),
) (map[string]any, error) {
	if err := f.socket.ensureOpen(ctx); err != nil {
		var accessErr *AccessError
		if errors.As(err, &accessErr) {
			accessErr.unsent = true
		}
		return nil, err
	}

	if hooks == nil {
		hooks = &pendingFileOp{}
	}
	pending := &pendingFileOp{onEntries: hooks.onEntries, onData: hooks.onData, done: make(chan fileResult, 1)}

	id := newRequestID()
	f.mu.Lock()
	f.pending[id] = pending
	f.mu.Unlock()

	request := map[string]any{"id": id, "op": op}
	for key, value := range fields {
		request[key] = value
	}

	if !f.socket.sendJSON(ctx, request) {
		f.take(id)
		lost := newError("disconnected", "The file connection was lost.", true)
		lost.unsent = true
		return nil, lost
	}

	finished := make(chan struct{})
	defer close(finished)

	if upload != nil {
		go upload(id, func() bool {
			select {
			case <-finished:
				return false
			default:
				return true
			}
		})
	}

	select {
	case result := <-pending.done:
		if result.err != nil {
			return nil, result.err
		}
		return result.fields, nil
	case <-ctx.Done():
		f.take(id)
		f.socket.sendJSON(context.Background(), map[string]any{"id": newRequestID(), "op": "cancel", "target": id})
		return nil, ctx.Err()
	}
}

func withFileRetry[T any](ctx context.Context, f *Files, canResend bool, run func() (T, error)) (T, error) {
	for attempt := 1; ; attempt++ {
		value, err := run()
		if err == nil {
			return value, nil
		}

		var accessErr *AccessError
		resendable := errors.As(err, &accessErr) &&
			accessErr.Retryable &&
			(canResend || accessErr.unsent) &&
			attempt < maxFileAttempts &&
			!f.socket.isClosed()
		if !resendable {
			return value, err
		}

		select {
		case <-time.After(resendDelay * time.Duration(attempt)):
		case <-ctx.Done():
			return value, ctx.Err()
		}
	}
}

func toFileEntry(raw any, parent string) FileEntry {
	data, _ := raw.(map[string]any)
	name, _ := data["name"].(string)
	rawPath, _ := data["path"].(string)
	target, _ := data["target"].(string)
	fileType, _ := data["type"].(string)
	size, _ := data["size"].(float64)
	mode, _ := data["mode"].(float64)
	mtime, _ := data["mtime"].(float64)

	path := rawPath
	if parent != "" {
		child := rawPath
		if name != "" {
			child = name
		}
		path = strings.TrimRight(parent, "/") + "/" + strings.TrimLeft(child, "/")
	}
	if name == "" {
		trimmed := strings.TrimRight(path, "/")
		name = trimmed[strings.LastIndex(trimmed, "/")+1:]
		if name == "" {
			name = "/"
		}
	}

	return FileEntry{
		Name:       name,
		Path:       path,
		Type:       fileType,
		Size:       int64(size),
		Mode:       uint32(mode),
		ModifiedAt: time.UnixMilli(int64(mtime)),
		Target:     target,
	}
}
