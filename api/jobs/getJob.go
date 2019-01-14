package jobs

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"

	jobs "github.com/davlis/cron/api/jobs/types"
	"github.com/davlis/cron/storage"
	"github.com/labstack/echo"
)

func GetJob(c echo.Context) error {
	log.Println("api/jobs/getJob : Get Job started")

	id := c.Param("id")

	job := &jobs.Job{}

	err := storage.Connection.Collection("jobs").FindById(bson.ObjectIdHex(id), job)

	if err != nil {
		log.Printf("%s: %s", "api/jobs/getJob : Failed to fetch job from database", err)
		return c.JSON(http.StatusNotFound, `Job not found`) // TODO: serialized error
	}

	log.Printf("api/jobs/getJob : Get Job finished")
	return c.JSON(http.StatusOK, job)
}
