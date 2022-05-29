package car

import (
	"sort"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/car/repo"
	"github.com/blackhorseya/irent/internal/pkg/base/distance"
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
	return &impl{logger: logger, repo: repo}
}

func (i *impl) NearTopN(ctx contextx.Contextx, top int, latitude, longitude float64) (cars []*pb.Car, total int, err error) {
	if top <= 0 {
		i.logger.Error(er.ErrInvalidN.Error(), zap.Int("top", top))
		return nil, 0, er.ErrInvalidN
	}

	cars, err = i.repo.List(ctx)
	if err != nil {
		i.logger.Error(er.ErrListCars.Error(), zap.Error(err), zap.Int("top", top))
		return nil, 0, er.ErrListCars
	}
	if len(cars) == 0 {
		i.logger.Error(er.ErrCarNotExists.Error(), zap.Int("top", top))
		return nil, 0, er.ErrCarNotExists
	}

	for _, c := range cars {
		c.Distance = distance.CalcWithGeo(latitude, longitude, c.Latitude, c.Longitude, "K")
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Distance < cars[j].Distance
	})

	return cars[:top], len(cars), nil
}
