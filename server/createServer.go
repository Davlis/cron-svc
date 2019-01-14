package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateServer() error {
	log.Println("server/createServer : Create Server started")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	CreateRouter(e.Group(""))

	e.Logger.Fatal(e.Start(":3000"))

	log.Println("server/createServer : Create Server finished")

	return nil
}
