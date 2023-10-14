package daos

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stewie1520/blog/daos/dao_account"
	"github.com/stewie1520/blog/daos/dao_post"
	"github.com/stewie1520/blog/daos/dao_user"
)

type Dao struct {
	db *pgxpool.Pool

	Builder squirrel.StatementBuilderType
	User    *dao_user.Queries
	Account *dao_account.Queries
	Post    *dao_post.Queries
}

func New(db *pgxpool.Pool) *Dao {
	return &Dao{
		db:      db,
		User:    dao_user.New(db),
		Account: dao_account.New(db),
		Post:    dao_post.New(db),
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (dao *Dao) DB() *pgxpool.Pool {
	return dao.db
}
