package core

import "time"

type Task struct {
	Id          uint32
	Name        string
	Description string
	Priority    string
	Assignee    uint32
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
