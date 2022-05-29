package car

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"sort"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/car/repo"
	"github.com/blackhorseya/irent/internal/pkg/base/distance"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo) IBiz {
	return &impl{logger: logger, repo: repo}
}

func (i *impl) NearTopN(ctx contextx.Contextx, top int, latitude, longitude float64) (infos []*car.Info, total int, err error) {
	if top <= 0 {
		i.logger.Error(er.ErrInvalidN.Error(), zap.Int("top", top))
		return nil, 0, er.ErrInvalidN
	}

	ret, err := i.repo.List(ctx)
	if err != nil {
		i.logger.Error(er.ErrListCars.Error(), zap.Error(err), zap.Int("top", top))
		return nil, 0, er.ErrListCars
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrCarNotExists.Error(), zap.Int("top", top))
		return nil, 0, er.ErrCarNotExists
	}

	for _, target := range ret {
		target.Distance = distance.CalcWithGeo(latitude, longitude, target.Latitude, target.Longitude, "K")
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Distance < ret[j].Distance
	})

	return ret[:top], len(ret), nil
}
