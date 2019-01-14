package jobs

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func DeleteJob(c echo.Context) error {
	log.Println("api/jobs/deleteJob : Delete Job started")

	log.Printf("api/jobs/deleteJob : Delete Job finished")
	return c.JSON(http.StatusNotImplemented, "Not implemented")
}
