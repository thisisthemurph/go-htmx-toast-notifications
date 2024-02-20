package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"hx-toast/handler"
	"hx-toast/toast"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func customErrorHandler(err error, c echo.Context) {
	te, ok := err.(toast.Toast)
	if !ok {
		fmt.Println(err)
		te = toast.Danger("there has been an unexpected error")
	}

	if te.Level != toast.SUCCESS {
		c.Response().Header().Set("HX-Reswap", "none")
	}

	te.SetHXTriggerHeader(c)
}

func main() {
	// Precompile templates
	t := &Template{
		template.Must(template.ParseGlob("view/*.html")),
	}

	// Init echo
	e := echo.New()
	e.Renderer = t
	e.HTTPErrorHandler = customErrorHandler
	e.Static("/", "static")

	h := handler.NewHomeHandler()

	// Set up handlers
	e.GET("/", h.HandleIndexPage)
	e.POST("/newsletter", h.HandleNewsletterSignUp)

	// Start the server
	e.Logger.Fatal(e.Start(":5000"))
}
