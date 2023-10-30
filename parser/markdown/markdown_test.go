package markdown

import (
	_ "embed"
	"testing"

	"github.com/ivaaaan/mira/task"
	"github.com/stretchr/testify/assert"
)

func newTask(title string, level int, description string, children []*task.Task) *task.Task {
	t := task.NewTask(title, level)
	t.WriteDescription(description)
	t.Children = children

	return t
}

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		expected *task.Task
	}{
		{
			name: "parse hierarchy",
			markdown: `# Level 1
## Level 2

### Level 3

### Level 3

## Level 2`,
			expected: newTask("Level 1", 1, "", []*task.Task{
				newTask("Level 2", 2, "", []*task.Task{
					newTask("Level 3", 3, "", nil),
					newTask("Level 3", 3, "", nil),
				}),
				newTask("Level 2", 2, "", nil),
			}),
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := parser.Parse([]byte(tt.markdown))
			assert.NoError(t, err)
			assert.EqualExportedValues(t, *tt.expected, *task)
		})
	}
}
