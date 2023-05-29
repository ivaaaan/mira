package markdown

import (
	"bytes"

	"github.com/ivaaaan/mira/task"
	bf "github.com/russross/blackfriday/v2"
)

func NewParser() *markdownParser {
	md := bf.New(
		bf.WithExtensions(bf.CommonExtensions),
	)
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
		var curr *task.Task
		if len(tasks) > 0 {
			curr = tasks[len(tasks)-1]
		}

		switch n.Type {
		case bf.Heading:
			if !entering {
				return bf.GoToNext
			}

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
			return bf.SkipChildren
		case bf.Link:
			if entering {
				curr.Description.Write(n.Destination)
				return bf.SkipChildren
			}
		case bf.List:
			if entering {
				return bf.GoToNext
			}
		case bf.Item:
			if entering {
				curr.Description.Write([]byte("\n- "))
				return bf.GoToNext
			}
		default:
			if curr != nil {
				if entering {
					curr.Description.Write(n.Literal)
				} else {
					curr.Description.Write([]byte("\n"))
				}
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
