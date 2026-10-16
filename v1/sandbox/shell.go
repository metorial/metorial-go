package sandbox

import (
	"bytes"
	"context"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/coder/websocket"
)

const streamStderr = 2

// SpawnOptions configure a streaming process. Stdout and Stderr are written to from
// the connection's reader goroutine, so they should not block.
type SpawnOptions struct {
	Cwd     string
	Env     map[string]string
	Timeout time.Duration
	Stdin   bool
	Stdout  io.Writer
	Stderr  io.Writer
}

type CommandOptions struct {
	Cwd     string
	Env     map[string]string
	Timeout time.Duration
	Stdin   []byte
	Stdout  io.Writer
	Stderr  io.Writer
}

type ProcessExit struct {
	// ExitCode is nil when the process was ended by a signal.
	ExitCode *int
	Signal   string
	TimedOut bool
}

type CommandResult struct {
	ProcessExit
	Stdout []byte
	Stderr []byte
}

func (r *CommandResult) StdoutString() string { return string(r.Stdout) }
func (r *CommandResult) StderrString() string { return string(r.Stderr) }

// Process is a command running in the machine.
type Process struct {
	ID string

	socket *muxSocket
	stdout io.Writer
	stderr io.Writer

	started  chan struct{}
	finished chan struct{}

	mu   sync.Mutex
	pid  int
	exit *ProcessExit
	err  error
}

func (p *Process) PID() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pid
}

// Started waits until the machine reports the process id.
func (p *Process) Started(ctx context.Context) (int, error) {
	select {
	case <-p.started:
	case <-p.finished:
	case <-ctx.Done():
		return 0, ctx.Err()
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if p.pid == 0 && p.err != nil {
		return 0, p.err
	}
	return p.pid, nil
}

func (p *Process) Wait(ctx context.Context) (*ProcessExit, error) {
	select {
	case <-p.finished:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	return p.exit, p.err
}

// Write sends data to the process's stdin. The process must be spawned with Stdin set.
func (p *Process) Write(data []byte) (int, error) {
	if err := p.ensureRunning(); err != nil {
		return 0, err
	}
	if !p.socket.send(context.Background(), websocket.MessageBinary, encodeFrame(p.ID, data)) {
		return 0, lostCommandError()
	}
	return len(data), nil
}

func (p *Process) CloseStdin() error {
	return p.control(map[string]any{"type": "stdin_eof"})
}

// Signal sends a signal such as "TERM", "INT", or "SIGKILL" to the process.
func (p *Process) Signal(signal string) error {
	name := strings.TrimPrefix(strings.ToUpper(signal), "SIG")
	return p.control(map[string]any{"type": "signal", "signal": name})
}

func (p *Process) Kill() error {
	return p.Signal("KILL")
}

func (p *Process) control(fields map[string]any) error {
	if err := p.ensureRunning(); err != nil {
		return err
	}
	fields["id"] = p.ID
	if !p.socket.sendJSON(context.Background(), fields) {
		return lostCommandError()
	}
	return nil
}

func (p *Process) ensureRunning() error {
	select {
	case <-p.finished:
		return newCommandError("unknown_command", "The command has already finished.")
	default:
		return nil
	}
}

func (p *Process) handleOutput(stream byte, data []byte) {
	writer := p.stdout
	if stream == streamStderr {
		writer = p.stderr
	}
	if writer != nil {
		_, _ = writer.Write(data)
	}
}

func (p *Process) handleStarted(pid int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case <-p.started:
	default:
		p.pid = pid
		close(p.started)
	}
}

func (p *Process) finish(exit *ProcessExit, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case <-p.finished:
	default:
		p.exit, p.err = exit, err
		close(p.finished)
	}
}

func lostCommandError() *CommandError {
	return newCommandError("connection_interrupted", "The shell connection was lost.")
}

// Shell is a connection for running commands in a machine. Many commands can run
// concurrently over one connection. It is safe for concurrent use.
type Shell struct {
	socket *muxSocket

	mu        sync.Mutex
	processes map[string]*Process
}

func connectShell(ctx context.Context, access *accessCache, onNotice func(Notice)) (*Shell, error) {
	shell := &Shell{processes: map[string]*Process{}}
	shell.socket = newMuxSocket(access, muxHandlers{
		onText:       shell.handleText,
		onBinary:     shell.handleBinary,
		onDisconnect: shell.handleDisconnect,
		onNotice:     onNotice,
	})

	if err := shell.socket.ensureOpen(ctx); err != nil {
		shell.socket.close()
		return nil, err
	}
	return shell, nil
}

// Spawn starts argv and returns immediately with a handle for streaming its I/O.
func (s *Shell) Spawn(ctx context.Context, argv []string, opts *SpawnOptions) (*Process, error) {
	if opts == nil {
		opts = &SpawnOptions{}
	}
	if err := s.socket.ensureOpen(ctx); err != nil {
		return nil, err
	}

	process := &Process{
		ID:       newRequestID(),
		socket:   s.socket,
		stdout:   opts.Stdout,
		stderr:   opts.Stderr,
		started:  make(chan struct{}),
		finished: make(chan struct{}),
	}

	s.mu.Lock()
	s.processes[process.ID] = process
	s.mu.Unlock()

	request := map[string]any{"id": process.ID, "op": "run", "v": 1, "command": argv}
	if opts.Cwd != "" {
		request["cwd"] = opts.Cwd
	}
	if opts.Env != nil {
		request["env"] = opts.Env
	}
	if opts.Stdin {
		request["stdin"] = true
	}
	if opts.Timeout > 0 {
		request["timeoutMs"] = opts.Timeout.Milliseconds()
	}

	if !s.socket.sendJSON(ctx, request) {
		s.take(process.ID)
		return nil, lostCommandError()
	}
	return process, nil
}

// Exec runs argv to completion and returns its buffered output.
func (s *Shell) Exec(ctx context.Context, argv []string, opts *CommandOptions) (*CommandResult, error) {
	if opts == nil {
		opts = &CommandOptions{}
	}

	var stdout, stderr bytes.Buffer
	process, err := s.Spawn(ctx, argv, &SpawnOptions{
		Cwd:     opts.Cwd,
		Env:     opts.Env,
		Timeout: opts.Timeout,
		Stdin:   opts.Stdin != nil,
		Stdout:  teeWriter(&stdout, opts.Stdout),
		Stderr:  teeWriter(&stderr, opts.Stderr),
	})
	if err != nil {
		return nil, err
	}

	if opts.Stdin != nil {
		if len(opts.Stdin) > 0 {
			if _, err := process.Write(opts.Stdin); err != nil {
				return nil, err
			}
		}
		if err := process.CloseStdin(); err != nil {
			return nil, err
		}
	}

	exit, err := process.Wait(ctx)
	if err != nil {
		if ctx.Err() != nil {
			_ = process.Kill()
		}
		return nil, err
	}

	return &CommandResult{ProcessExit: *exit, Stdout: stdout.Bytes(), Stderr: stderr.Bytes()}, nil
}

// Command runs script with /bin/sh -c and returns its buffered output.
func (s *Shell) Command(ctx context.Context, script string, opts *CommandOptions) (*CommandResult, error) {
	return s.Exec(ctx, []string{"/bin/sh", "-c", script}, opts)
}

// Close ends the connection. Running commands fail with code "closed".
func (s *Shell) Close() error {
	s.socket.close()
	return nil
}

func (s *Shell) take(id string) *Process {
	s.mu.Lock()
	defer s.mu.Unlock()

	process := s.processes[id]
	delete(s.processes, id)
	return process
}

func (s *Shell) lookup(id string) *Process {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.processes[id]
}

func (s *Shell) handleText(message map[string]any) {
	id, _ := message["id"].(string)
	if id == "" {
		return
	}

	switch message["type"] {
	case "started":
		if process := s.lookup(id); process != nil {
			pid, _ := message["pid"].(float64)
			process.handleStarted(int(pid))
		}
	case "exit":
		if process := s.take(id); process != nil {
			exit := &ProcessExit{}
			if code, ok := message["code"].(float64); ok {
				value := int(code)
				exit.ExitCode = &value
			}
			exit.Signal, _ = message["signal"].(string)
			exit.TimedOut, _ = message["timedOut"].(bool)
			process.finish(exit, nil)
		}
	case "error":
		if process := s.take(id); process != nil {
			process.finish(nil, &CommandError{AccessError: *serverError(message["error"])})
		}
	}
}

func (s *Shell) handleBinary(data []byte) {
	id, tag, payload, ok := decodeFrame(data, true)
	if !ok {
		return
	}
	if process := s.lookup(id); process != nil {
		process.handleOutput(tag, append([]byte(nil), payload...))
	}
}

func (s *Shell) handleDisconnect(err *AccessError) {
	s.mu.Lock()
	processes := s.processes
	s.processes = map[string]*Process{}
	s.mu.Unlock()

	var reason error = newCommandError("connection_interrupted", err.Message)
	if err.Code == "closed" {
		reason = err
	}

	for _, process := range processes {
		process.finish(nil, reason)
	}
}

func teeWriter(buffer *bytes.Buffer, extra io.Writer) io.Writer {
	if extra == nil {
		return buffer
	}
	return io.MultiWriter(buffer, extra)
}
