package main

import (
	"awesomeProject/db"
	"awesomeProject/router"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := router.New()

	e.Use(middleware.CORS())

	db.Init()

	log.Println("Server Up!")
	e.Logger.Fatal(e.Start(":8080"))
}
