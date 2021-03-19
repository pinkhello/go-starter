package libs

import "go.uber.org/fx"

var XormModule = fx.Provide(NewXorm)
