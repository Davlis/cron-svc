package main

import (
	"github.com/davlis/cron/config"
	"github.com/davlis/cron/lib/queue_manager"
	"github.com/davlis/cron/scheduler"
	"github.com/davlis/cron/server"
	"github.com/davlis/cron/storage"
)

func main() {
	config.Load()

	storage.Connect()
	queue_manager.Connect()

	go scheduler.Initialize()

	server.CreateServer()
}
