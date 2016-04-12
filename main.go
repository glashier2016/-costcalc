package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/glashier2016/costcalc/pages"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}

	e := echo.New()
	e.SetRenderer(t)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Get("/auth", pages.Auth)
	e.Run(standard.New(":8080"))
}
