package api

import (
	"log"

	"github.com/labstack/echo"

	"github.com/davlis/cron/api/jobs"
)

func CreateRouter(r *echo.Group) error {
	log.Println("api/createRouter : Create Router started")

	jobs.CreateRouter(r.Group("/jobs"))

	log.Println("api/createRouter : Create Router finished")
	return nil
}
