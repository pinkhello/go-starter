package repository

import (
	"go-starter/internal/repository/mysql"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	mysql.NewBusinessGroupRepository,
)
