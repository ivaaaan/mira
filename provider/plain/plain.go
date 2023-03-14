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
	descInt := strings.Repeat(" ", t.Level)
	fmt.Fprintf(w, "%s %s\n%s %s\n", titleInt, t.Title, descInt, t.GetDescription())
	for _, c := range t.Children {
		writeTask(w, c)
	}
}

func (p *plainProvider) Push(_ context.Context, t *task.Task) error {
	writeTask(p.out, t)
	return nil
}
