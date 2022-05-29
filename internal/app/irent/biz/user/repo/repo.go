package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IRepo declare repo service function
type IRepo interface {
	// Login serve caller to log in the system
	Login(ctx contextx.Contextx, id, password string) (info *pb.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, NewOptions)
