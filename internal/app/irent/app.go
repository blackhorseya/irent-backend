package irent

import (
	// import swagger docs
	_ "github.com/blackhorseya/irent/api/docs"
	"github.com/blackhorseya/irent/internal/pkg/app"
	"github.com/blackhorseya/irent/internal/pkg/infra/runner"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options declare options configuration
type Options struct {
	Name string
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, err
	}

	logger.Info("load application options success")

	return o, nil
}

// New serve caller to create an *app.Application
func New(o *Options, logger *zap.Logger, hs *http.Server, runner runner.Runner) (*app.Application, error) {
	a, err := app.New(o.Name, logger, app.HTTPServerOption(hs), app.RunnerOption(runner))
	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New, NewOptions)
