package biz

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing"
	"github.com/blackhorseya/irent/internal/app/irent/biz/car"
	"github.com/blackhorseya/irent/internal/app/irent/biz/order"
	"github.com/blackhorseya/irent/internal/app/irent/biz/user"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	car.ProviderSet,
	user.ProviderSet,
	billing.ProviderSet,
	order.ProviderSet,
)
