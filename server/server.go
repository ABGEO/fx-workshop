package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/abgeo/fx-workshop/config"
	"github.com/abgeo/fx-workshop/model"
	"github.com/abgeo/fx-workshop/route"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Params struct {
	fx.In

	Server *http.Server
	Logger *zap.Logger
	Config *config.Config
	DB     *gorm.DB
}

func New(routes []route.IRoute, conf *config.Config) (*http.Server, error) {
	if conf.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	if conf.Env == "test" {
		gin.SetMode(gin.TestMode)
	}

	engine := gin.New()
	if err := engine.SetTrustedProxies(conf.Server.TrustedProxies); err != nil {
		return nil, fmt.Errorf("unable to set trusted proxies: %w", err)
	}

	for _, r := range routes {
		r.Register(engine)
	}

	return &http.Server{
		Addr:              net.JoinHostPort(conf.Server.ListenAddr, conf.Server.Port),
		Handler:           engine,
		ReadHeaderTimeout: 0,
	}, nil
}

func Run(params Params, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("Running database migrations")

			if err := params.DB.AutoMigrate(&model.Product{}); err != nil {
				return fmt.Errorf("failed to run migrations: %w", err)
			}

			params.Logger.Info(
				"Starting HTTP Server",
				zap.String("address", params.Config.Server.ListenAddr),
				zap.String("port", params.Config.Server.Port),
			)

			go func() {
				if err := params.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					params.Logger.Fatal("Unable to start HTTP Server", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Shutting down HTTP Server")

			if err := params.Server.Shutdown(ctx); err != nil {
				return fmt.Errorf("failed to shutdown server: %w", err)
			}

			return nil
		},
	})

	return nil
}
