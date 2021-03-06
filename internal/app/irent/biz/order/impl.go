package order

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/order/repo"
	"github.com/blackhorseya/irent/internal/pkg/base/stringutils"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options declare premium configuration
type Options struct {
	People []string `json:"people"`
}

// NewOptions return *Options
func NewOptions(v *viper.Viper) (*Options, error) {
	o := new(Options)

	err := v.UnmarshalKey("premium", &o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

type impl struct {
	o      *Options
	logger *zap.Logger
	repo   repo.IRepo

	premiumBookings map[*user.Profile]*order.Booking
}

// NewImpl serve caller to create an IBiz
func NewImpl(o *Options, logger *zap.Logger, repo repo.IRepo) IBiz {
	return &impl{
		o:               o,
		logger:          logger.With(zap.String("type", "OrderBiz")),
		repo:            repo,
		premiumBookings: make(map[*user.Profile]*order.Booking),
	}
}

func (i *impl) List(ctx contextx.Contextx, start, end int, from *user.Profile) (orders []*order.Info, err error) {
	if start < 0 {
		i.logger.Error(er.ErrInvalidStart.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("from", from))
		return nil, er.ErrInvalidStart
	}

	if end < 0 {
		i.logger.Error(er.ErrInvalidEnd.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("from", from))
		return nil, er.ErrInvalidEnd
	}

	if len(from.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("from", from))
		return nil, er.ErrMissingToken
	}

	ret, err := i.repo.QueryBookings(ctx, from)
	if err != nil {
		i.logger.Error(er.ErrListBooking.Error(), zap.Error(err), zap.Int("start", start), zap.Int("end", end), zap.Any("from", from))
		return nil, er.ErrListBooking
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrBookingNotExists.Error(), zap.Int("start", start), zap.Int("end", end), zap.Any("from", from))
		return nil, er.ErrBookingNotExists
	}

	return ret, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string, from *user.Profile) (info *order.Info, err error) {
	if len(from.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Any("from", from))
		return nil, er.ErrMissingToken
	}

	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.Any("from", from))
		return nil, er.ErrMissingID
	}

	ret, err := i.repo.QueryBookings(ctx, from)
	if err != nil {
		i.logger.Error(er.ErrGetBookingByID.Error(), zap.Any("from", from))
		return nil, er.ErrGetBookingByID
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrBookingNotExists.Error(), zap.Any("from", from))
		return nil, er.ErrBookingNotExists
	}

	for _, o := range ret {
		if o.No == id {
			return o, nil
		}
	}

	return nil, er.ErrBookingNotExists
}

func (i *impl) BookCar(ctx contextx.Contextx, id, projID string, from *user.Profile, circularly bool) (info *order.Booking, err error) {
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.String("projID", projID), zap.Any("from", from))
		return nil, er.ErrMissingID
	}

	if len(projID) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.String("projID", projID), zap.Any("from", from))
		return nil, er.ErrMissingID
	}

	if len(from.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.String("projID", projID), zap.Any("from", from))
		return nil, er.ErrMissingToken
	}

	ret, err := i.repo.Book(ctx, id, projID, from)
	if err != nil {
		i.logger.Error(er.ErrBook.Error(), zap.Error(err), zap.String("projID", projID), zap.Any("from", from))
		return nil, er.ErrBook
	}

	isPremium := stringutils.ContainsInSlice(i.o.People, from.ID)
	if isPremium && circularly {
		i.premiumBookings[from] = ret
	}

	return ret, nil
}

func (i *impl) CancelBooking(ctx contextx.Contextx, id string, from *user.Profile) (err error) {
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.Any("from", from))
		return er.ErrMissingID
	}

	if len(from.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Any("from", from))
		return er.ErrMissingToken
	}

	err = i.repo.CancelBooking(ctx, id, from)
	if err != nil {
		i.logger.Error(er.ErrCancelBooking.Error(), zap.Any("from", from))
		return er.ErrCancelBooking
	}

	for profile := range i.premiumBookings {
		if profile.AccessToken == from.AccessToken {
			delete(i.premiumBookings, profile)
		}
	}

	return nil
}

func (i *impl) ReBookCar(ctx contextx.Contextx, no, id, projID string, from *user.Profile) (info *order.Booking, err error) {
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.Any("from", from))
		return nil, er.ErrMissingID
	}

	if len(projID) == 0 {
		i.logger.Error(er.ErrMissingID.Error(), zap.String("projID", projID), zap.Any("from", from))
		return nil, er.ErrMissingID
	}

	if len(from.AccessToken) == 0 {
		i.logger.Error(er.ErrMissingToken.Error(), zap.Any("from", from))
		return nil, er.ErrMissingToken
	}

	err = i.repo.CancelBooking(ctx, no, from)
	if err != nil {
		i.logger.Error(er.ErrCancelBooking.Error(), zap.Any("from", from))
		return nil, er.ErrCancelBooking
	}

	ret, err := i.repo.Book(ctx, id, projID, from)
	if err != nil {
		i.logger.Error(er.ErrBook.Error(), zap.Error(err))
		return nil, er.ErrBook
	}

	return ret, nil
}

func (i *impl) ListPremiumBookings(ctx contextx.Contextx) (info map[*user.Profile]*order.Booking, err error) {
	return i.premiumBookings, nil
}

func (i *impl) UpdatePremiumBooking(ctx contextx.Contextx, from *user.Profile, booking *order.Booking) (info *order.Booking, err error) {
	i.premiumBookings[from] = booking

	return booking, nil
}
