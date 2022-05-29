package er

import (
	"github.com/blackhorseya/gocommon/pkg/er"
	"net/http"
)

var (
	// ErrListCars means list car is failure
	ErrListCars = er.NewAPPError(http.StatusInternalServerError, 50000, "list car is failure")

	// ErrLogin means login is failure
	ErrLogin = er.NewAPPError(http.StatusInternalServerError, 50001, "login is failure")

	// ErrQueryArrears means query arrears is failure
	ErrQueryArrears = er.NewAPPError(http.StatusInternalServerError, 50002, "query arrears is failure")

	// ErrListBooking means list bookings is failure
	ErrListBooking = er.NewAPPError(http.StatusInternalServerError, 50003, "list bookings is failure")

	// ErrBook means book a car is failure
	ErrBook = er.NewAPPError(http.StatusInternalServerError, 50004, "book a car is failure")

	// ErrCancelBooking means cancel a booking is failure
	ErrCancelBooking = er.NewAPPError(http.StatusInternalServerError, 50005, "cancel a booking is failure")

	// ErrGetBookingByID means get booking by id is failure
	ErrGetBookingByID = er.NewAPPError(http.StatusInternalServerError, 50006, "get booking by id is failure")
)
