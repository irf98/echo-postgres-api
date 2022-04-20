package router

import (
	"awesomeProject/api"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// entry
	e.GET("/", api.Entry)
	e.GET("/who", api.WhoAmI)

	// users
	u := e.Group("/api/users")

	u.GET("/", api.GetUsers)
	u.GET("/:id", api.GetUserById)
	u.GET("/logout", api.Logout)
	u.POST("/login", api.Login)
	u.POST("/email", api.GetUserByEmail)
	u.POST("/", api.CreateUser)
	u.PUT("/:id/email", api.UpdateUserEmail)
	u.PUT("/:id/password", api.UpdateUserPassword)
	u.DELETE("/:id", api.DeleteUser)
	u.DELETE("/:id/final", api.DeleteUserPermanently)

	return e
}
