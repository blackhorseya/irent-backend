package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IRepo declare repo service function
type IRepo interface {
	// QueryArrears serve caller to query arrears
	QueryArrears(ctx contextx.Contextx, user *pb.Profile) (info *pb.Arrears, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, NewOptions)
