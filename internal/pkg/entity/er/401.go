package er

import (
	"github.com/blackhorseya/gocommon/pkg/er"
	"net/http"
)

var (
	// ErrMissingToken means missing token in header
	ErrMissingToken = er.NewAPPError(http.StatusUnauthorized, 40100, "missing token")

	// ErrAuthHeaderFormat means must provide Authorization header with format `Bearer {token}`
	ErrAuthHeaderFormat = er.NewAPPError(http.StatusUnauthorized, 40101, "Must provide Authorization header with format `Bearer {token}`")
)
