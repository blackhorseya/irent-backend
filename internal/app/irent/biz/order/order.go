package order

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/order/repo"
	"github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IBiz declare order service function
//go:generate mockery --name=IBiz
type IBiz interface {
	// List serve caller to list all orders
	List(ctx contextx.Contextx, start, end int, from *user.Profile) (orders []*order.Info, err error)

	// GetByID serve caller to get an order info by id
	GetByID(ctx contextx.Contextx, id string, from *user.Profile) (info *order.Info, err error)

	// BookCar serve caller to book a car
	BookCar(ctx contextx.Contextx, id, projID string, from *user.Profile) (info *order.Booking, err error)

	// CancelBooking serve caller to cancel an order by order's id
	CancelBooking(ctx contextx.Contextx, id string, from *user.Profile) (err error)

	// ReBookCar serve caller to rebook car
	ReBookCar(ctx contextx.Contextx, id, projID string, from *user.Profile) (info *order.Booking, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
