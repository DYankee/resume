package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// templeting stuff
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// handlers
func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
func resume(c echo.Context) error {
	return c.Render(http.StatusOK, "resume", nil)
}
func sites(c echo.Context) error {
	return c.Render(http.StatusOK, "sites", nil)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	//css files
	e.Static("/dist", "dist")

	//init handlers
	e.GET("/", index)
	e.GET("views/home", home)
	e.GET("views/resume", resume)
	e.GET("views/sites", sites)

	e.Logger.Fatal(e.Start(":8080"))
}
