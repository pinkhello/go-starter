package controller

import "go.uber.org/fx"

var Module = fx.Invoke(
	InitIndexController,
	InitBusinessGroupController,
)
