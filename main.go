package main

import (
	"github.com/abgeo/fx-workshop/config"
	"github.com/abgeo/fx-workshop/database"
	"github.com/abgeo/fx-workshop/model"
	"github.com/abgeo/fx-workshop/server"
	"go.uber.org/zap"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log, err := zap.NewDevelopment()
	panicOnErr(err)

	conf, err := config.New()
	panicOnErr(err)

	srv, err := server.New(conf.Env, conf.Server.ListenAddr, conf.Server.Port, conf.Server.TrustedProxies)
	panicOnErr(err)

	db, err := database.New(conf.Database.FilePath)
	panicOnErr(err)

	// Run Migrations.
	err = db.AutoMigrate(&model.Product{})
	panicOnErr(err)

	log.Info(
		"Starting HTTP Server",
		zap.String("address", conf.Server.ListenAddr),
		zap.String("port", conf.Server.Port),
	)

	err = srv.ListenAndServe()
	panicOnErr(err)
}
