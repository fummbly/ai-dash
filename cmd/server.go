package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	status "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"www.github.com/fummbly/ai-dash/internal/adapters/ai"
	"www.github.com/fummbly/ai-dash/internal/adapters/http"
	"www.github.com/fummbly/ai-dash/internal/domain"
	"www.github.com/fummbly/ai-dash/internal/service"
)

type Data struct {
	Response domain.Models
}

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Renderer = NewTemplate()

	e.Use(middleware.RequestLogger())

	// TODO create seperate handler functions for routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(status.StatusOK, "index", Data{})
	})

	e.GET("/chat", func(c echo.Context) error {
		return c.Render(status.StatusOK, "chat", Data{})
	})

	e.GET("/generate", func(c echo.Context) error {
		return c.Render(status.StatusOK, "generate", Data{})
	})

	e.GET("/stream", func(c echo.Context) error {
		log.Printf("SSE client connected, ip: %v", c.RealIP())
		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(status.StatusOK)

		dataChan := make(chan string)
		go http.StreamPost("http://localhost:11434/api/generate", "application/json", `{
				"model": "gemma3:1b",
				"prompt": "Why is the sky blue?"
				}`, dataChan)
		message := ""
		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("SSE client disconnected, ip: %v", c.RealIP())
				return nil
			case data, ok := <-dataChan:
				if !ok {
					log.Printf("Generated data finished")
					if _, err := fmt.Fprintf(w, "event: GenerateFinished\ndata: <div>Generated Data finished</div>\n\n"); err != nil {
						return err
					}
					w.Flush()

					return nil
				}
				log.Printf("Generated response: %s", data)
				message += data
				if _, err := fmt.Fprintf(w, "data: <div>%s</div>\n\n", message); err != nil {
					return err
				}
				w.Flush()
			}
		}
	})

	modelService := service.NewModelService(ai.NewAIModelEnpoint("http://localhost:11434/api"))

	modelHandler := http.NewModelHandler(*modelService)

	e.GET("/models", modelHandler.ListModels)

	e.Logger.Fatal(e.Start(":1323"))
}
