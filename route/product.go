package route

import (
	"github.com/abgeo/fx-workshop/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductRoute struct {
	handler *handler.ProductHandler
	logger  *zap.Logger
}

func NewProductRoute(handler *handler.ProductHandler, logger *zap.Logger) *ProductRoute {
	return &ProductRoute{
		handler: handler,
		logger:  logger,
	}
}

func (route *ProductRoute) Register(engine *gin.Engine) {
	route.logger.Debug("Setting up route", zap.String("route", "product"))

	group := engine.Group("/product")
	{
		group.POST("", route.handler.Create)
		group.GET("", route.handler.GetAll)
		group.GET(":id", route.handler.GetSingle)
	}
}
