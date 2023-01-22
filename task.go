package main

type TaskType int

const (
	Epic TaskType = iota
	Story
	Subtask
)

type Task struct {
	ID          string
	Title       string
	Description string
	Properties  map[string]string
	Type        TaskType
	Parent      *Task
	Children    []Task
}
