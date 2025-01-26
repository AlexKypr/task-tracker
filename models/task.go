package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskProgress string

type Task struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(description string) *Task {
	now := time.Now()
	return &Task{
		ID:          uuid.NewString(),
		Description: description,
		Status:      Todo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Task) hasValidID() bool {
	_, err := uuid.Parse(t.ID)
	return err == nil
}
