package middlewares

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/gin-gonic/gin"
)

// ContextMiddleware serve caller to added Contextx into gin
func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := contextx.Background()
		c.Set("ctx", ctx)

		c.Next()
	}
}
