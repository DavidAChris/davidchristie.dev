package main

import (
	"davidchristie-dev/database"
	"davidchristie-dev/services"
	"log"

	mw "davidchristie-dev/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
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
	server.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))
	server.Use(middleware.Gzip())
	server.Use(middleware.Decompress())
	server.Use(middleware.Secure())

	authReq := server.Group("/")
	authReq.Use(mw.Auth)
	authReq.File("/", "./views/index.html")
	authReq.File("home", "./views/index.html")

	server.File("/login", "./views/login.html")
	server.File("/signup", "./views/signup.html")

	server.Static("/", "./static/")
	return server
}

func main() {
	setupAws()
	server := setupRouter()

	server.Logger.Fatal(server.Start(":3000"))
}
