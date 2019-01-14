package scheduler

import (
	"github.com/davlis/cron/lib/events"
	"github.com/jasonlvhit/gocron"
)

func ScheduleAll(tasks []SchTask, sh *gocron.Scheduler) {
	var i int
	var length = len(tasks)
	for i = 0; i < length; i++ {
		Schedule(tasks[i], sh)
	}
}

func Schedule(task SchTask, sh *gocron.Scheduler) {
	sh.Every(task.Interval).Seconds().Do(events.CheckPendingTasks, task.Name, task.Data)
}
