//go:build wireinject
// +build wireinject

package cli

import (
	"github.com/google/wire"
)

func initJobsCollection() *jobsCollection {

	wire.Build(IteractorCollection)

	return &jobsCollection{}
}
