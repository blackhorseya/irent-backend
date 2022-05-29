package order

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/order/repo"
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
		logger: logger.With(zap.String("type", "OrderBiz")),
		repo:   repo,
	}
}

func (i *impl) List(ctx contextx.Contextx, start, end int, user *pb.Profile) (orders []*pb.OrderInfo, err error) {
	if start < 0 {
		i.logger.Error(er.ErrInvalidStart.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("user", user))
		return nil, er.ErrInvalidStart
	}

	if end < 0 {
		i.logger.Error(er.ErrInvalidEnd.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("user", user))
		return nil, er.ErrInvalidEnd
	}

	if len(user.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("user", user))
		return nil, er.ErrMissingToken
	}

	ret, err := i.repo.QueryBookings(ctx, user)
	if err != nil {
		i.logger.Error(er.ErrListBooking.Error(), zap.Error(err), zap.Int("start", start), zap.Int("end", end), zap.Any("user", user))
		return nil, er.ErrListBooking
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrBookingNotExists.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("user", user))
		return nil, er.ErrBookingNotExists
	}

	return ret, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string, user *pb.Profile) (info *pb.OrderInfo, err error) {
	if len(user.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Any("user", user))
		return nil, er.ErrMissingToken
	}

	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.Any("user", user))
		return nil, er.ErrMissingID
	}

	ret, err := i.repo.QueryBookings(ctx, user)
	if err != nil {
		i.logger.Error(er.ErrGetBookingByID.Error(), zap.Any("user", user))
		return nil, er.ErrGetBookingByID
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrBookingNotExists.Error(), zap.Any("user", user))
		return nil, er.ErrBookingNotExists
	}

	for _, o := range ret {
		if o.No == id {
			return o, nil
		}
	}

	return nil, er.ErrBookingNotExists
}

func (i *impl) BookCar(ctx contextx.Contextx, id, projID string, user *pb.Profile) (info *pb.Booking, err error) {
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.String("projID", projID), zap.Any("user", user))
		return nil, er.ErrMissingID
	}

	if len(projID) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.String("projID", projID), zap.Any("user", user))
		return nil, er.ErrMissingID
	}

	if len(user.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.String("projID", projID), zap.Any("user", user))
		return nil, er.ErrMissingToken
	}

	ret, err := i.repo.Book(ctx, id, projID, user)
	if err != nil {
		i.logger.Error(er.ErrBook.Error(), zap.Error(err), zap.String("projID", projID), zap.Any("user", user))
		return nil, er.ErrBook
	}

	return ret, nil
}

func (i *impl) CancelBooking(ctx contextx.Contextx, id string, user *pb.Profile) (err error) {
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.Any("user", user))
		return er.ErrMissingID
	}

	if len(user.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Any("user", user))
		return er.ErrMissingToken
	}

	err = i.repo.CancelBooking(ctx, id, user)
	if err != nil {
		i.logger.Error(er.ErrCancelBooking.Error(), zap.Any("user", user))
		return er.ErrCancelBooking
	}

	return nil
}
