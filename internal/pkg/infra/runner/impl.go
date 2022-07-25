package runner

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/order"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

// Options declare runner configuration
type Options struct {
	Interval time.Duration
}

// NewOptions return *Options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	o := new(Options)

	err := v.UnmarshalKey("runner", &o)
	if err != nil {
		return nil, err
	}

	logger.Info("load runner options success")

	return o, nil
}

type impl struct {
	o      *Options
	logger *zap.Logger

	app     string
	project string
	env     string

	orderBiz order.IBiz

	taskC chan time.Time
	done  chan bool
}

// NewImpl return Runner
func NewImpl(o *Options, logger *zap.Logger, orderBiz order.IBiz) Runner {
	return &impl{
		o:        o,
		logger:   logger,
		app:      "",
		project:  "",
		env:      "",
		orderBiz: orderBiz,
		taskC:    make(chan time.Time, 10000),
		done:     make(chan bool),
	}
}

func (i *impl) Application(app, project, env string) {
	i.app = app
	i.project = project
	i.env = env
}

func (i *impl) Start() error {
	i.logger.Info("runner engine starting...")

	go i.worker()

	return nil
}

func (i *impl) Stop() error {
	i.logger.Info("stopping runner engine...")

	i.done <- true

	return nil
}

func (i *impl) worker() {
	ticker := time.NewTicker(i.o.Interval)

	for {
		select {
		case <-i.done:
			return
		case <-ticker.C:
			i.ExecuteTo()
		case t := <-i.taskC:
			err := i.Execute(t)
			if err != nil {
				i.logger.Error("executor occurs error", zap.Error(err))
			}
		}
	}
}

func (i *impl) ExecuteTo() {
	select {
	case i.taskC <- time.Now():
	case <-time.After(50 * time.Millisecond):
		return
	}
}

func (i *impl) Execute(t time.Time) error {
	ctx := contextx.Background()

	bookings, err := i.orderBiz.ListPremiumBookings(ctx)
	if err != nil {
		return err
	}

	for profile, booking := range bookings {
		if booking.LastPickAt.Add(-5 * time.Minute).Before(t) {
			ret, err := i.orderBiz.ReBookCar(ctx, booking.No, booking.CarID, booking.ProjID, profile)
			if err != nil {
				return err
			}

			_, err = i.orderBiz.UpdatePremiumBooking(ctx, profile, ret)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
