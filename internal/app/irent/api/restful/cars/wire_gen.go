// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cars

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/car"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz car.IBiz) (IHandler, error) {
	iHandler := NewImpl(logger, biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
