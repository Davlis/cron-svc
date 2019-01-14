package jobs

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func UpdateJob(c echo.Context) error {
	log.Println("api/jobs/updateJob : Update Job started")

	log.Printf("api/jobs/updateJob : Update Job finished")
	return c.JSON(http.StatusNotImplemented, "Not implemented")
}
