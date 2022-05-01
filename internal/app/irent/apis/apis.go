package apis

import (
	"github.com/blackhorseya/irent/internal/app/irent/apis/billing"
	"github.com/blackhorseya/irent/internal/app/irent/apis/booking"
	"github.com/blackhorseya/irent/internal/app/irent/apis/cars"
	"github.com/blackhorseya/irent/internal/app/irent/apis/health"
	"github.com/blackhorseya/irent/internal/app/irent/apis/user"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(
	healthH health.IHandler,
	carH cars.IHandler,
	userH user.IHandler,
	billingH billing.IHandler,
	bookingH booking.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", healthH.Readiness)
			api.GET("liveness", healthH.Liveness)

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			v1 := api.Group("v1")
			{
				carG := v1.Group("cars")
				{
					carG.GET("", carH.NearTopN)
				}

				authG := v1.Group("auth")
				{
					authG.POST("login", userH.Login)
				}

				billingG := v1.Group("billing")
				{
					billingG.GET(":id/arrears", middlewares.AuthMiddleware(), billingH.GetArrears)
				}

				bookingG := v1.Group("bookings", middlewares.AuthMiddleware())
				{
					bookingG.GET("", bookingH.ListBookings)
					bookingG.GET(":id", bookingH.GetBookingByID)
					bookingG.POST("", bookingH.Book)
					bookingG.DELETE(":id", bookingH.CancelBooking)
				}
			}
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	cars.ProviderSet,
	user.ProviderSet,
	billing.ProviderSet,
	booking.ProviderSet,
	CreateInitHandlerFn,
)
