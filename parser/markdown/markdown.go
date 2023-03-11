package markdown

import (
	"github.com/ivaaaan/mira/task"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func NewParser() *markdownParser {
	defaultParser := goldmark.DefaultParser()
	defaultParser.AddOptions(parser.WithAttribute())
	return &markdownParser{
		p: defaultParser,
	}
}

type markdownParser struct {
	p parser.Parser
}

func (p *markdownParser) Parse(b []byte) (*task.Task, error) {
	node := p.p.Parse(text.NewReader(b))

	var tasks []*task.Task
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering == false {
			return ast.WalkContinue, nil
		}

		switch n := n.(type) {
		case *ast.Heading:
			t := string(n.Text(b))
			newTask := &task.Task{
				Title: t,
				Level: n.Level,
			}

			tasks = append(tasks, newTask)
		case *ast.Paragraph:
			tasks[len(tasks)-1].Description = string(n.Text(b))
		}

		return ast.WalkContinue, nil
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
