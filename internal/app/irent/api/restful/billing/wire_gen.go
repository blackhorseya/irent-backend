// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package billing

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing"
	"github.com/google/wire"
	"go.uber.org/zap"
)

import (
	_ "github.com/blackhorseya/gocommon/pkg/er"
)

// Injectors from wire.go:

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz billing.IBiz) (IHandler, error) {
	iHandler := NewImpl(logger, biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
