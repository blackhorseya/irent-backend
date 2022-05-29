package er

import (
	"github.com/blackhorseya/gocommon/pkg/er"
	"net/http"
)

var (
	// ErrInvalidStart means given start is invalid
	ErrInvalidStart = er.NewAPPError(http.StatusBadRequest, 40000, "start is invalid")

	// ErrInvalidEnd means given end is invalid
	ErrInvalidEnd = er.NewAPPError(http.StatusBadRequest, 40001, "end is invalid")

	// ErrInvalidN means given N is invalid
	ErrInvalidN = er.NewAPPError(http.StatusBadRequest, 40002, "N is invalid")

	// ErrInvalidLatitude means given latitude is invalid
	ErrInvalidLatitude = er.NewAPPError(http.StatusBadRequest, 40003, "latitude is invalid")

	// ErrInvalidLongitude means given longitude is invalid
	ErrInvalidLongitude = er.NewAPPError(http.StatusBadRequest, 40004, "longitude is invalid")

	// ErrMissingID means given id is empty
	ErrMissingID = er.NewAPPError(http.StatusBadRequest, 40005, "id is empty")

	// ErrMissingPassword means given password is empty
	ErrMissingPassword = er.NewAPPError(http.StatusBadRequest, 40006, "password is empty")
)
