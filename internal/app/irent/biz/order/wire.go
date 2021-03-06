//go:build wireinject

package order

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/order/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(o *Options, logger *zap.Logger, repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
