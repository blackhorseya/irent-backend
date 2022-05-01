package cars

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare apis handler's function
type IHandler interface {
	// NearTopN serve caller to list closer car
	NearTopN(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
