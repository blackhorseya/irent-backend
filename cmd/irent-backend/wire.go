//go:build wireinject

package main

import (
	"github.com/blackhorseya/irent/internal/app/irent"
	"github.com/blackhorseya/irent/internal/app/irent/apis"
	"github.com/blackhorseya/irent/internal/app/irent/biz"
	"github.com/blackhorseya/irent/internal/pkg/app"
	"github.com/blackhorseya/irent/internal/pkg/entity/config"
	"github.com/blackhorseya/irent/internal/pkg/infra/log"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	irent.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
