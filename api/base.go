package api

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Entry(c echo.Context) error {
	return c.String(http.StatusOK, "Server Up!")
}

func WhoAmI(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if sess.Values["user"] == nil {
		return c.JSON(http.StatusOK, "anonymous")
	}

	return c.JSON(http.StatusOK, sess.Values["user"])
}
