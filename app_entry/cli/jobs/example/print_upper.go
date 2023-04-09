package example

import (
	"context"
	"strings"
)

func (j *jobsCmd) PrintUpper(ctx context.Context, args ...string) (msg string, err error) {
	r := strings.Join(args, " ")
	return strings.ToUpper(r), nil
}
