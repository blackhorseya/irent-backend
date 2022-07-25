package restclient

import (
	wire "github.com/google/wire"
	"net/http"
)

// HTTPClient interface
//go:generate mockery --name=HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient return HTTPClient
func NewClient() HTTPClient {
	return &http.Client{}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewClient)
