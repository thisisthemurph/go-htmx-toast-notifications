package handler

import (
	"github.com/labstack/echo/v4"
	"hx-toast/toast"
	"net/http"
	"strings"
)

type HomeHandler struct{}

func NewHomeHandler() HomeHandler {
	return HomeHandler{}
}

func (h HomeHandler) HandleIndexPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (h HomeHandler) HandleNewsletterSignUp(c echo.Context) error {
	name := strings.TrimSpace(c.FormValue("name"))
	email := strings.TrimSpace(c.FormValue("email"))

	if err := validateName(name); err != nil {
		return err
	}

	if err := validateEmailAddress(email); err != nil {
		return err
	}

	toast.Success(c, "Successfully signed up to the newsletter!")
	return c.Render(http.StatusOK, "success_partial.html", name)
}

// validateName returns an error if the name is not valid or if the name is Tom.
// This is a terrible implementation.
func validateName(name string) error {
	if name == "" {
		return toast.Warning("How are we supposed to spam you if we don't know your name?")
	}

	if strings.ToLower(name) == "tom" {
		return toast.Danger("Tom is a genius, there is no need to subscribe!")
	}

	return nil
}

// validateEmailAddress returns an error if the email is not valid.
// This is a terrible implementation.
func validateEmailAddress(email string) error {
	if email == "" {
		return toast.Warning("An email address must be provided")
	}

	if !strings.Contains(email, "@") {
		return toast.Warning("the email address must include an @ symbol")
	}

	return nil
}
