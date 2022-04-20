package config

import (
	"awesomeProject/api/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GenerateSession(c echo.Context, user models.User) error {
	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
	}

	sess.Values["user"] = user.Email
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func DeleteSession(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["user"] == nil {
		return c.NoContent(http.StatusBadRequest)
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
	}

	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
