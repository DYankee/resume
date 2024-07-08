package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type siteInfo struct {
	SiteName string
	SiteDesc string
	SiteLink string
	SiteImg  string
}

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

func sites(c echo.Context) error {
	sites := []siteInfo{
		{
			SiteName: "The world of AI art",
			SiteDesc: "Website I created in my second semester of collage. Uses html and css",
			SiteLink: "https://woa.gearyhs.com/",
			SiteImg:  "",
		},
		{
			SiteName: "OCC connect",
			SiteDesc: "website I created in my third semester of collage. Uses php, mySQL, html, and css.",
			SiteLink: "https://occ-connect.zgeary.dev",
			SiteImg:  "",
		},
		{
			SiteName: "Dota 2 Fan Site",
			SiteDesc: "website I created in my first semester of collage. Uses html and css.",
			SiteLink: "https://occ-connect.zgeary.dev",
			SiteImg:  "",
		},
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
	e.Static("/res", "res")
	e.Static("./oldsites", "dota")

	//init handlers
	e.GET("/", index)
	e.GET("views/home", home)
	e.GET("views/sites", sites)

	e.Logger.Fatal(e.Start(":8080"))
}
