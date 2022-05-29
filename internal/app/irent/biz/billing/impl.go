package billing

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing/repo"
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
		logger: logger.With(zap.String("type", "BillingBiz")),
		repo:   repo,
	}
}

func (i *impl) GetArrears(ctx contextx.Contextx, from *user.Profile) (info *user.Arrears, err error) {
	if len(from.ID) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.Any("from", from))
		return nil, er.ErrMissingID
	}

	if len(from.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Any("from", from))
		return nil, er.ErrMissingToken
	}

	ret, err := i.repo.QueryArrears(ctx, from)
	if err != nil {
		i.logger.Error(er.ErrQueryArrears.Error(), zap.Any("from", from))
		return nil, er.ErrQueryArrears
	}

	return ret, nil
}
