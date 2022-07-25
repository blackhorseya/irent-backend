package user

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/user/repo"
	"github.com/blackhorseya/irent/internal/pkg/base/encrypt"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
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

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *user.Profile, err error) {
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
