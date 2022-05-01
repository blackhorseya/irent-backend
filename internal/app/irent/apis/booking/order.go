package booking

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare booking api handler function
type IHandler interface {
	// ListBookings serve caller to list all bookings
	ListBookings(c *gin.Context)

	// GetBookingByID serve caller to get a booking by id
	GetBookingByID(c *gin.Context)

	// Book serve caller to book a car
	Book(c *gin.Context)

	// CancelBooking serve caller to cancel a booking by id
	CancelBooking(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
