package main

import (
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

type Parser interface {
	Parse(b []byte) ([]Task, error)
}

func NewParser() Parser {
	return &MarkdownParser{}
}

type MarkdownParser struct {
}

func (p *MarkdownParser) Parse(b []byte) ([]Task, error) {
	tasks := []Task{}

	node := goldmark.DefaultParser().Parse(text.NewReader(b))

	fmt.Println(node)

	return tasks, nil
}
