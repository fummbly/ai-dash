package main

import (
	"html/template"
	"io"
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

	responseService := service.NewResponseService(ai.NewAIResponseEndpoint("http://localhost:11434/api"))
	responseHandler := http.NewResponseHandler(*responseService)

	e.GET("/stream", responseHandler.StreamResponse)

	modelService := service.NewModelService(ai.NewAIModelEnpoint("http://localhost:11434/api"))

	modelHandler := http.NewModelHandler(*modelService)

	e.GET("/models", modelHandler.ListModels)

	e.Logger.Fatal(e.Start(":1323"))
}
