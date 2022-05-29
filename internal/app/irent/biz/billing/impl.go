package billing

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing/repo"
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
		logger: logger.With(zap.String("type", "BillingBiz")),
		repo:   repo,
	}
}

func (i *impl) GetArrears(ctx contextx.Contextx, user *pb.Profile) (info *pb.Arrears, err error) {
	if len(user.Id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.Any("user", user))
		return nil, er.ErrMissingID
	}

	if len(user.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Any("user", user))
		return nil, er.ErrMissingToken
	}

	ret, err := i.repo.QueryArrears(ctx, user)
	if err != nil {
		i.logger.Error(er.ErrQueryArrears.Error(), zap.Any("user", user))
		return nil, er.ErrQueryArrears
	}

	return ret, nil
}
