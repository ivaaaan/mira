package task

import "io"

type TaskType int

const (
	Epic TaskType = iota
	Story
	Subtask
)

type Task struct {
	ID          string
	Title       string
	Description io.ReadWriter
	Properties  map[string]string
	Type        TaskType
	Children    []*Task
	Level       int
}

func (t Task) GetDescription() string {
	b, err := io.ReadAll(t.Description)
	if err != nil {
		return ""
	}

	return string(b)
}
