package er

import (
	"github.com/blackhorseya/gocommon/pkg/er"
	"net/http"
)

var (
	// ErrCarNotExists means car not exists
	ErrCarNotExists = er.NewAPPError(http.StatusNotFound, 40400, "car not exists")

	// ErrBookingNotExists means booking is not exists
	ErrBookingNotExists = er.NewAPPError(http.StatusNotFound, 40401, "booking is not exists")
)
