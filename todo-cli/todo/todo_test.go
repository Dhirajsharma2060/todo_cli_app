package todo

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		initials    Todos
		description string
		wantErr     bool
	}{
		{
			name:        "add first todo successfully",
			initials:    Todos{},
			description: "Learns gotest",
			wantErr:     false,
		},
		{
			name:        "error on empty description",
			initials:    Todos{},
			description: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todos := tt.initials
			err := todos.Add(tt.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name        string
		initials    Todos
		description string
		wantOut     string
	}{
		{
			name:        "empty list",
			initials:    Todos{},
			description: "test for empty test",
			wantOut:     "",
		},
		{
			name: "proper list",
			initials: Todos{
				{ID: 1, Description: "buy milk", Completed: true, CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)},
			},
			description: "List is proper",
			wantOut:     "[1] ✓ buy milk (Completed) - Created: 2023-10-01\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			output := os.Stdout
			os.Stdout = w

			test.initials.List()
			w.Close()
			os.Stdout = output
			var buf bytes.Buffer

			buf.ReadFrom(r)

			got := buf.String()
			if got != test.wantOut {
				t.Errorf("List() got = %v, want %v", got, test.wantOut)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		initials Todos
		wantErr  bool
	}{
		{
			name: "deleted todo exist",
			initials: Todos{
				{ID: 1, Description: "buy milk", Completed: false, CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: false,
		},
		{
			name: "deleted todo does not  exist",
			initials: Todos{
				{ID: 1, Description: "buy milk", Completed: false, CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			idToDelete := 1
			if test.wantErr {
				idToDelete = 2
			}
			err := test.initials.Delete(idToDelete)
			if (err != nil) != test.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func StatusChange(t *testing.T) {
	tests := []struct {
		name     string
		initials Todos
		wantErr  bool
	}{
		{
			name: "status is pending",
			initials: Todos{
				{ID: 1, Description: "buy milk", Completed: false, CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: false,
		},
		{
			name: "status is completed",
			initials: Todos{
				{ID: 1, Description: "buy milk", Completed: true, CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			idToChange := 1
			if test.wantErr {
				idToChange = 1
			}
			err := test.initials.StatusChange(idToChange)
			if (err != nil) != test.wantErr {
				t.Errorf("StatusChange() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
