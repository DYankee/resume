package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type siteInfo struct {
	SiteName string
	SiteLink string
	SiteImg  string
	SiteDesc string
	SiteTime string
	SiteTech []string
}

// templating stuff
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

func sites(c echo.Context) error {
	var sites = []siteInfo{}
	file, err := os.ReadFile("resources/site-data.json")
	if err != nil {
		c.Logger().Panic()
	}
	err = json.Unmarshal(file, &sites)
	if err != nil {
		c.Logger().Panic()
	}
	return c.Render(http.StatusOK, "sites", sites)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	//css files
	e.Static("/dist", "dist")
	e.Static("/resources", "resources")
	e.Static("./oldsites", "dota")

	//init handlers
	e.GET("/", index)
	e.GET("views/home", home)
	e.GET("views/sites", sites)

	e.Logger.Fatal(e.Start(":8080"))
}
