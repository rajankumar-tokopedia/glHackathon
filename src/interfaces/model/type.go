package model

import "time"

type Model interface {
	InsertSubmission() bool
	GetLeaders(eventId int64, howMany int64) []interface{}
}

type TestCase struct {
	Input []interface{} `json:"input" db:"input"`
	Ouput []interface{} `json:"output" db:"output"`
	TTL   time.Time     `json:"ttl" db:"ttl"`
}
