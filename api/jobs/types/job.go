package jobs

import "github.com/go-bongo/bongo"

type Job struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string `json:"name" validate:"required"`
	Interval           uint64 `json:"interval" validate:"required"`
	Data               string `json:"data" validate:"required"`
}

// TODO: move me somewhere
