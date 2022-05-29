package user

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/user/repo"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IBiz declare user biz service function
type IBiz interface {
	// GetByID serve caller to given user's id to get user profile
	GetByID(ctx contextx.Contextx, id string) (info *pb.Profile, err error)

	// GetByAccessToken serve caller to given user's access token to get user profile
	GetByAccessToken(ctx contextx.Contextx, token string) (info *pb.Profile, err error)

	// Login serve caller to given id and password then login the system
	Login(ctx contextx.Contextx, id, password string) (info *pb.Profile, err error)

	// RefreshToken serve caller to refresh user's access token
	RefreshToken(ctx contextx.Contextx, user *pb.Profile) (info *pb.Profile, err error)

	// Logout serve caller to logout the system
	Logout(ctx contextx.Contextx, user *pb.Profile) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
