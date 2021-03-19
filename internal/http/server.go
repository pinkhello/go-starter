package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-starter/config"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewServer)

func NewServer(lifecycle fx.Lifecycle, config config.Config) *echo.Echo {
	logrus.SetLevel(logrus.DebugLevel)
	instance := echo.New()
	middleware := InitMiddleware()

	instance.Use(middleware.CORS)
	instance.Use(middleware.Logger)
	instance.Use(middleware.Recover)
	instance.Use(middleware.JWT)

	instance.HTTPErrorHandler = middleware.ErrorHandler

	instance.GET("/swagger/*", echoSwagger.WrapHandler)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logrus.Print("Start Http Server.")
			go func() {
				err := instance.Start(config.Server.Address)
				if err != nil {
					logrus.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Print("Stopping Http Server.")
			return instance.Shutdown(ctx)
		},
	})
	return instance
}
