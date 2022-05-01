package runner

import (
	"time"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options declare runner configuration
type Options struct {
	Interval time.Duration
}

// NewOptions serve caller to create an Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	err = v.UnmarshalKey("runner", &o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// Engine declare a runner engine struct
type Engine struct {
	o      *Options
	logger *zap.Logger
	app    string

	done chan bool
}

// NewEngine serve caller to create a Runner
func NewEngine(o *Options, logger *zap.Logger) (Runner, error) {
	ret := &Engine{
		o:      o,
		logger: logger.With(zap.String("type", "Runner")),
		done:   make(chan bool),
	}

	return ret, nil
}

// Application serve caller to set app name
func (e *Engine) Application(app string) {
	e.app = app
}

// Start runner engine
func (e *Engine) Start() error {
	e.logger.Info("runner engine starting...")

	go e.work()

	return nil
}

// Stop runner engine
func (e *Engine) Stop() error {
	e.logger.Info("stopping runner engine")

	e.done <- true

	return nil
}

func (e *Engine) work() {
	ticker := time.NewTicker(e.o.Interval)

	for {
		select {
		case <-e.done:
			return
		case t := <-ticker.C:
			e.logger.Debug("ticker", zap.Time("at", t))
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewOptions, NewEngine)
