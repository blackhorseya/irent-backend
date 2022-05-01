//go:build wireinject

package booking

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/order"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz order.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
