package booking

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/irent/internal/app/irent/biz/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	orderE "github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/pb"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type impl struct {
	logger *zap.Logger
	biz    order.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz order.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "OrderHandler")),
		biz:    biz,
	}
}

// ListBookings
// @Summary List all bookings
// @Description List all bookings
// @Tags Bookings
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=[]pb.OrderInfo}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/bookings [get]
func (i *impl) ListBookings(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	token, ok := c.MustGet("token").(string)
	if !ok || len(token) == 0 {
		i.logger.Error(er.ErrMissingToken.Error())
		_ = c.Error(er.ErrMissingToken)
		return
	}

	orders, err := i.biz.List(ctx, 0, 0, &user.Profile{AccessToken: token})
	if err != nil {
		_ = c.Error(err)
		return
	}

	var ret []*pb.OrderInfo
	for _, info := range orders {
		ret = append(ret, orderE.NewOrderInfoResponse(info))
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// GetBookingByID
// @Summary Get a booking by id
// @Description Get a booking by id
// @Tags Bookings
// @Accept application/json
// @Produce application/json
// @Param id path string true "id"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=pb.OrderInfo}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/bookings/{id} [get]
func (i *impl) GetBookingByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	token, ok := c.MustGet("token").(string)
	if !ok || len(token) == 0 {
		i.logger.Error(er.ErrMissingToken.Error())
		_ = c.Error(er.ErrMissingToken)
		return
	}

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrMissingID.Error(), zap.Error(err))
		_ = c.Error(er.ErrMissingID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID, &user.Profile{ID: req.ID, AccessToken: token})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(orderE.NewOrderInfoResponse(ret)))
}

// Book
// @Summary Book a car
// @Description Book a car
// @Tags Bookings
// @Accept application/json
// @Produce application/json
// @Param car body bookRequest true "information of car"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=pb.Booking}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/bookings [post]
func (i *impl) Book(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	token, ok := c.MustGet("token").(string)
	if !ok || len(token) == 0 {
		i.logger.Error(er.ErrMissingToken.Error())
		_ = c.Error(er.ErrMissingToken)
		return
	}

	var data *bookRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrBook.Error(), zap.Error(err))
		_ = c.Error(er.ErrBook)
		return
	}

	ret, err := i.biz.BookCar(ctx, data.ID, data.ProjectID, &user.Profile{ID: data.UserID, AccessToken: token}, data.Circularly)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(orderE.NewBookingResponse(ret)))
}

// CancelBooking
// @Summary Cancel a booking by id
// @Description Cancel a booking by id
// @Tags Bookings
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of booking"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=string}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/bookings/{id} [delete]
func (i *impl) CancelBooking(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	token, ok := c.MustGet("token").(string)
	if !ok || len(token) == 0 {
		i.logger.Error(er.ErrMissingToken.Error())
		_ = c.Error(er.ErrMissingToken)
		return
	}

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrMissingID.Error(), zap.Error(err))
		_ = c.Error(er.ErrMissingID)
		return
	}

	err := i.biz.CancelBooking(ctx, req.ID, &user.Profile{ID: req.ID, AccessToken: token})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(req.ID))
}
