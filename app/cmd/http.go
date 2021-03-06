package cmd

import (
	"github.com/spf13/cobra"
	"go-starter/config"
	"go-starter/internal/controller"
	"go-starter/internal/http"
	"go-starter/internal/lib"
	"go-starter/internal/repository"
	"go-starter/internal/service"
	"go-starter/utils"
	"go.uber.org/fx"
)

var (
	httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Start Http REST API",
		Run:   initHttp,
	}
)

func initHttp(cmd *cobra.Command, args []string) {
	fx.New(inject()).Run()
}

func inject() fx.Option {
	return fx.Options(
		fx.Provide(
			config.NewConfig,
			utils.NewTimeoutContext,
		),
		lib.Module,
		repository.Module,
		service.Module,
		controller.Module,
		http.Module,
	)
}
