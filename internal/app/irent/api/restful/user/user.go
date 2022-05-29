package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare api handler's function
type IHandler interface {
	// Login serve caller to login the system
	Login(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
