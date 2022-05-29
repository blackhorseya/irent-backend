package middlewares

import (
	"github.com/blackhorseya/gocommon/pkg/er"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseMiddleware serve caller to format api response
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()
			c.Errors = c.Errors[:0]

			switch err.Err.(type) {
			case *er.APPError:
				appError := err.Err.(*er.APPError)
				c.AbortWithStatusJSON(appError.Status, appError)
				break
			default:
				c.AbortWithStatus(http.StatusInternalServerError)
				break
			}
		}()

		c.Next()

	}
}
