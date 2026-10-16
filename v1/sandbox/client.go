// Package sandbox opens file, command, and port connections to sandbox machines.
package sandbox

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/metorial/metorial-go/v1/endpoints"
)

type ConnectOptions struct {
	// TTL is how long each access token stays valid; zero uses the server default.
	TTL      time.Duration
	OnNotice func(Notice)
}

type PortOptions struct {
	TTL        time.Duration
	HTTPClient *http.Client
}

type Client struct {
	access *endpoints.SandboxMachinesAccessEndpoint
}

func NewClient(access *endpoints.SandboxMachinesAccessEndpoint) *Client {
	return &Client{access: access}
}

// Files opens a connection to the machine's file service.
func (c *Client) Files(ctx context.Context, machineID string, opts *ConnectOptions) (*Files, error) {
	if opts == nil {
		opts = &ConnectOptions{}
	}

	ttl := ttlSeconds(opts.TTL)
	access := newAccessCache(func(context.Context) (*Grant, error) {
		grant, err := c.access.Files(machineID, &endpoints.SandboxMachinesAccessEndpointFilesBody{TtlSeconds: ttl})
		if err != nil {
			return nil, err
		}
		return newGrant(grant.Urls.Files, grant.Token, grant.ExpiresAt, "File")
	})

	return connectFiles(ctx, access, opts.OnNotice)
}

// Shell opens a connection for running commands in the machine.
func (c *Client) Shell(ctx context.Context, machineID string, opts *ConnectOptions) (*Shell, error) {
	if opts == nil {
		opts = &ConnectOptions{}
	}

	ttl := ttlSeconds(opts.TTL)
	access := newAccessCache(func(context.Context) (*Grant, error) {
		grant, err := c.access.Shell(machineID, &endpoints.SandboxMachinesAccessEndpointShellBody{TtlSeconds: ttl})
		if err != nil {
			return nil, err
		}
		return newGrant(grant.Urls.Exec, grant.Token, grant.ExpiresAt, "Command")
	})

	return connectShell(ctx, access, opts.OnNotice)
}

// Connect opens an HTTP connection to a port inside the machine.
func (c *Client) Connect(ctx context.Context, machineID string, port int, opts *PortOptions) (*Port, error) {
	if opts == nil {
		opts = &PortOptions{}
	}

	ttl := ttlSeconds(opts.TTL)
	access := newAccessCache(func(context.Context) (*Grant, error) {
		grant, err := c.access.Ports(machineID, &endpoints.SandboxMachinesAccessEndpointPortsBody{
			Ports:      []int{port},
			TtlSeconds: ttl,
		})
		if err != nil {
			return nil, err
		}

		var url *string
		if grant.Urls.Ports != nil {
			if value, ok := (*grant.Urls.Ports)[strconv.Itoa(port)]; ok {
				url = &value
			}
		}
		return newGrant(url, grant.Token, grant.ExpiresAt, fmt.Sprintf("Port %d", port))
	})

	return connectPort(ctx, port, access, opts.HTTPClient)
}

func newGrant(url *string, token string, expiresAt time.Time, label string) (*Grant, error) {
	if url == nil || *url == "" {
		return nil, newError("unavailable", label+" access is not available for this machine.", false)
	}
	return &Grant{URL: *url, Token: token, ExpiresAt: expiresAt}, nil
}

func ttlSeconds(ttl time.Duration) *float64 {
	if ttl <= 0 {
		return nil
	}
	seconds := ttl.Seconds()
	return &seconds
}
