package middlewares

import (
	"strings"

	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// AuthMiddleware serve caller to extract authorization header value
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		err := c.ShouldBindHeader(&h)
		if err != nil {
			c.Error(er.ErrMissingToken)
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")
		if len(idTokenHeader) < 2 {
			c.Error(er.ErrAuthHeaderFormat)
			c.Abort()
			return
		}

		// todo: 2021-05-13|11:40|doggy|get by token then set user entity
		c.Set("token", idTokenHeader[1])

		c.Next()
	}
}
