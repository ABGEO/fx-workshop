package route

import (
	"github.com/abgeo/fx-workshop/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HelloRoute struct {
	handler *handler.HelloHandler
	logger  *zap.Logger
}

func NewHelloRoute(handler *handler.HelloHandler, logger *zap.Logger) *HelloRoute {
	return &HelloRoute{
		handler: handler,
		logger:  logger,
	}
}

func (route *HelloRoute) Register(engine *gin.Engine) {
	route.logger.Debug("Setting up route", zap.String("route", "hello"))

	group := engine.Group("/hello")
	{
		group.GET("", route.handler.Hello)
	}
}
