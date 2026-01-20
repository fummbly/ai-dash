# AGENTS.md - AI Coding Agent Guidelines

This file contains guidelines and commands for agentic coding agents working in this Go-based AI dashboard project.

## Project Overview

Go web application using Echo framework with clean architecture:
- **cmd/**: Application entry points
- **internal/adapters/**: External integrations (HTTP, AI)
- **internal/domain/**: Business entities and interfaces
- **internal/service/**: Business logic layer
- **internal/parser/**: Text processing utilities

## Build, Test, and Development Commands

### Build Commands
```bash
go build -o ./tmp/main cmd/server.go
air  # Build and run with hot reload
```

### Test Commands
```bash
go test ./...                    # Run all tests
go test -v ./...                 # Verbose output
go test ./internal/parser/parser_test.go  # Specific test file
go test -run TestBoldParse ./internal/parser/  # Specific test function
go test -cover ./...             # With coverage
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

### Lint and Quality Commands
```bash
golangci-lint run
go fmt ./...
go vet ./...
go mod tidy
```

### Development Server
```bash
go run cmd/server.go
air  # Hot reload (recommended)
```

## Code Style Guidelines

### Import Organization
Group imports: standard library → third-party → internal packages. Blank line between groups.

```go
import (
    "html/template"
    "io"
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "www.github.com/fummbly/ai-dash/internal/adapters/ai"
    "www.github.com/fummbly/ai-dash/internal/domain"
)
```

### Naming Conventions
- **Packages**: lowercase, single word (e.g., `parser`, `service`, `domain`)
- **Structs**: PascalCase, descriptive (`ResponseHandler`, `AIResponseEndpoint`)
- **Functions**: PascalCase exported, camelCase unexported
- **Variables**: camelCase, descriptive
- **Constants**: UPPER_SNAKE_CASE for exported
- **Interfaces**: Often end with "er" suffix or descriptive (`ResponseInterface`, `ModelInterface`)

### Error Handling
Handle errors explicitly, use descriptive messages, return without wrapping unless necessary.

```go
data, err := process(input)
if err != nil {
    return fmt.Errorf("failed to process input: %w", err)
}
```

### Struct and Function Patterns
Use constructor functions, dependency injection through constructors, consistent receiver names.

```go
type ResponseHandler struct {
    responseService service.ResponseService
}

func NewResponseHandler(svc service.ResponseService) *ResponseHandler {
    return &ResponseHandler{responseService: svc}
}

func (h *ResponseHandler) StreamResponse(c echo.Context) error { }
```

### Interface Design
Define interfaces in domain layer, keep small and focused, use dependency injection.

### Testing Guidelines
Use table-driven tests, descriptive names, `github.com/stretchr/testify/assert`.

```go
func TestConvertBold(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"basic bold conversion", "text **bold** text", "text <strong>bold</strong> text"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ConvertBold(tt.input)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

### Channel and Concurrency Patterns
Use channels for streaming, always close when done, use select statements, handle context cancellation.

```go
resChan := make(chan domain.Response)
defer close(resChan)
go func() {
    err := h.responseService.Generate(resChan, question)
    if err != nil { /* handle error */ }
}()
```

## Project-Specific Patterns

### Clean Architecture
- **Domain Layer**: Interfaces and business entities
- **Service Layer**: Business logic using domain interfaces
- **Adapter Layer**: External integrations (HTTP clients, AI endpoints)
- **Presentation Layer**: Echo handlers and routing

### AI Integration
- Ollama API (default: http://localhost:11434/api)
- Streaming responses using Server-Sent Events
- JSON unmarshaling for AI response parsing

### HTML Template Rendering
- Go `html/template` package, templates in `public/views/`
- Custom renderer implements Echo's `Render` interface

## Configuration

- **.air.toml**: Air hot reload configuration
- **.golangci.yml**: Golangci-lint configuration (minimal setup)
- **go.mod**: Go module dependencies
- **.gitignore**: Git ignore patterns (only /tmp/)

## Development Notes

- Application runs on port 1323 by default
- Uses Echo v4 web framework
- Integrates with Ollama for AI functionality
- Implements Server-Sent Events for streaming responses
- Follows clean architecture principles