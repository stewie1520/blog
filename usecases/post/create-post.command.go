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

type CreatePostCommand struct {
	app     core.App
	daoPost *dao_post.Queries

	Content string `json:"content"`
	UserID  string `json:"-"`
}

var _ usecases.Command[*CreatePostResponse] = (*CreatePostCommand)(nil)

type CreatePostResponse struct {
	*dao_post.Post
}

func NewCreatePostCommand(app core.App) *CreatePostCommand {
	return &CreatePostCommand{
		app:     app,
		daoPost: app.Dao().Post,
	}
}

func (cmd *CreatePostCommand) Validate() error {
	return validation.ValidateStruct(cmd,
		validation.Field(&cmd.Content, validation.Required),
		validation.Field(&cmd.UserID, validation.Required, is.UUIDv4),
	)
}

func (cmd *CreatePostCommand) Execute(ctx context.Context) (*CreatePostResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	newPost, err := cmd.daoPost.Create(ctx, dao_post.CreateParams{
		ID:      uuid.New(),
		Content: cmd.Content,
		UserID:  uuid.MustParse(cmd.UserID),
	})

	if err != nil {
		return nil, err
	}

	return &CreatePostResponse{
		&newPost,
	}, nil
}
