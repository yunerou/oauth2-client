package server

import (
	"github.com/google/wire"
	appentry "github.com/yunerou/oauth2-client/app_entry"
	commonpages "github.com/yunerou/oauth2-client/usecases/delivery/common_pages"
)

// Make new type for sumary handler interactor
type routesCollection struct {
	commonpages commonpages.CommonPages
}

func NewroutesCollection(
	commonpages commonpages.CommonPages,
) *routesCollection {
	return &routesCollection{
		commonpages,
	}
}

var IteractorCollection = wire.NewSet(
	appentry.IteractorCollection,
	NewroutesCollection,
)
