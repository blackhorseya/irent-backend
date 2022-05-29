//go:build wireinject

package main

import (
	"github.com/blackhorseya/gocommon/pkg/config"
	"github.com/blackhorseya/gocommon/pkg/log"
	"github.com/blackhorseya/irent/internal/app/irent"
	"github.com/blackhorseya/irent/internal/app/irent/api/restful"
	"github.com/blackhorseya/irent/internal/app/irent/biz"
	"github.com/blackhorseya/irent/internal/pkg/app"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	http.ProviderSet,
	irent.ProviderSet,
	restful.ProviderSet,
	biz.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
