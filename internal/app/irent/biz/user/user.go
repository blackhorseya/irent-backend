package user

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/user/repo"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IBiz declare user biz service function
//go:generate mockery --name=IBiz
type IBiz interface {
	// Login serve caller to given id and password then login the system
	Login(ctx contextx.Contextx, id, password string) (info *user.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
