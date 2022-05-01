//go:build wireinject

package cars

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/car"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz car.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
