//go:build wireinject

package user

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/user/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(logger *zap.Logger, repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
