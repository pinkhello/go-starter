package repository

import (
	"go-starter/internal/repository/mysql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		mysql.NewBusinessGroupRepository,
	),
)
