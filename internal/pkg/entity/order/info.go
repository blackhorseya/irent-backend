package order

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
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
