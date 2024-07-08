package main

import (
	"fmt"
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

type SiteState struct {
	current_view string
	siteInfo     []siteInfo
}

var mySites = []siteInfo{
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

// templeting stuff
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// handlers
func index(c echo.Context) error {

	cookie, err := c.Cookie("current_page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)

	state := SiteState{
		current_view: cookie.Value,
		siteInfo:     mySites,
	}

	return c.Render(http.StatusOK, "index", state)
}

func home(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "current_page"
	cookie.Path = "/"
	cookie.Value = "home"
	c.SetCookie(cookie)
	return c.Render(http.StatusOK, "home", nil)
}

func sites(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "current_page"
	cookie.Path = "/"

	cookie.Value = "sites"
	c.SetCookie(cookie)

	return c.Render(http.StatusOK, "sites", mySites)
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
