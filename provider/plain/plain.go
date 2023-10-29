package plain

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/ivaaaan/mira/task"
)

func NewProvider(out io.Writer) *plainProvider {
	return &plainProvider{out: out}
}

type plainProvider struct {
	out io.Writer
}

func writeTask(w io.Writer, t *task.Task) {
	titleInt := strings.Repeat("-", t.Level)

	descLines := strings.Split(t.Description(), "\n")
	for i, d := range descLines {
		descLines[i] = strings.Repeat(" ", t.Level) + d
	}
	descInt := strings.Join(descLines, "\n")
	fmt.Fprintf(w, "%s %s%s\n", titleInt, t.Title, descInt)
	for _, c := range t.Children {
		writeTask(w, c)
	}
}

func (p *plainProvider) Push(_ context.Context, t *task.Task) error {
	writeTask(p.out, t)
	return nil
}
