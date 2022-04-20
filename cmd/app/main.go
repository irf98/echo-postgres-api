package main

import (
	"awesomeProject/db"
	"awesomeProject/router"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := router.New()

	e.Use(middleware.CORS())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	db.Init()

	log.Println("Server Up!")
	e.Logger.Fatal(e.Start(":8080"))
}
