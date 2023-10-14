package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog/api/middleware"
	"github.com/stewie1520/blog/api/response"
	"github.com/stewie1520/blog/core"
	_ "github.com/stewie1520/blog/tools/types"
	usecases_post "github.com/stewie1520/blog/usecases/post"
)

type postApi struct {
	app core.App
}

func bindPostApi(app core.App, ginEngine *gin.Engine) {
	api := &postApi{
		app: app,
	}

	subGroup := ginEngine.Group("/post")
	subGroup.Use(middleware.RequireAuth(app))
	subGroup.POST("/", api.create)
	subGroup.GET("/", api.list)

}

// create create new post
// @Summary Create new post
// @Tags post
// @Accept json
// @Produce json
// @Param payload body usecases_post.CreatePostCommand true "Create post payload"
// @Success 200 {object} usecases_post.CreatePostResponse
// @Security Authorization
// @Router /post [post]
func (api *postApi) create(c *gin.Context) {
	cmd := usecases_post.NewCreatePostCommand(api.app)

	if err := c.ShouldBindJSON(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	user := getUserFromContext(c)
	cmd.UserID = user.ID.String()

	if res, err := cmd.Execute(); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// list list posts given user
// @Summary List posts given user
// @Tags post
// @Accept json
// @Produce json
// @Param pagination query usecases_post.ListByUserQuery true "Pagination"
// @Success 200 {object} types.Pagination[dao_post.Post]
// @Security Authorization
// @Router /post [get]
func (api *postApi) list(c *gin.Context) {
	q := usecases_post.NewListByUserQuery(api.app)

	if err := c.ShouldBindQuery(q); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	q.UserID = getUserFromContext(c).ID.String()

	if res, err := q.Execute(); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
