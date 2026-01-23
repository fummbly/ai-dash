# AGENTS.md - AI Coding Agent Guidelines

This file provides guidelines, commands, and style rules for agents working in the Go‑based AI dashboard project.

## Project Overview

- `cmd/`: application entry points
- `internal/adapters/`: external integrations (HTTP, AI)
- `internal/domain/`: business entities and interfaces
- `internal/service/`: business logic layer
- `internal/parser/`: text processing utilities

## Build, Test, and Development Commands

### Build
```bash
# Build binary
go build -o ./tmp/main cmd/server.go
# Hot‑reload with Air
air
```

### Tests
```bash
# Run all tests
go test ./...
# Verbose output
go test -v ./...
# Specific test file
go test ./internal/parser/parser_test.go
# Specific test function
go test -run TestBoldParse ./internal/parser/
# Test with coverage
go test -cover ./...
# Generate HTML coverage report
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

### Lint & Quality
```bash
golangci-lint run
go fmt ./...
go vet ./...
go mod tidy
```

### Development Server
```bash
go run cmd/server.go
air  # recommends hot‑reload
```

## Code Style Guidelines

### Import Organization
Group imports in the following order: standard library → third‑party → internal packages. Separate groups with a blank line.
```go
import (
    "html/template"
    "io"
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/fummbly/ai-dash/internal/adapters/ai"
    "github.com/fummbly/ai-dash/internal/domain"
)
```

### Naming Conventions
- Packages: lowercase, single word
- Structs/Interfaces: PascalCase
- Functions: exported PascalCase, unexported camelCase
- Variables: camelCase
- Constants: UPPER_SNAKE_CASE

### Error Handling
Always return errors explicitly with context. Wrap only when adding meaningful information.
```go
data, err := process(input)
if err != nil {
    return fmt.Errorf("process failed: %w", err)
}
```

### Constructors & Dependency Injection
```go
type ResponseHandler struct {
    svc service.ResponseService
}
func NewResponseHandler(svc service.ResponseService) *ResponseHandler {
    return &ResponseHandler{svc: svc}
}
```

### Testing Guidelines
Use table‑driven tests and `stretchr/testify/assert`.
```go
func TestConvertBold(t *testing.T) {
    cases := []struct{ name, in, want string }{
        {"basic", "\**bold\**", "<strong>bold</strong>"},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            got := ConvertBold(c.in)
            assert.Equal(t, c.want, got)
        })
    }
}
```

### Concurrency & Channels
Always close channels; use `select` when awaiting cancellation.
```go
resC := make(chan domain.Response)
defer close(resC)
go func() {
    if err := h.svc.Generate(resC, q); err != nil {
        log.Println(err)
    }
}()
```

## Project‑Specific Patterns
- **Clean Architecture**: Domain → Service → Adapters → Presentation
- **AI Integration**: Ollama API at `http://localhost:11434/api`, SSE for streaming
- **Template Rendering**: `html/template` templates stored in `public/views/`

## Configuration
- `.air.toml`: Air hot‑reload config
- `.golangci.yml`: Lint configuration
- `go.mod`: Module dependencies
- `.gitignore`: Only ignores `/tmp/`

## Development Notes
- Default port: **1323**
- Echo v4
- Follow clean architecture throughout the codebase

## Cursor / Copilot Rules
None defined in this repository.
