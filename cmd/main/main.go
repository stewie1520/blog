package main

import (
	"flag"
	"fmt"

	"github.com/stewie1520/blog/api"
	"github.com/stewie1520/blog/config"
	"github.com/stewie1520/blog/core"
	docs "github.com/stewie1520/blog/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var debug = flag.Bool("debug", false, "debug mode")

func init() {
	flag.Parse()
}

// @title Blog API
// @version 1.0
// @BasePath /
func main() {
	cfg, err := config.Init()
	panicIfError(err)

	app := core.NewBaseApp(core.BaseAppConfig{
		Config:  cfg,
		IsDebug: *debug,
	})

	err = app.Bootstrap()
	panicIfError(err)

	router, err := api.InitApi(app)
	panicIfError(err)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(fmt.Sprintf(":%d", cfg.Port))
}

func panicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
