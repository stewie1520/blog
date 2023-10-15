package user

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgx/v5"
	"github.com/stewie1520/blog/core"
	"github.com/stewie1520/blog/daos/dao_account"
	"github.com/stewie1520/blog/daos/dao_user"
	"github.com/stewie1520/blog/tools/securities"
	"github.com/stewie1520/blog/usecases"
)

var _ usecases.Command[*TokensResponse] = (*LoginCommand)(nil)

type LoginCommand struct {
	app        core.App
	daoUser    *dao_user.Queries
	daoAccount *dao_account.Queries

	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginCommand(app core.App) *LoginCommand {
	return &LoginCommand{
		app:        app,
		daoUser:    app.Dao().User,
		daoAccount: app.Dao().Account,
	}
}

func (cmd *LoginCommand) Validate() error {
	return validation.ValidateStruct(cmd,
		validation.Field(&cmd.Email, validation.Required, is.Email, validation.Length(1, 255)),
		validation.Field(&cmd.Password, validation.Required, validation.Length(8, 255)),
	)
}

func (cmd *LoginCommand) Execute(ctx context.Context) (*TokensResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	dbAccount, err := cmd.daoAccount.FindByEmail(ctx, cmd.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, usecases.ErrInvalidCredentials
	}

	if err != nil {
		return nil, err
	}

	if ok := securities.CompareHashAndPassword(dbAccount.Password, cmd.Password); !ok {
		return nil, usecases.ErrInvalidCredentials
	}

	dbUser, err := cmd.daoUser.GetUserByAccountID(ctx, dbAccount.ID)

	if err != nil {
		return nil, err
	}

	return createTokens(cmd.app.Config(), dbUser.ID.String(), dbAccount.ID.String())
}
