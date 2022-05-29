package billing

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing/repo"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IBiz declare arrears service function
type IBiz interface {
	// GetArrears serve caller to given user then get user's arrears information
	GetArrears(ctx contextx.Contextx, user *pb.Profile) (info *pb.Arrears, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
