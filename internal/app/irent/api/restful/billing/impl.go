package billing

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	_ "github.com/blackhorseya/gocommon/pkg/er" // import er struct
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing"
	cer "github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type impl struct {
	logger *zap.Logger
	biz    billing.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz billing.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "BillingHandler")),
		biz:    biz,
	}
}

type reqID struct {
	ID string `uri:"id" binding:"required"`
}

// GetArrears
// @Summary Get arrears by user's id
// @Description Get arrears by user's id
// @Tags Billing
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of user"
// @Success 200 {object} response.Response{data=[]pb.Arrears}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Security ApiKeyAuth
// @Router /v1/billing/{id}/arrears [get]
func (i *impl) GetArrears(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	token, ok := c.MustGet("token").(string)
	if !ok || len(token) == 0 {
		i.logger.Error(cer.ErrMissingToken.Error())
		_ = c.Error(cer.ErrMissingToken)
		return
	}

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(cer.ErrMissingID.Error(), zap.Error(err))
		_ = c.Error(cer.ErrMissingID)
		return
	}

	ret, err := i.biz.GetArrears(ctx, &user.Profile{ID: req.ID, AccessToken: token})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
