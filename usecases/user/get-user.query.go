package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/stewie1520/blog/core"
	"github.com/stewie1520/blog/daos/dao_user"
	"github.com/stewie1520/blog/usecases"
)

var _ usecases.Query[dao_user.User] = (*GetUserQuery)(nil)

func NewGetUserQuery(app core.App) *GetUserQuery {
	return &GetUserQuery{
		app: app,
		dao: app.Dao().User,
	}
}

type GetUserQuery struct {
	app core.App
	dao *dao_user.Queries

	ID string `json:"id"`
}

// Validate implements Query.
func (q *GetUserQuery) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(&q.ID, validation.Required, is.UUIDv4),
	)
}

// Execute implements Query.
func (q *GetUserQuery) Execute(ctx context.Context) (dao_user.User, error) {
	if err := q.Validate(); err != nil {
		return dao_user.User{}, err
	}

	user, err := q.dao.FindUserById(ctx, uuid.MustParse(q.ID))
	return user, err
}
