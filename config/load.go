package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	MongoUrl          string `env:"MONGO_URL,required"`
	MongoDatabaseName string `env:"MONGO_DATABASE_NAME,required"`
	AmqpUrl           string `env:"AMQP_URL",required"`
}

var Cfg Config

func Load() Config {
	cfg := Config{}
	err := env.Parse(&cfg)

	if err == nil {
		log.Printf("config/load : Environment variables loaded from shell")
		Cfg = cfg
		return Cfg
	}

	godotenv.Load()

	err = env.Parse(&cfg)
	if err == nil {
		log.Printf("config/load : Environment variables loaded from .env file")
		Cfg = cfg
		return Cfg
	}

	if err != nil {
		log.Printf("config/load : Failed to load environment variables")
		panic(fmt.Sprintf("%+v\n", err))
	}

	return Config{}
}
