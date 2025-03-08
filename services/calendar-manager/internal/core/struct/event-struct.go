package core

import "time"

type Event struct {
	Id          uint32
	Name        string
	Description string
	Time        time.Time
	Duration    uint32
	Attendees   []uint32
}
