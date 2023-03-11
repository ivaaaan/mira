package parser

import "github.com/ivaaaan/mira/task"

type Parser interface {
	Parse(b []byte) (*task.Task, error)
}
