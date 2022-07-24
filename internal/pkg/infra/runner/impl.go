package runner

import (
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

	// todo: 2022/7/25|sean|impl me

	return nil
}

func (i *impl) Stop() error {
	i.logger.Info("stopping runner engine...")

	// todo: 2022/7/25|sean|impl me

	return nil
}
