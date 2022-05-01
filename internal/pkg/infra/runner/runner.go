package runner

// Runner declare a runner functions
type Runner interface {
	Application(app string)
	Start() error
	Stop() error
}
