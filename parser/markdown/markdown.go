package markdown

import (
	"strings"

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
			tasks = append(tasks, task.NewTask(string(textNode.Literal), n.Level))
			return bf.SkipChildren
		case bf.Link:
			if entering {
				curr.WriteDescription(string(n.Destination))
				return bf.SkipChildren
			}
		case bf.Item:
			if entering {
				curr.WriteDescription("\n- ")
				return bf.GoToNext
			}
		case bf.CodeBlock:
			curr.WriteDescription("{code:" + string(n.Info) + "}\n")
			curr.WriteDescription(string(n.Literal))
			curr.WriteDescription("{code}")
		case bf.Paragraph:
			// Check if it is an empty paragraph, and it's not a part of a list
			if entering && len(strings.TrimSpace(string(n.Literal))) == 0 && n.Prev != nil && n.Prev.Type != bf.Item {
				curr.WriteDescription("\n\n")
			}
		default:
			if curr != nil {
				if entering {
					curr.WriteDescription(string(n.Literal))
				} else {
					curr.WriteDescription("\n")
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
