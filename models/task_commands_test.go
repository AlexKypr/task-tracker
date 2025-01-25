package models

import "testing"

func TestNewTaskCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected TaskCommands
		valid    bool
	}{
		{"add", Add, true},
		{"update", Update, true},
		{"delete", Delete, true},
		{"mark-in-progress", MarkInProgress, true},
		{"mark-done", MarkDone, true},
		{"list", List, true},
		{"help", Help, true},
		{"invalid-command", "", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			command, valid := NewTaskCommand(test.input)
			if valid != test.valid {
				t.Errorf("expected valid to be %v, got %v", test.valid, valid)
			}
			if command != test.expected {
				t.Errorf("expected command to be %v, got %v", test.expected, command)
			}
		})
	}
}
