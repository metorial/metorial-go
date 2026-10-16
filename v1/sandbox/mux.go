package sandbox

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
)

const (
	maxConnectAttempts = 3
	connectBaseDelay   = 300 * time.Millisecond
	maxReconnectDelay  = 10 * time.Second
	maxMessageSize     = 16 * 1024 * 1024
	connectTimeout     = 30 * time.Second
)

var reconnectCodes = map[string]bool{
	"disconnected":        true,
	"unauthorized":        true,
	"rate_limited":        true,
	"machine_restarting":  true,
	"machine_unavailable": true,
}

// Notice is an informational message the sandbox proxy sends on an open connection,
// such as a warning that the connection is about to close for being idle.
type Notice struct {
	Notice        string
	Reason        string
	IdleInSeconds float64
}

var (
	idCounter atomic.Uint64
	idPrefix  = func() string {
		buf := make([]byte, 3)
		_, _ = rand.Read(buf)
		return hex.EncodeToString(buf)
	}()
)

func newRequestID() string {
	return fmt.Sprintf("%s_%x", idPrefix, idCounter.Add(1))
}

func socketProtocols(token string) []string {
	return []string{"metorial-sandbox-v1", "metorial-sandbox-token." + token}
}

func encodeFrame(id string, payload []byte, tag ...byte) []byte {
	frame := make([]byte, 0, 1+len(id)+len(tag)+len(payload))
	frame = append(frame, byte(len(id)))
	frame = append(frame, id...)
	frame = append(frame, tag...)
	return append(frame, payload...)
}

func decodeFrame(data []byte, hasTag bool) (id string, tag byte, payload []byte, ok bool) {
	if len(data) == 0 {
		return "", 0, nil, false
	}

	offset := 1 + int(data[0])
	if hasTag {
		offset++
	}
	if len(data) < offset {
		return "", 0, nil, false
	}

	id = string(data[1 : 1+int(data[0])])
	if hasTag {
		tag = data[offset-1]
	}
	return id, tag, data[offset:], true
}

type muxHandlers struct {
	onText       func(message map[string]any)
	onBinary     func(data []byte)
	onDisconnect func(err *AccessError)
	onNotice     func(notice Notice)
}

type muxSocket struct {
	access   *accessCache
	handlers muxHandlers

	ctx    context.Context
	cancel context.CancelFunc

	dialMu sync.Mutex
	mu     sync.Mutex
	conn   *websocket.Conn
	closed bool

	reconnecting     bool
	reconnectAttempt int
}

func newMuxSocket(access *accessCache, handlers muxHandlers) *muxSocket {
	ctx, cancel := context.WithCancel(context.Background())
	return &muxSocket{access: access, handlers: handlers, ctx: ctx, cancel: cancel}
}

func (m *muxSocket) isClosed() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.closed
}

func (m *muxSocket) current() *websocket.Conn {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.conn
}

func (m *muxSocket) ensureOpen(ctx context.Context) error {
	if m.isClosed() {
		return closedError()
	}
	if m.current() != nil {
		return nil
	}

	m.dialMu.Lock()
	defer m.dialMu.Unlock()

	if m.isClosed() {
		return closedError()
	}
	if m.current() != nil {
		return nil
	}

	return m.connect(ctx)
}

func (m *muxSocket) connect(ctx context.Context) error {
	var lastErr *AccessError

	for attempt := 1; attempt <= maxConnectAttempts; attempt++ {
		if attempt > 1 {
			select {
			case <-time.After(connectBaseDelay * time.Duration(1<<(attempt-2))):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		if m.isClosed() {
			return closedError()
		}

		grant, err := m.access.get(ctx)
		if err != nil {
			return err
		}

		dialCtx, cancel := context.WithTimeout(ctx, connectTimeout)
		conn, resp, err := websocket.Dial(dialCtx, grant.URL, &websocket.DialOptions{
			Subprotocols: socketProtocols(grant.Token),
		})
		cancel()

		if err == nil {
			conn.SetReadLimit(maxMessageSize)

			m.mu.Lock()
			if m.closed {
				m.mu.Unlock()
				conn.Close(websocket.StatusNormalClosure, "closed")
				return closedError()
			}
			m.conn = conn
			m.mu.Unlock()

			go m.read(conn)
			return nil
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		status := 0
		if resp != nil {
			status = resp.StatusCode
		}
		lastErr = handshakeError(status, err)
		if !lastErr.Retryable {
			return lastErr
		}
		m.access.invalidate()
	}

	lastErr.unsent = true
	return lastErr
}

func (m *muxSocket) send(ctx context.Context, messageType websocket.MessageType, data []byte) bool {
	conn := m.current()
	if conn == nil {
		return false
	}
	return conn.Write(ctx, messageType, data) == nil
}

func (m *muxSocket) sendJSON(ctx context.Context, value any) bool {
	data, err := json.Marshal(value)
	if err != nil {
		return false
	}
	return m.send(ctx, websocket.MessageText, data)
}

func (m *muxSocket) close() {
	m.mu.Lock()
	if m.closed {
		m.mu.Unlock()
		return
	}
	m.closed = true
	conn := m.conn
	m.conn = nil
	m.mu.Unlock()

	m.cancel()
	if conn != nil {
		conn.Close(websocket.StatusNormalClosure, "closed")
	}

	m.handlers.onDisconnect(closedError())
}

func (m *muxSocket) read(conn *websocket.Conn) {
	var readErr error

	for {
		messageType, data, err := conn.Read(m.ctx)
		if err != nil {
			readErr = err
			break
		}

		if messageType == websocket.MessageBinary {
			m.handlers.onBinary(data)
		} else {
			m.handleText(data)
		}
	}

	m.mu.Lock()
	if m.closed || m.conn != conn {
		m.mu.Unlock()
		return
	}
	m.conn = nil
	m.mu.Unlock()

	code, reason := -1, ""
	var closeErr websocket.CloseError
	if errors.As(readErr, &closeErr) {
		code, reason = int(closeErr.Code), closeErr.Reason
	}

	accessErr := describeClose(code, reason)
	if accessErr.Code == "unauthorized" {
		m.access.invalidate()
	}

	conn.CloseNow()
	m.handlers.onDisconnect(accessErr)

	if reconnectCodes[accessErr.Code] {
		m.scheduleReconnect()
	}
}

func (m *muxSocket) scheduleReconnect() {
	m.mu.Lock()
	if m.closed || m.reconnecting {
		m.mu.Unlock()
		return
	}
	m.reconnecting = true
	delay := connectBaseDelay * time.Duration(1<<min(m.reconnectAttempt, 10))
	if delay > maxReconnectDelay {
		delay = maxReconnectDelay
	}
	m.reconnectAttempt++
	m.mu.Unlock()

	go func() {
		select {
		case <-time.After(delay):
		case <-m.ctx.Done():
			return
		}

		err := m.ensureOpen(m.ctx)

		m.mu.Lock()
		m.reconnecting = false
		if err == nil {
			m.reconnectAttempt = 0
		}
		m.mu.Unlock()

		var accessErr *AccessError
		if err != nil && errors.As(err, &accessErr) && reconnectCodes[accessErr.Code] {
			m.scheduleReconnect()
		}
	}()
}

func (m *muxSocket) handleText(data []byte) {
	var message map[string]any
	if err := json.Unmarshal(data, &message); err != nil {
		return
	}

	if message["type"] == "connection_notice" {
		if m.handlers.onNotice != nil {
			notice, _ := message["notice"].(string)
			reason, _ := message["reason"].(string)
			idle, _ := message["idleInSeconds"].(float64)
			m.handlers.onNotice(Notice{Notice: notice, Reason: reason, IdleInSeconds: idle})
		}
		return
	}

	m.handlers.onText(message)
}
