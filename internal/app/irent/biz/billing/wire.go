//go:build wireinject

package billing

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(logger *zap.Logger, repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
