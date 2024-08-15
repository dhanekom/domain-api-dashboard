package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("internal/views/*.html")),
	}
}

func routes(handlerRepo *HandlerRepo) *echo.Echo {
	e := echo.New()

	e.Static("/public", "public")
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.GET("/", handlerRepo.DashboardHandler)
	e.POST("/search", handlerRepo.DashboardSearchResultsHandler)

	return e
}
