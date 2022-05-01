package er

import "net/http"

var (
	// ErrRateLimit means too many requests
	ErrRateLimit = newAPPError(http.StatusTooManyRequests, 42900, "too many requests")

	// ErrMissingToken means missing token in header
	ErrMissingToken = newAPPError(http.StatusUnauthorized, 40100, "missing token")

	// ErrAuthHeaderFormat means must provide Authorization header with format `Bearer {token}`
	ErrAuthHeaderFormat = newAPPError(http.StatusUnauthorized, 40101, "Must provide Authorization header with format `Bearer {token}`")
)

var (
	// ErrInvalidStart means given start is invalid
	ErrInvalidStart = newAPPError(http.StatusBadRequest, 40000, "start is invalid")

	// ErrInvalidEnd means given end is invalid
	ErrInvalidEnd = newAPPError(http.StatusBadRequest, 40001, "end is invalid")

	// ErrInvalidN means given N is invalid
	ErrInvalidN = newAPPError(http.StatusBadRequest, 40002, "N is invalid")

	// ErrInvalidLatitude means given latitude is invalid
	ErrInvalidLatitude = newAPPError(http.StatusBadRequest, 40003, "latitude is invalid")

	// ErrInvalidLongitude means given longitude is invalid
	ErrInvalidLongitude = newAPPError(http.StatusBadRequest, 40004, "longitude is invalid")

	// ErrMissingID means given id is empty
	ErrMissingID = newAPPError(http.StatusBadRequest, 40005, "id is empty")

	// ErrMissingPassword means given password is empty
	ErrMissingPassword = newAPPError(http.StatusBadRequest, 40006, "password is empty")
)

var (
	// ErrCarNotExists means car not exists
	ErrCarNotExists = newAPPError(http.StatusNotFound, 40400, "car not exists")

	// ErrListCars means list car is failure
	ErrListCars = newAPPError(http.StatusInternalServerError, 50000, "list car is failure")
)

var (
	// ErrLogin means login is failure
	ErrLogin = newAPPError(http.StatusInternalServerError, 50001, "login is failure")
)

var (
	// ErrQueryArrears means query arrears is failure
	ErrQueryArrears = newAPPError(http.StatusInternalServerError, 50002, "query arrears is failure")
)

var (
	// ErrListBooking means list bookings is failure
	ErrListBooking = newAPPError(http.StatusInternalServerError, 50003, "list bookings is failure")

	// ErrBookingNotExists means booking is not exists
	ErrBookingNotExists = newAPPError(http.StatusNotFound, 40401, "booking is not exists")

	// ErrGetBookingByID means get booking by id is failure
	ErrGetBookingByID = newAPPError(http.StatusInternalServerError, 50006, "get booking by id is failure")

	// ErrBook means book a car is failure
	ErrBook = newAPPError(http.StatusInternalServerError, 50004, "book a car is failure")

	// ErrCancelBooking means cancel a booking is failure
	ErrCancelBooking = newAPPError(http.StatusInternalServerError, 50005, "cancel a booking is failure")
)
