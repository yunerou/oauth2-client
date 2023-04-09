package cli

import (
	appentry "github.com/yunerou/oauth2-client/app_entry"
	"github.com/yunerou/oauth2-client/app_entry/cli/jobs/example"

	"github.com/google/wire"
)

// Make new type for sumary handler interactor
type jobsCollection struct {
	jobsExample example.JobsCmd
}

func NewJobsCollection(
	jobsExample example.JobsCmd,
) *jobsCollection {
	return &jobsCollection{
		jobsExample,
	}
}

var IteractorCollection = wire.NewSet(
	appentry.IteractorCollection,
	NewJobsCollection,
)
