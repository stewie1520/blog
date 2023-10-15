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

	subGroup := ginEngine.Group("/posts")
	subGroup.Use(middleware.RequireAuth(app))
	subGroup.POST("/", api.create)
	subGroup.GET("/", api.list)
	subGroup.PUT("/:id", api.update)
	subGroup.DELETE("/:id", api.remove)
}

// create create new post
// @Summary Create new post
// @Tags post
// @Accept json
// @Produce json
// @Param payload body usecases_post.CreatePostCommand true "Create post payload"
// @Success 200 {object} usecases_post.CreatePostResponse
// @Security Authorization
// @Router /posts [post]
func (api *postApi) create(c *gin.Context) {
	cmd := usecases_post.NewCreatePostCommand(api.app)

	if err := c.ShouldBindJSON(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	user := getUserFromContext(c)
	cmd.UserID = user.ID.String()

	if res, err := cmd.Execute(c); err != nil {
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
// @Router /posts [get]
func (api *postApi) list(c *gin.Context) {
	q := usecases_post.NewListByUserQuery(api.app)

	if err := c.ShouldBindQuery(q); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	q.UserID = getUserFromContext(c).ID.String()

	if res, err := q.Execute(c); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// update update content of a post
// @Summary Update content of a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body usecases_post.UpdatePostCommand true "post"
// @Param id path string true "post id"
// @Success 200 {object} usecases_post.UpdatePostResponse
// @Security Authorization
// @Router /posts/{id} [put]
func (api *postApi) update(c *gin.Context) {
	cmd := usecases_post.NewUpdatePostCommand(api.app)

	if err := c.ShouldBindUri(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	if err := c.ShouldBindJSON(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	cmd.UserID = getUserFromContext(c).ID.String()

	if res, err := cmd.Execute(c); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// remove remove a post
// @Summary remove a post
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "post id"
// @Success 200 {object} usecases_post.RemovePostResponse
// @Security Authorization
// @Router /posts/{id} [delete]
func (api *postApi) remove(c *gin.Context) {
	cmd := usecases_post.NewRemovePostCommand(api.app)
	if err := c.ShouldBindUri(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	cmd.UserID = getUserFromContext(c).ID.String()

	if res, err := cmd.Execute(c); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
