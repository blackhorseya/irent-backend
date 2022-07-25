//go:build wireinject

package runner

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/order"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRunner(o *Options, logger *zap.Logger, orderBiz order.IBiz) (Runner, error) {
	panic(wire.Build(testProviderSet))
}
