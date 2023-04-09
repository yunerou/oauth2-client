package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/yunerou/oauth2-client/common/comctx"
)

type JobRunCmd func(ctx context.Context, args ...string) (msg string, err error)

// Start ...run a cli to exec. "./app --mode=job --job=print_upper print this"
func Start(jobName string, args ...string) {

	// new context
	newCtx := context.Background()
	newCtx = context.WithValue(newCtx, comctx.RequestTraceIDKey, "xxx")

	// exec job
	if execFunc, ok := getJobMapping()[jobName]; ok {
		msg, err := execFunc(newCtx, args...)
		if err != nil {
			handleError(err)
			return
		}
		fmt.Println(msg)
		return
	}
	panic(fmt.Sprintf("[%s] job has not exist", jobName))
}

func handleError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
