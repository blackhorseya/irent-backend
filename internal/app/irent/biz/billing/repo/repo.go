package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IRepo declare repo service function
//go:generate mockery --name=IRepo
type IRepo interface {
	// QueryArrears serve caller to query arrears
	QueryArrears(ctx contextx.Contextx, user *user.Profile) (info *user.Arrears, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, NewOptions)
