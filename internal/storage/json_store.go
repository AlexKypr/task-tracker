package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/AlexKypr/task-tracker/models"
)

type JSONStore struct {
	filePath string
	mu       sync.RWMutex
}

func NewJSONStore(filePath string) *JSONStore {
	return &JSONStore{
		filePath: filePath,
	}
}

// WriteTasks writes all the tasks to json file
// @WARN: File will either be created or truncated, therefore previous data will be lost if old data is not prepended
func (s *JSONStore) WriteTasks(tasks []*models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	f, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			// Log the error if file closing fails
			err = closeErr
		}
	}()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")

	// Write each task as a separate JSON object
	for _, task := range tasks {
		if err := enc.Encode(task); err != nil {
			return err
		}
	}

	return nil
}

// ReadTasks read all the tasks from json file and returnes them
// @WARN: All data will be loaded on memory, so for large files use with caution
func (s *JSONStore) ReadTasks() ([]*models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	dec := json.NewDecoder(bytes.NewReader(data))
	for {
		var task models.Task
		if err := dec.Decode(&task); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (s *JSONStore) Add(task *models.Task) error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, task)
	return s.WriteTasks(tasks)
}

func (s *JSONStore) List(status *models.TaskStatus) ([]*models.Task, error) {
	tasks, err := s.ReadTasks()
	if err != nil {
		return nil, err
	}

	if status == nil {
		return tasks, nil
	}

	var filtered []*models.Task
	for _, task := range tasks {
		if task.Status == *status {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func (s *JSONStore) Delete(id string) error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return err
	}

	var updated []*models.Task
	found := false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		updated = append(updated, task)
	}

	if !found {
		return fmt.Errorf("task: %v not found", id)
	}

	return s.WriteTasks(updated)
}

func (s *JSONStore) Update(id string, updatedTask *models.Task) error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return err
	}

	updated := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("task: %v not found", id)
	}

	return s.WriteTasks(tasks)
}
