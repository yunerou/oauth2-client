package appentry

import (
	"github.com/google/wire"

	cliJobExample "github.com/yunerou/oauth2-client/app_entry/cli/jobs/example"

	commonpages "github.com/yunerou/oauth2-client/usecases/delivery/common_pages"
)

var IteractorCollection = wire.NewSet(
	cliJobExample.NewJobsCmd,
	commonpages.NewCommonPages,
)
