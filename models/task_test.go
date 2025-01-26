package models

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTask(t *testing.T) {
	description := "Test Task Description"
	task := NewTask(description)

	// Assert that task is not nil pointer
	if task == nil {
		t.Fatalf("NewTask() returned nil, expected a valid task pointer")
	}

	// Check Description has initialized based on input
	if task.Description != description {
		t.Errorf("Expected Description: %s, got: %s", task.Description, description)
	}

	// Check Status that has initialized correctly with default value
	if task.Status != Todo {
		t.Errorf("Expected Status: %s, got: %s", task.Status, Todo)
	}

	// Check ID is valid UUID
	if _, err := uuid.Parse(task.ID); err != nil {
		t.Errorf("Invalid UUID for ID: %s", task.ID)
	}

	// Check createdAt and updatedAt are non zero
	if task.CreatedAt.IsZero() {
		t.Errorf("CreatedAt not set, got: %v", task.CreatedAt)
	}
	if task.UpdatedAt.IsZero() {
		t.Errorf("UpdatedAt not set, got: %v", task.UpdatedAt)
	}
}

func TestHasValidID(t *testing.T) {
	description := "Test Task Description"
	task := NewTask(description)

	if !task.hasValidID() {
		t.Errorf("Expected hasValidID() to return true for a valid UUID, got false")
	}

	task.ID = "random-invalid-id"
	if task.hasValidID() {
		t.Errorf("Expected hasValidID() to return false for an invalid UUID, got true")
	}

	task.ID = ""
	if task.hasValidID() {
		t.Errorf("Expected hasValidID() to return false for an empty UUID, got true")
	}
}
