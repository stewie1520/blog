package post

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/stewie1520/blog/core"
	"github.com/stewie1520/blog/daos/dao_post"
	"github.com/stewie1520/blog/usecases"
)

var _ usecases.Command[*RemovePostResponse] = (*RemovePostCommand)(nil)

type RemovePostCommand struct {
	app      core.App
	dao_post *dao_post.Queries

	ID     string `json:"-" uri:"id"`
	UserID string `json:"-"`
}

type RemovePostResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewRemovePostCommand(app core.App) *RemovePostCommand {
	return &RemovePostCommand{
		app:      app,
		dao_post: app.Dao().Post,
	}
}

func (cmd *RemovePostCommand) Validate() error {
	return validation.ValidateStruct(cmd,
		validation.Field(&cmd.ID, validation.Required, is.UUIDv4),
		validation.Field(&cmd.UserID, validation.Required, is.UUIDv4),
	)
}

func (cmd *RemovePostCommand) Execute(ctx context.Context) (*RemovePostResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	dbPost, err := cmd.dao_post.FindById(ctx, uuid.MustParse(cmd.ID))
	if err != nil {
		return nil, usecases.ErrEntityNotFound
	}

	if dbPost.UserID != uuid.MustParse(cmd.UserID) {
		return nil, usecases.ErrEntityForbidden
	}

	err = cmd.dao_post.RemovePost(ctx, dbPost.ID)

	if err != nil {
		return nil, err
	}

	return &RemovePostResponse{
		ID: dbPost.ID,
	}, nil
}
