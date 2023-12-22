package main

import (
	"github.com/abgeo/fx-workshop/config"
	"github.com/abgeo/fx-workshop/database"
	"github.com/abgeo/fx-workshop/handler"
	"github.com/abgeo/fx-workshop/repository"
	"github.com/abgeo/fx-workshop/route"
	"github.com/abgeo/fx-workshop/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func provideLogger(conf *config.Config) (*zap.Logger, error) {
	if conf.Env == "prod" {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}

func main() {
	fxApp := fx.New(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Provide(
			provideLogger,
			config.New,
			database.New,
			fx.Annotate(server.New, fx.ParamTags(`group:"route"`)),
		),
		handler.Provide(),
		repository.Provide(),
		route.Provide(),
		fx.Invoke(server.Run),
	)

	fxApp.Run()
}
