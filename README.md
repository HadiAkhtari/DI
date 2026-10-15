# DI

A small Go HTTP service demonstrating dependency injection and application lifecycle management with Uber Fx.

## Requirements and development

Use Go 1.25.1 or newer. No database or credentials are required.

```sh
go mod download
go run ./cmd/api
```

The server defaults to port 8080. `CONFIG_ADDR` sets the complete listen address and takes precedence over `CONFIG_PORT`:

```sh
CONFIG_ADDR=127.0.0.1:8080 go run ./cmd/api
```

## API

`GET /greet?name=Hadi` returns JSON such as:

```json
{"message":"Good afternoon Hadi"}
```

The greeting follows the server's local clock: morning before 12:00, afternoon from 12:00 to 17:59, and night from 18:00. Omitted, empty, or whitespace-only names default to `myfriend`. Names are trimmed. Other HTTP methods receive status 405 with `Allow: GET`.

## Structure

- `cmd/api`: application entry point.
- `internal/fx`: dependency wiring.
- `internal/config`: environment configuration.
- `internal/controller`: HTTP routes, handlers, and server lifecycle.
- `internal/service`: greeting logic and replaceable time provider.
- `internal/logging`: simple console logging.

The server binds its socket before reporting startup success, sets request and idle timeouts, and shuts down gracefully on termination. An unexpected serving failure requests application shutdown with a nonzero exit code.

## Validation

```sh
go test -race ./...
go vet ./...
go build ./cmd/api
curl --fail 'http://127.0.0.1:8080/greet?name=Hadi'
```

Tests cover greeting time boundaries, HTTP responses and methods, configuration precedence, and failure to start on an occupied address.
