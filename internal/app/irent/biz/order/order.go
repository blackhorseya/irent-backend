package order

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/order/repo"
	"github.com/blackhorseya/irent/internal/pkg/base/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IBiz declare order service function
type IBiz interface {
	// List serve caller to list all orders
	List(ctx contextx.Contextx, start, end int, user *pb.Profile) (orders []*pb.OrderInfo, err error)

	// GetByID serve caller to get an order info by id
	GetByID(ctx contextx.Contextx, id string, user *pb.Profile) (info *pb.OrderInfo, err error)

	// BookCar serve caller to book a car
	BookCar(ctx contextx.Contextx, id, projID string, user *pb.Profile) (info *pb.Booking, err error)

	// CancelBooking serve caller to cancel an order by order's id
	CancelBooking(ctx contextx.Contextx, id string, user *pb.Profile) (err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
