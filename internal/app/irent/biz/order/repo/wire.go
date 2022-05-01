//go:build wireinject

package repo

import (
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIRepo serve caller to create an IRepo
func CreateIRepo(o *Options) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
