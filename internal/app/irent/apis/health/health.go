package health

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare apis handler's function
type IHandler interface {
	// Readiness to know when an application is ready to start accepting traffic
	Readiness(c *gin.Context)

	// Liveness to know when to restart an application
	Liveness(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
