package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "unsafe.html", map[string]interface{}{
		"username": "Not Yet Set",
		"password": "Not Yet Set",
	})
}

func nocheck(c echo.Context) error {
	return c.Render(http.StatusOK, "unsafe.html", map[string]interface{}{
		"username": c.FormValue("username"),
		"password": c.FormValue("password"),
	})
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = t

	e.GET("/", index)
	e.POST("/postform", nocheck)

	fmt.Println("Listening on http://localhost:8080/")

	e.Logger.Fatal(e.Start(":8080"))

}
