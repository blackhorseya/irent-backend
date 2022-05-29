//go:build wireinject

package user

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/user"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz user.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
