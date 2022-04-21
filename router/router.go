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

	// posts
	p := e.Group("/api/posts")

	p.GET("/:id", api.GetUserPosts)
	p.POST("/tag", api.GetPostsByTag)
	p.POST("/", api.CreatePost)
	p.PUT("/:id/upvote", api.UpVotePost)
	p.PUT("/:id/downvote", api.DownVotePost)
	p.DELETE("/:id", api.DeletePost)

	return e
}
