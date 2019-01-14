package jobs

import (
	"log"

	"github.com/labstack/echo"
)

func CreateRouter(r *echo.Group) error {
	log.Println("api/jobs/createRouter : Create Router started")

	r.POST("/", CreateJob)
	r.GET("/:id", GetJob)
	r.PUT("/:id", UpdateJob)
	r.DELETE("/:id", DeleteJob)

	log.Println("api/jobs/createRouter : Create Router finished")
	return nil
}
