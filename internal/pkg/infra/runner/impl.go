package runner

import (
	"github.com/pkg/errors"
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

	taskC chan time.Time
	done  chan bool
}

// NewImpl return Runner
func NewImpl(o *Options, logger *zap.Logger) Runner {
	return &impl{
		o:       o,
		logger:  logger,
		app:     "",
		project: "",
		env:     "",
		taskC:   make(chan time.Time, 10000),
		done:    make(chan bool),
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
	// todo: 2022/7/25|sean|impl me
	return errors.New("impl me")
}
