package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/abgeo/fx-workshop/handler"
	"github.com/gin-gonic/gin"
)

func New(env, address, port string, trustedProxies []string) (*http.Server, error) {
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	if env == "test" {
		gin.SetMode(gin.TestMode)
	}

	engine := gin.New()
	if err := engine.SetTrustedProxies(trustedProxies); err != nil {
		return nil, fmt.Errorf("unable to set trusted proxies: %w", err)
	}

	err := registerHandlers(engine)
	if err != nil {
		return nil, fmt.Errorf("failed to register handlers: %w", err)
	}

	return &http.Server{
		Addr:              net.JoinHostPort(address, port),
		Handler:           engine,
		ReadHeaderTimeout: 0,
	}, nil
}

func registerHandlers(engine *gin.Engine) error {
	helloHandler := handler.NewHelloHandler()

	productHandler, err := handler.NewProductHandler()
	if err != nil {
		return fmt.Errorf("failed to initialize product handler: %w", err)
	}

	helloGroup := engine.Group("/hello")
	{
		helloGroup.GET("", helloHandler.Hello)
	}

	productGroup := engine.Group("/product")
	{
		productGroup.POST("", productHandler.Create)
		productGroup.GET("", productHandler.GetAll)
		productGroup.GET(":id", productHandler.GetSingle)
	}

	return nil
}
