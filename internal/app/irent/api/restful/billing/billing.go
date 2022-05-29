package billing

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare billing api handler function
type IHandler interface {
	// GetArrears serve caller to get user's arrears
	GetArrears(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
