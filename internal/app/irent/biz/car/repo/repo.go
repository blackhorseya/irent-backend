package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"github.com/google/wire"
)

// IRepo declare app repo function
//go:generate mockery --name=IRepo
type IRepo interface {
	// List serve caller to list all car
	List(ctx contextx.Contextx) (cars []*car.Info, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, NewOptions)
