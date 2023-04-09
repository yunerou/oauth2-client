package example

import (
	"context"

	)

type JobsCmd interface {
	PrintUpper(ctx context.Context, args ...string) (msg string, err error)
	CheckPanicRecovery(ctx context.Context, args ...string) (msg string, err error)
}

type jobsCmd struct {
	
}

func NewJobsCmd() JobsCmd {
	return &jobsCmd{
	}
}
