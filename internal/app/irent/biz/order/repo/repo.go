package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/wire"
)

// IRepo declare booking repo service function
type IRepo interface {
	// QueryBookings serve caller to query all bookings
	QueryBookings(ctx contextx.Contextx, user *pb.Profile) (orders []*pb.OrderInfo, err error)

	// Book serve caller to book a car
	Book(ctx contextx.Contextx, id, projID string, user *pb.Profile) (info *pb.Booking, err error)

	// CancelBooking serve caller to cancel a booking by id
	CancelBooking(ctx contextx.Contextx, id string, user *pb.Profile) (err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, NewOptions)
