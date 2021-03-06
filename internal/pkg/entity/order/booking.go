package order

import (
	"github.com/blackhorseya/irent/pb"
	"time"
)

// Booking declare booking information
type Booking struct {
	No         string    `json:"no"`
	LastPickAt time.Time `json:"last_pick_at"`

	CarID  string `json:"car_id"`
	ProjID string `json:"proj_id"`
}

// NewBookingResponse return *pb.Booking
func NewBookingResponse(from *Booking) *pb.Booking {
	return &pb.Booking{
		No:         from.No,
		LastPickAt: from.LastPickAt.UTC().Format(time.RFC3339),
	}
}
