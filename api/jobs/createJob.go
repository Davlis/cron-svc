package jobs

import (
	"log"
	"net/http"

	jobs "github.com/davlis/cron/api/jobs/types"
	"github.com/davlis/cron/scheduler"
	"github.com/davlis/cron/storage"
	"github.com/labstack/echo"
)

// TODO: research about body parsing layer & validation layer
// https://echo.labstack.com/guide/request

// TODO: Add service layer / processors

func CreateJob(c echo.Context) error {
	log.Println("api/jobs/createJob : Create Job started")

	job := new(jobs.Job)
	if err := c.Bind(&job); err != nil {
		return err
	}

	err := storage.Connection.Collection("jobs").Save(job)

	if err != nil {
		log.Fatalf("%s: %s", "api/jobs/createJob : Failed to save job to database", err)
		return c.JSON(http.StatusBadRequest, "Error occured while saving to database") // TODO serialized error
	}

	go scheduler.WithJob(job)

	log.Printf("api/jobs/createJob : Create Job finished")

	return c.JSON(http.StatusCreated, job)
}
