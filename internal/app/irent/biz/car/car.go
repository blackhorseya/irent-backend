package car

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/car/repo"
	"github.com/blackhorseya/irent/internal/pkg/base/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IBiz declare user biz service function
type IBiz interface {
	// NearTopN serve caller to list closer current location
	NearTopN(ctx contextx.Contextx, top int, latitude, longitude float64) (cars []*pb.Car, total int, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
