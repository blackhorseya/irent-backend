package user

import (
	"github.com/blackhorseya/gocommon/pkg/response"
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/user"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    user.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz user.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "UserHandler")),
		biz:    biz,
	}
}

// Login
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce application/json
// @Param id formData string true "user id"
// @Param password formData string true "user password"
// @Success 201 {object} response.Response{data=pb.Profile}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/auth/login [post]
func (i *impl) Login(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	id := c.PostForm("id")
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error())
		_ = c.Error(er.ErrMissingID)
		return
	}

	password := c.PostForm("password")
	if len(password) == 0 {
		i.logger.Error(er.ErrMissingPassword.Error())
		_ = c.Error(er.ErrMissingPassword)
		return
	}

	ret, err := i.biz.Login(ctx, id, password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}
