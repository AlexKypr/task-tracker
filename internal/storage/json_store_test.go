package storage

import (
	"os"
	"testing"
	"time"

	"github.com/AlexKypr/task-tracker/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTempFile(t *testing.T) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "test_tasks_*.json")
	assert.NoError(t, err)
	defer tmpFile.Close()
	return tmpFile.Name()
}

func cleanTempFile(t *testing.T, filePath string) {
	t.Helper()
	err := os.Remove(filePath)
	assert.NoError(t, err)
}

func createTestTask(id string, description string, status models.TaskStatus) *models.Task {
	now := time.Now()
	return &models.Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func TestJSONStore_WriteAndReadTasks(t *testing.T) {
	filePath := setupTempFile(t)
	defer cleanTempFile(t, filePath)

	store := NewJSONStore(filePath)

	// Write tasks
	tasks := []*models.Task{
		createTestTask(uuid.NewString(), "Task 1", models.Todo),
		createTestTask(uuid.NewString(), "Task 2", models.InProgress),
	}
	err := store.WriteTasks(tasks)
	assert.NoError(t, err)

	// Read tasks
	readTasks, err := store.ReadTasks()
	assert.NoError(t, err)
	assert.Len(t, readTasks, len(readTasks))
	assert.Equal(t, tasks[0].Description, readTasks[0].Description)
	assert.Equal(t, tasks[1].Status, readTasks[1].Status)
}

func TestJSONStore_Add(t *testing.T) {
	filePath := setupTempFile(t)
	defer cleanTempFile(t, filePath)

	store := NewJSONStore(filePath)

	task := createTestTask(uuid.NewString(), "New Task", models.Todo)
	err := store.Add(task)
	assert.NoError(t, err)

	readTasks, err := store.ReadTasks()
	assert.NoError(t, err)
	assert.Len(t, readTasks, 1)
	assert.Equal(t, task.Description, readTasks[0].Description)
}

func TestJSONStore_List(t *testing.T) {
	filePath := setupTempFile(t)
	defer cleanTempFile(t, filePath)

	store := NewJSONStore(filePath)

	// Write tasks
	tasks := []*models.Task{
		createTestTask(uuid.NewString(), "Task 1", models.Todo),
		createTestTask(uuid.NewString(), "Task 2", models.Done),
	}
	err := store.WriteTasks(tasks)
	assert.NoError(t, err)

	// List all tasks
	allTasks, err := store.List(nil)
	assert.NoError(t, err)
	assert.Len(t, allTasks, len(allTasks))

	// Filter tasks by status
	filter := models.Todo
	todoTasks, err := store.List(&filter)
	assert.NoError(t, err)
	assert.Len(t, todoTasks, len(todoTasks))
	assert.Equal(t, "Task 1", todoTasks[0].Description)
}

func TestJSONStore_Delete(t *testing.T) {
	filePath := setupTempFile(t)
	defer cleanTempFile(t, filePath)

	store := NewJSONStore(filePath)

	// Write tasks
	tasks := []*models.Task{
		createTestTask("1", "Task 1", models.Todo),
		createTestTask("2", "Task 2", models.Done),
	}
	err := store.WriteTasks(tasks)
	assert.NoError(t, err)

	// Delete a task
	err = store.Delete("1")
	assert.NoError(t, err)

	// Verify task was deleted
	readTasks, err := store.List(nil)
	assert.NoError(t, err)
	assert.Len(t, readTasks, len(readTasks))
	assert.Equal(t, "Task 2", readTasks[0].Description)
}

func TestJSONStore_Update(t *testing.T) {
	filePath := setupTempFile(t)
	defer cleanTempFile(t, filePath)

	store := NewJSONStore(filePath)

	// Write tasks
	tasks := []*models.Task{
		createTestTask("1", "Task 1", models.Todo),
	}
	err := store.WriteTasks(tasks)
	assert.NoError(t, err)

	// Update a task
	updatedTask := createTestTask("1", "Updated Task", models.Done)
	err = store.Update("1", updatedTask)
	assert.NoError(t, err)

	// Verify task was updated
	readTasks, err := store.List(nil)
	assert.NoError(t, err)
	assert.Len(t, readTasks, 1)
	assert.Equal(t, "Updated Task", readTasks[0].Description)
	assert.Equal(t, models.Done, readTasks[0].Status)
}
