package repository

import "go.uber.org/fx"

func Provide() fx.Option {
	return fx.Provide(
		fx.Annotate(
			NewProductRepository,
			fx.As(new(IProductRepository)),
			fx.ResultTags(`name:"ProductRepository"`),
		),
	)
}
