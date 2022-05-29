package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IRepo declare app repo function
type IRepo interface {
	// List serve caller to list all car
	List(ctx contextx.Contextx) (cars []*pb.Car, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, NewOptions)
