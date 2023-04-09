//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
)

func initRoutesCollection() *routesCollection {

	wire.Build(IteractorCollection)

	return &routesCollection{}
}
