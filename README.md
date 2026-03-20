# Metorial Go SDK

Go client library for the [Metorial API](https://metorial.com). Provides typed access to all Metorial resources including providers, deployments, sessions, and more.

## Installation

```bash
go get github.com/metorial/metorial-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    metorial "github.com/metorial/metorial-go/v1"
)

func main() {
    sdk, err := metorial.New(
        metorial.WithAPIKey("metorial_sk_..."),
    )
    if err != nil {
        log.Fatal(err)
    }

    providers, err := sdk.Providers.List(nil)
    if err != nil {
        log.Fatal(err)
    }

    for _, p := range providers.Items {
        fmt.Printf("%s: %s\n", p.Id, p.Name)
    }
}
```

## Configuration

```go
// Required: API key
sdk, err := metorial.New(
    metorial.WithAPIKey("metorial_sk_..."),
)

// Optional: custom API host
sdk, err := metorial.New(
    metorial.WithAPIKey("metorial_sk_..."),
    metorial.WithAPIHost("https://custom-api.example.com"),
)

// Optional: additional headers
sdk, err := metorial.New(
    metorial.WithAPIKey("metorial_sk_..."),
    metorial.WithHeaders(map[string]string{
        "X-Custom-Header": "value",
    }),
)
```

## Usage

All endpoint methods follow the same pattern: path parameters as positional arguments, optional query/body parameters as the last argument (pass `nil` to use defaults).

### Providers

Providers are read-only templates for MCP server integrations (like GitHub or Slack).

```go
// List providers
providers, err := sdk.Providers.List(nil)

// List with query parameters
limit := float64(10)
providers, err := sdk.Providers.List(&endpoints.ProvidersEndpointListParams{
    Limit: &limit,
})

// Get a specific provider
provider, err := sdk.Providers.Get("prov_abc123")

// Access provider sub-resources
versions, err := sdk.ProvidersVersions.List(nil)
tools, err := sdk.ProvidersTools.List("prov_abc123", nil)
```

### Provider Deployments

Deployments are configured instances of a provider within your project.

```go
// List deployments
deployments, err := sdk.ProviderDeployments.List(nil)

// Create a deployment
deployment, err := sdk.ProviderDeployments.Create(&endpoints.ProviderDeploymentsEndpointCreateBody{
    ProviderId: "prov_abc123",
})

// Update a deployment
deployment, err := sdk.ProviderDeployments.Update("pd_abc123", &endpoints.ProviderDeploymentsEndpointUpdateBody{
    Name: ptr("Updated Name"),
})

// Delete a deployment
_, err := sdk.ProviderDeployments.Delete("pd_abc123")

// Deployment sub-resources
configs, err := sdk.ProviderDeploymentsConfigs.List(nil)
authConfigs, err := sdk.ProviderDeploymentsAuthConfigs.List(nil)
```

### Sessions

Sessions represent active MCP communication channels.

```go
// List sessions
sessions, err := sdk.Sessions.List(nil)

// Create a session
session, err := sdk.Sessions.Create(&endpoints.SessionsEndpointCreateBody{
    Name:      ptr("My Session"),
    Providers: []map[string]any{
        {"provider_deployment_id": "pd_abc123"},
    },
})

// Get a specific session
session, err := sdk.Sessions.Get("sess_abc123")

// Update a session
session, err := sdk.Sessions.Update("sess_abc123", &endpoints.SessionsEndpointUpdateBody{
    Name: ptr("Renamed Session"),
})

// Delete a session
_, err := sdk.Sessions.Delete("sess_abc123")

// Session sub-resources (separate endpoint fields)
messages, err := sdk.SessionsMessages.List(nil)
events, err := sdk.SessionsEvents.List(nil)
connections, err := sdk.SessionsConnections.List(nil)
```

### Session Templates

Templates define reusable session configurations.

```go
templates, err := sdk.SessionTemplates.List(nil)

template, err := sdk.SessionTemplates.Create(&endpoints.SessionTemplatesEndpointCreateBody{
    Name: "My Template",
})
```

### Custom Providers

Define and manage your own MCP server providers.

```go
customs, err := sdk.CustomProviders.List(nil)

custom, err := sdk.CustomProviders.Get("cp_abc123")
```

## Pagination

List endpoints return paginated results. The response includes a `Pagination` field with `HasMoreBefore` and `HasMoreAfter` booleans. Use the `After`, `Before`, or `Cursor` parameters to navigate.

```go
limit := float64(10)

// First page
page, err := sdk.Providers.List(&endpoints.ProvidersEndpointListParams{
    Limit: &limit,
})

for _, item := range page.Items {
    fmt.Println(item.Id)
}

// Check if there are more results
if page.Pagination.HasMoreAfter {
    // Fetch next page using the last item's ID
    lastID := page.Items[len(page.Items)-1].Id
    next, err := sdk.Providers.List(&endpoints.ProvidersEndpointListParams{
        Limit: &limit,
        After: &lastID,
    })
}
```

## Error Handling

All methods return `(result, error)`. API errors are returned as `*metorial.Error`.

```go
import (
    "errors"
    "fmt"

    metorial "github.com/metorial/metorial-go/v1"
)

provider, err := sdk.Providers.Get("invalid_id")
if err != nil {
    var apiErr *metorial.Error
    if errors.As(err, &apiErr) {
        fmt.Println("Status:", apiErr.StatusCode)
        fmt.Println("Message:", apiErr.Message)
        fmt.Println("Request ID:", apiErr.RequestID)

        switch {
        case apiErr.IsNotFound():
            fmt.Println("Resource not found")
        case apiErr.IsUnauthorized():
            fmt.Println("Invalid API key")
        case apiErr.IsRateLimited():
            fmt.Println("Rate limited")
        }
    } else {
        // Network error or other non-API error
        fmt.Println("Request failed:", err)
    }
}
```

Rate-limited requests (HTTP 429) are automatically retried up to 3 times with backoff.

## Pointer Helpers

Optional fields use pointer types. A generic helper is useful:

```go
func ptr[T any](v T) *T {
    return &v
}

// Usage
sdk.ProviderDeployments.Update("pd_abc", &endpoints.ProviderDeploymentsEndpointUpdateBody{
    Name:        ptr("New Name"),
    Description: ptr("New description"),
})
```

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Support

[Documentation](https://metorial.com/docs) · [GitHub Issues](https://github.com/metorial/metorial-node/issues) · [Email Support](mailto:support@metorial.com)
