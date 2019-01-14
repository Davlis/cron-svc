package queue_manager

import (
	"log"

	"github.com/davlis/cron/config"
	"github.com/streadway/amqp"
)

var Channel *amqp.Channel

func Connect() error {
	c := make(chan *amqp.Error)

	go func() {
		err := <-c
		log.Println("lib/queue_manager/channel : Reconnected " + err.Error())
		Connect()
	}()

	conn, err := amqp.Dial(config.Cfg.AmqpUrl)
	if err != nil {
		panic("lib/queue_manager/channel : Failed to connect message broker")
	}

	log.Println("lib/queue_manager/channel : Connected to AMQP")

	ch, err := conn.Channel()
	if err != nil {
		panic("lib/queue_manager/channel : Failed to open a message broker channel")
	}

	log.Println("lib/queue_manager/channel : AMQP channel initializing")

	conn.NotifyClose(c)
	ch.NotifyClose(c)

	Channel = ch

	log.Println("lib/queue_manager/channel : AMQP channel initialized")

	return nil
}
