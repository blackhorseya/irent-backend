//go:build wireinject

package billing

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz billing.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
