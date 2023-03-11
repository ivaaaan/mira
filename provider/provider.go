package provider

import (
	"context"

	"github.com/ivaaaan/mira/task"
)

type Provider interface {
	Push(ctx context.Context, t *task.Task) error
}
