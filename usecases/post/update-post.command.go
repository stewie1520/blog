package post

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stewie1520/blog/core"
	"github.com/stewie1520/blog/daos/dao_post"
	"github.com/stewie1520/blog/usecases"
)

var _ usecases.Command[*UpdatePostResponse] = (*UpdatePostCommand)(nil)

type UpdatePostCommand struct {
	app      core.App
	dao_post *dao_post.Queries

	ID      string `json:"-" uri:"id"`
	Content string `json:"content"`
	UserID  string `json:"-"`
}

type UpdatePostResponse struct {
	*dao_post.Post
}

func NewUpdatePostCommand(app core.App) *UpdatePostCommand {
	return &UpdatePostCommand{
		app:      app,
		dao_post: app.Dao().Post,
	}
}

func (cmd *UpdatePostCommand) Validate() error {
	return validation.ValidateStruct(cmd,
		validation.Field(&cmd.ID, validation.Required, is.UUIDv4),
		validation.Field(&cmd.UserID, validation.Required, is.UUIDv4),
		validation.Field(&cmd.Content, validation.Required),
	)
}

func (cmd *UpdatePostCommand) Execute(ctx context.Context) (*UpdatePostResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	dbPost, err := cmd.dao_post.FindById(ctx, uuid.MustParse(cmd.ID))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, usecases.ErrEntityNotFound
	}

	if err != nil {
		return nil, err
	}

	if dbPost.UserID != uuid.MustParse(cmd.UserID) {
		return nil, usecases.ErrEntityForbidden
	}

	dbPost.Content = cmd.Content
	err = cmd.dao_post.UpdatePost(ctx, dao_post.UpdatePostParams{
		ID:      dbPost.ID,
		Content: dbPost.Content,
	})

	if err != nil {
		return nil, err
	}

	return &UpdatePostResponse{&dbPost}, nil
}
