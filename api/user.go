package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog/api/response"
	"github.com/stewie1520/blog/core"
	"github.com/stewie1520/blog/usecases"
)

type userApi struct {
	app core.App
}

func bindUserApi(app core.App, ginEngine *gin.Engine) {
	api := &userApi{
		app: app,
	}

	subGroup := ginEngine.Group("/user")
	subGroup.GET("/me", api.getCurrentUserInfo)
}

// getCurrentUserInfo Return current logged in user information
// @Summary Get current user information
// @Description Return current logged in user information
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} usecases.GetUserByAccountIDResponse
// @Router /user/me [get]
func (api *userApi) getCurrentUserInfo(c *gin.Context) {
	q := usecases.NewGetUserByAccountIDQuery(api.app)
	q.AccountID = "123" // TODO: implement this

	if user, err := q.Execute(); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
