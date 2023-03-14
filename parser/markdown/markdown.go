package markdown

import (
	"bytes"

	"github.com/ivaaaan/mira/task"
	bf "github.com/russross/blackfriday/v2"
)

func NewParser() *markdownParser {
	md := bf.New(bf.WithExtensions(bf.CommonExtensions))
	return &markdownParser{
		p: md,
	}
}

type markdownParser struct {
	p *bf.Markdown
}

func (p *markdownParser) Parse(b []byte) (*task.Task, error) {
	node := p.p.Parse(b)

	var tasks []*task.Task
	node.Walk(func(n *bf.Node, entering bool) bf.WalkStatus {
		if entering == false {
			return bf.GoToNext
		}

		switch n.Type {
		case bf.Heading:
			textNode := n.FirstChild
			if textNode == nil {
				return bf.GoToNext
			}
			newTask := &task.Task{
				Title:       string(textNode.Literal),
				Level:       n.Level,
				Description: bytes.NewBuffer([]byte{}),
			}

			tasks = append(tasks, newTask)
		default:
			if len(tasks) > 0 {
				tasks[len(tasks)-1].Description.Write(n.Literal)
			}
		}

		return bf.GoToNext
	})

	root := tasks[0]
	stack := []*task.Task{root}
	for _, t := range tasks[1:] {
		if t.Level > len(stack) {
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, t)
		} else {
			stack = stack[:t.Level-1]
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, t)
		}

		stack = append(stack, t)
	}

	return root, nil
}
