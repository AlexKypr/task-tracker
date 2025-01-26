package storage

import "github.com/AlexKypr/task-tracker/models"

type Store interface {
	Add(task *models.Task) error
	List(status *models.TaskStatus) ([]*models.Task, error)
	Delete(id string) error
	Update(id string, task *models.Task) error
}
