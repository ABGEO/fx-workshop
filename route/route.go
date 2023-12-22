package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type IRoute interface {
	Register(engine *gin.Engine)
}

func AsRoute(constructor interface{}) interface{} {
	return fx.Annotate(
		constructor,
		fx.As(new(IRoute)),
		fx.ResultTags(`group:"route"`),
	)
}

func Provide() fx.Option {
	return fx.Provide(
		AsRoute(NewHelloRoute),
		AsRoute(NewProductRoute),
	)
}
