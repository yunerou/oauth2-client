package example

import (
	"context"
	"strings"
)

func (j *jobsCmd) CheckPanicRecovery(ctx context.Context, args ...string) (msg string, err error) {
	r := strings.Join(args, " ")
	panic(r)
}
