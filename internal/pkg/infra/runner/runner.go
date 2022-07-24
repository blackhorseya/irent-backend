package runner

import (
	"github.com/google/wire"
)

// Runner declare a runner functions
type Runner interface {
	// Application set up then application information
	Application(app, project, env string)

	// Start the runner
	Start() error

	// Stop the runner
	Stop() error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewOptions, NewImpl)
