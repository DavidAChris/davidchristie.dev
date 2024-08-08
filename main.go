package main

import (
	"davidchristie-dev/database"
	"davidchristie-dev/handlers"
	"davidchristie-dev/services"
	"log"

	mw "davidchristie-dev/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupAws() {
	s3Svc, details := services.ConfigAws()
	log.Println("About to Init DB")
	err := database.InitDb(s3Svc, details)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *echo.Echo {
	server := echo.New()
	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Use(middleware.BodyLimit("5M"))

	authReq := server.Group("/")
	authReq.Use(mw.Auth)
	authReq.GET("/", handlers.Home)
	authReq.GET("home", handlers.Home)

	server.Static("/static", "static")
	server.GET("/login", handlers.Login)
	server.GET("/signup", handlers.Signup)
	return server
}

func main() {
	setupAws()
	server := setupRouter()

	server.Logger.Fatal(server.Start(":3000"))
}
