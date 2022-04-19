package router

import (
	"awesomeProject/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()

	// entry
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Up!")
	})

	// users
	u := e.Group("/api/users")
	u.GET("/", api.GetUsers)
	u.GET("/:id", api.GetUserById)
	u.POST("/email", api.GetUserByEmail)
	u.POST("/", api.CreateUser)
	u.POST("/login", api.Login)

	r := e.Group("/api/admin/users")
	r.PUT("/:id/email", api.UpdateUserEmail)
	r.PUT("/:id/password", api.UpdateUserPassword)
	r.DELETE("/:id", api.DeleteUser)
	r.DELETE("/:id/final", api.DeleteUserPermanently)

	return e
}
