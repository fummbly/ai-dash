package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Response any
}

func NewData() *Data {
	return &Data{
		Response: "",
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}

func main() {
	e := echo.New()

	e.Renderer = NewTemplate()
	data := NewData()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})
	e.GET("/models", func(c echo.Context) error {
		res, err := http.Get("http://localhost:11434/api/tags")
		if err != nil {
			return c.Render(404, "index", data)
		}
		defer res.Body.Close()

		dat, err := io.ReadAll(res.Body)
		if err != nil {
			return c.Render(404, "index", data)
		}

		var js map[string]interface{}

		if err = json.Unmarshal(dat, &js); err != nil {
			return c.Render(404, "index", data)
		}

		data.Response = js
		return c.Render(200, "response", data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
