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
	"github.com/stewie1520/blog/tools/types"
	"github.com/stewie1520/blog/usecases"
)

var _ usecases.Query[*types.Pagination[dao_post.Post]] = (*ListByUserQuery)(nil)

type ListByUserQuery struct {
	app core.App
	dao *dao_post.Queries

	UserID string `json:"-"`
	*types.PaginationParams
}

func NewListByUserQuery(app core.App) *ListByUserQuery {
	return &ListByUserQuery{
		app:              app,
		dao:              app.Dao().Post,
		PaginationParams: &types.PaginationParams{},
	}
}

func (q *ListByUserQuery) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(&q.UserID, validation.Required, is.UUIDv4),
	)
}

func (q *ListByUserQuery) Execute(ctx context.Context) (*types.Pagination[dao_post.Post], error) {
	posts, err := q.dao.ListByUserId(ctx, dao_post.ListByUserIdParams{
		UserID: uuid.MustParse(q.UserID),
		Limit:  int32(q.Limit),
		Offset: int32(q.Offset),
	})

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	total, err := q.dao.CountByUserId(ctx, uuid.MustParse(q.UserID))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	return types.NewPagination(posts, total, int64(q.Offset), int64(q.Limit)), nil
}
