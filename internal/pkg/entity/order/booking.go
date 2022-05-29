package order

import (
	"time"
)

// Booking declare booking information
type Booking struct {
	No         string    `json:"no"`
	LastPickAt time.Time `json:"last_pick_at"`
}
