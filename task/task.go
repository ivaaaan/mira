package task

import (
	"bytes"
	"io"
)

type TaskType int

const (
	Epic TaskType = iota
	Story
	Subtask
)

type Task struct {
	ID       string
	Title    string
	Type     TaskType
	Children []*Task
	Level    int

	description io.ReadWriter
}

func NewTask(title string, l int) *Task {
	return &Task{
		Title:       title,
		Level:       l,
		description: bytes.NewBuffer([]byte{}),
	}
}

// Append string to a current description.
// If buffer does not exist, it will create a new buffer
func (t Task) WriteDescription(s string) error {
	if t.description == nil {
		t.description = bytes.NewBuffer([]byte{})
	}

	_, err := t.description.Write([]byte(s))
	return err
}

// Reads all bytes from the buffer and returns a string
func (t Task) Description() string {
	b, err := io.ReadAll(t.description)
	if err != nil {
		return ""
	}

	return string(b)
}
