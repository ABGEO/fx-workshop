package handler

import "go.uber.org/fx"

func Provide() fx.Option {
	return fx.Provide(
		NewHelloHandler,
		fx.Annotate(NewProductHandler, fx.ParamTags(`name:"ProductRepository"`)),
	)
}
