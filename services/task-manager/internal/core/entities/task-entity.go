package core

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"date"`
}

func NewTask(name string) Task {
	return Task{
		Id:        uuid.New().ID(),
		Name:      name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
