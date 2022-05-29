package health

import (
	"github.com/blackhorseya/gocommon/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger) IHandler {
	return &impl{logger: logger}
}

// Readiness
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}

// Liveness
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}
