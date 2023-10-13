package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog/api/middleware"
	"github.com/stewie1520/blog/core"
)

func InitApi(app core.App) (*gin.Engine, error) {
	if app.IsDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	bindHealthApi(app, engine)

	engine.Use(middleware.Cors(app))

	bindUserApi(app, engine)

	// TODO: add more api group here

	return engine, nil
}
