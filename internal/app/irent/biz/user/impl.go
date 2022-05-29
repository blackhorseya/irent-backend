package user

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/user/repo"
	"github.com/blackhorseya/irent/internal/pkg/base/encrypt"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/pb"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "UserBiz")),
		repo:   repo,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (info *pb.Profile, err error) {
	// todo: 2021-05-11|11:59|doggy|implement me
	panic("implement me")
}

func (i *impl) GetByAccessToken(ctx contextx.Contextx, token string) (info *pb.Profile, err error) {
	// todo: 2021-05-11|11:59|doggy|implement me
	panic("implement me")
}

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *pb.Profile, err error) {
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error())
		return nil, er.ErrMissingID
	}

	if len(password) == 0 {
		i.logger.Error(er.ErrMissingPassword.Error())
		return nil, er.ErrMissingPassword
	}

	ret, err := i.repo.Login(ctx, id, encrypt.EncPWD(password))
	if err != nil {
		i.logger.Error(er.ErrLogin.Error(), zap.Error(err))
		return nil, er.ErrLogin
	}

	return ret, nil
}

func (i *impl) Logout(ctx contextx.Contextx, user *pb.Profile) error {
	// todo: 2021-05-11|11:59|doggy|implement me
	panic("implement me")
}

func (i *impl) RefreshToken(ctx contextx.Contextx, user *pb.Profile) (info *pb.Profile, err error) {
	// todo: 2021-05-11|11:59|doggy|implement me
	panic("implement me")
}
