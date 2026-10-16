package sandbox

import (
	"errors"
	"fmt"
	"strings"
)

// AccessError is returned for failures reaching or using a sandbox connection.
// Code is the server's error code (e.g. "not_found", "permission_denied") or one
// of the connection codes: "closed", "disconnected", "unauthorized", "rate_limited",
// "machine_paused", "machine_restarting", "machine_terminated", "machine_unavailable",
// "machine_failed", "idle_timeout", "unavailable".
type AccessError struct {
	Code      string
	Message   string
	Retryable bool
	Path      string

	unsent bool
}

func (e *AccessError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("sandbox: %s (%s): %s", e.Code, e.Path, e.Message)
	}
	return fmt.Sprintf("sandbox: %s: %s", e.Code, e.Message)
}

// CommandError is returned when a command fails to run or loses its connection.
type CommandError struct {
	AccessError
}

func (e *CommandError) Unwrap() error { return &e.AccessError }

// ErrorCode returns the Code of an AccessError or CommandError in err's chain.
func ErrorCode(err error) string {
	var accessErr *AccessError
	if errors.As(err, &accessErr) {
		return accessErr.Code
	}
	return ""
}

func newError(code, message string, retryable bool) *AccessError {
	if message == "" {
		message = strings.ReplaceAll(code, "_", " ")
	}
	return &AccessError{Code: code, Message: message, Retryable: retryable}
}

func newCommandError(code, message string) *CommandError {
	return &CommandError{AccessError: *newError(code, message, false)}
}

func closedError() *AccessError {
	return newError("closed", "The sandbox connection was closed.", false)
}

func serverError(raw any) *AccessError {
	data, _ := raw.(map[string]any)
	code, _ := data["code"].(string)
	message, _ := data["message"].(string)
	path, _ := data["path"].(string)
	retryable, _ := data["retryable"].(bool)

	if code == "" {
		code = "unknown"
	}

	err := newError(code, message, retryable)
	err.Path = path
	return err
}

var closeReasons = map[string]string{
	"machine_paused":     "The machine paused, so this connection was closed.",
	"machine_restarting": "The machine is restarting.",
	"machine_terminated": "The machine was terminated.",
	"idle_timeout":       "This connection was closed after being idle.",
}

func describeClose(code int, reason string) *AccessError {
	switch code {
	case 4401:
		return newError("unauthorized", "Access to the machine was rejected.", true)
	case 4404:
		return newError("not_found", "This machine could not be found.", false)
	case 4410:
		return newError("machine_terminated", closeReasons["machine_terminated"], false)
	case 4429:
		return newError("rate_limited", "Too many connections to the machine.", true)
	case 4503:
		if reason == "machine_failed" {
			return newError("machine_failed", "The machine has failed.", false)
		}
		return newError("machine_unavailable", "The machine is not available.", true)
	case 1012:
		kind := "machine_paused"
		if reason == "machine_terminated" || reason == "machine_restarting" {
			kind = reason
		}
		return newError(kind, closeReasons[kind], kind != "machine_terminated")
	case 1000:
		if reason == "idle_timeout" {
			return newError("idle_timeout", closeReasons["idle_timeout"], true)
		}
	}

	return newError("disconnected", "The connection to the machine was lost.", true)
}

func handshakeError(status int, cause error) *AccessError {
	switch status {
	case 401, 403:
		return newError("unauthorized", "Access to the machine was rejected.", true)
	case 404:
		return newError("not_found", "This machine could not be found.", false)
	case 429:
		return newError("rate_limited", "Too many connections to the machine.", true)
	}

	message := "The connection to the machine failed."
	if cause != nil {
		message = cause.Error()
	}
	return newError("disconnected", message, true)
}
