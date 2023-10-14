package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog/api/middleware"
	"github.com/stewie1520/blog/core"
	"github.com/stewie1520/blog/daos/dao_user"
)

func InitApi(app core.App) (*gin.Engine, error) {
	if app.IsDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())

	bindHealthApi(app, engine)

	engine.Use(middleware.Cors(app))
	engine.Use(middleware.LoadAuthContext(app))

	bindUserApi(app, engine)

	// TODO: add more api group here

	return engine, nil
}

func getUserFromContext(c *gin.Context) *dao_user.User {
	value, _ := c.Get(middleware.ContextUserKey)
	user, ok := value.(*dao_user.User)

	if user == nil || !ok {
		return nil
	}

	return user
}
