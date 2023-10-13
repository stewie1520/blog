package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog/api/response"
	"github.com/stewie1520/blog/core"
	usecases_user "github.com/stewie1520/blog/usecases/user"
)

type userApi struct {
	app core.App
}

func bindUserApi(app core.App, ginEngine *gin.Engine) {
	api := &userApi{
		app: app,
	}

	subGroup := ginEngine.Group("/user")
	subGroup.POST("/register", api.register)
}

// register Register new user
// @Summary Register new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body usecases_user.RegisterCommand true "Register payload"
// @Success 200 {object} usecases_user.RegisterResponse
// @Router /user/register [post]
func (api *userApi) register(c *gin.Context) {
	cmd := usecases_user.NewRegisterCommand(api.app)

	if err := c.ShouldBindJSON(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	if res, err := cmd.Execute(); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}
