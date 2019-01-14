package scheduler

import (
	"log"

	jobs "github.com/davlis/cron/api/jobs/types"
	"github.com/davlis/cron/storage"
	"github.com/globalsign/mgo/bson"
	"github.com/jasonlvhit/gocron"
)

type SchTask struct {
	Interval uint64
	Data     string
	Name     string
}

func Initialize() {
	log.Println("lib/scheduler/initialize#Initialize : starting")

	sh := _init()
	_prepareTaskQueue(sh)
	<-sh.Start()

	log.Println("lib/scheduler/initialize#Initialize : finished")
}

func WithJob(job *jobs.Job) {
	sh := _init()
	task := _serializeJobsToSchTask(job)
	Schedule(task, sh)
	<-sh.Start()
}

func _init() *gocron.Scheduler {
	return gocron.NewScheduler()
}

// TODO: Inject tasks via function parameter
// FIXME: Dependency inversion
func _prepareTaskQueue(sh *gocron.Scheduler) {
	job := &jobs.Job{}
	arrayOfJobs := []SchTask{}

	results := storage.Connection.Collection("jobs").Find(bson.M{})

	for results.Next(job) {
		arrayOfJobs = append(arrayOfJobs, _serializeJobsToSchTask(job))
	}

	ScheduleAll(arrayOfJobs, sh)
}

// TODO: Move me closer to domain logic
func _serializeJobsToSchTask(job *jobs.Job) SchTask {
	return SchTask{Interval: job.Interval, Name: job.Name, Data: job.Data}
}
