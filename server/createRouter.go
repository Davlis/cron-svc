package server

import (
	"log"

	"github.com/davlis/cron/api"
	"github.com/labstack/echo"
)

func CreateRouter(r *echo.Group) error {
	log.Println("server/createRouter : Create Router started")

	api.CreateRouter(r.Group("/api"))

	log.Println("server/createRouter : Create Router finished")
	return nil
}
