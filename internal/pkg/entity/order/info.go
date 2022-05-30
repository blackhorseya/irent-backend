package order

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"github.com/blackhorseya/irent/pb"
	"time"
)

// Info declare order information
type Info struct {
	No         string    `json:"no"`
	Car        *car.Info `json:"car"`
	StartedAt  time.Time `json:"started_at"`
	EndAt      time.Time `json:"end_at"`
	StopPickAt time.Time `json:"stop_pick_at"`
}

// NewOrderInfoResponse return *pb.OrderInfo
func NewOrderInfoResponse(from *Info) *pb.OrderInfo {
	return &pb.OrderInfo{
		No:           from.No,
		CarId:        from.Car.ID,
		CarLatitude:  from.Car.Latitude,
		CarLongitude: from.Car.Longitude,
		StartAt:      from.StartedAt.UTC().Format(time.RFC3339),
		EndAt:        from.EndAt.UTC().Format(time.RFC3339),
		StopPickAt:   from.StopPickAt.UTC().Format(time.RFC3339),
	}
}
