//go:build wireinject

package repo

import (
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIRepo serve caller to create an IRepo
func CreateIRepo(o *Options, client restclient.HTTPClient) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
