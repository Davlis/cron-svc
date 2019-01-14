package storage

import (
	"log"

	"github.com/davlis/cron/config"
	"github.com/go-bongo/bongo"
)

var Connection *bongo.Connection

func Connect() error {
	log.Println("storage/database : Database connection starting")

	config := &bongo.Config{
		ConnectionString: config.Cfg.MongoUrl,
		Database:         config.Cfg.MongoDatabaseName,
	}

	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	Connection = connection

	log.Println("storage/database : Database connection established")

	return nil
}
