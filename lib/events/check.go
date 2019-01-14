package events

import (
	"log"

	"github.com/davlis/cron/lib/queue_manager"
	"github.com/streadway/amqp"
)

func CheckPendingTasks(eventName string, data string) {
	chann := queue_manager.Channel

	err := chann.Publish("", eventName, true, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(data),
	})

	if err != nil {
		log.Fatalf("%s: %s", "failed to publish "+eventName+" message", err)
		return
	}

	log.Println(eventName + " event message sent with data : " + data)
}
