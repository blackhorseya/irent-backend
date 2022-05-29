package cars

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/car"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/entity/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    car.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz car.IBiz) IHandler {
	return &impl{logger: logger, biz: biz}
}

// NearTopN
// @Summary List closer car
// @Description List closer car
// @Tags Cars
// @Accept application/json
// @Produce application/json
// @Param n query integer false "n" default(10)
// @Param latitude query number false "latitude" default(0)
// @Param longitude query number false "longitude" default(0)
// @Success 200 {object} response.Response{data=[]pb.Car}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/cars [get]
func (i *impl) NearTopN(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	n, err := strconv.Atoi(c.DefaultQuery("n", "10"))
	if err != nil {
		i.logger.Error(er.ErrInvalidN.Error(), zap.Error(err), zap.String("n", c.Query("n")))
		_ = c.Error(er.ErrInvalidN)
		return
	}

	latitude, err := strconv.ParseFloat(c.DefaultQuery("latitude", "0"), 64)
	if err != nil {
		i.logger.Error(er.ErrInvalidLatitude.Error(), zap.Error(err), zap.String("latitude", c.Query("latitude")))
		_ = c.Error(er.ErrInvalidLatitude)
		return
	}

	longitude, err := strconv.ParseFloat(c.DefaultQuery("longitude", "0"), 64)
	if err != nil {
		i.logger.Error(er.ErrInvalidLongitude.Error(), zap.Error(err), zap.String("longitude", c.Query("longitude")))
		_ = c.Error(er.ErrInvalidLongitude)
		return
	}

	ret, _, err := i.biz.NearTopN(ctx, n, latitude, longitude)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(n))
	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
