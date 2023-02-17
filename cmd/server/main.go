package main

import (
	"html/template"
	"io"
	"todo-htmx/internal/handler"
	"todo-htmx/internal/store/inmem"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("web/*.html")),
	}

	// Echo instance
	e := echo.New()

	e.Binder = &echo.DefaultBinder{}
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handler.New(inmem.New())
	e.GET("/", h.Root)
	e.GET("/item", h.AddItem)
	e.POST("/item/", h.PostItem)
	// echo calls http.Request.ParseForm but it does not parse form body for DELETE requests
	e.POST("/delete/item/", h.DeleteItem)

	e.Logger.Fatal(e.Start(":8080"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	c.FormParams()
	return t.templates.ExecuteTemplate(w, name, data)
}
