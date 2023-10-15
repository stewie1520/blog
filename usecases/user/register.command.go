package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stewie1520/blog/config"
	"github.com/stewie1520/blog/core"
	"github.com/stewie1520/blog/daos/dao_account"
	"github.com/stewie1520/blog/daos/dao_user"
	"github.com/stewie1520/blog/tools/securities"
	"github.com/stewie1520/blog/tools/types"
	"github.com/stewie1520/blog/usecases"
	"go.uber.org/zap"
)

var _ usecases.Command[*TokensResponse] = (*RegisterCommand)(nil)

type RegisterCommand struct {
	app        core.App
	daoUser    *dao_user.Queries
	daoAccount *dao_account.Queries

	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Bio      string `json:"bio"`
}

func NewRegisterCommand(app core.App) *RegisterCommand {
	return &RegisterCommand{
		app:        app,
		daoUser:    app.Dao().User,
		daoAccount: app.Dao().Account,
	}
}

func (cmd *RegisterCommand) Validate() error {
	return validation.ValidateStruct(cmd,
		validation.Field(&cmd.FullName, validation.Required, validation.Length(1, 255)),
		validation.Field(&cmd.Password, validation.Required, validation.Length(8, 255)),
		validation.Field(&cmd.Email, validation.Required, is.Email, validation.Length(1, 255)),
	)
}

func (cmd *RegisterCommand) Execute(ctx context.Context) (*TokensResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	if password, err := securities.HashPassword(cmd.Password); err != nil {
		return nil, err
	} else {
		cmd.Password = password
	}

	tx, err := cmd.app.DB().BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil {
			cmd.app.Log().Info("Rollback error:", zap.Error(err))
		}
	}()

	daoAccountTx := cmd.daoAccount.WithTx(tx)
	daoUserTx := cmd.daoUser.WithTx(tx)

	dbAccount, err := daoAccountTx.CreateAccount(ctx, dao_account.CreateAccountParams{
		ID:       uuid.New(),
		Email:    cmd.Email,
		Password: cmd.Password,
	})

	if err != nil {
		return nil, err
	}

	dbUser, err := daoUserTx.CreateUser(ctx, dao_user.CreateUserParams{
		ID:        uuid.New(),
		FullName:  cmd.FullName,
		AccountID: dbAccount.ID,
		Bio:       types.NewNullString(&cmd.Bio),
	})

	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return createTokens(cmd.app.Config(), dbUser.ID.String(), dbAccount.ID.String())
}

func createTokens(config *config.Config, userId string, accountId string) (*TokensResponse, error) {
	accessToken, err := securities.NewPaseto(map[string]string{
		"userId":    userId,
		"accountId": accountId,
		"type":      "access",
	},
		config.Token.PrivateKey,
		config.Token.AccessTokenTTL,
	)

	if err != nil {
		return nil, err
	}

	refreshToken, err := securities.NewPaseto(map[string]string{
		"userId":    userId,
		"accountId": accountId,
		"type":      "refresh",
	},
		config.Token.PrivateKey,
		config.Token.RefreshTokenTTL,
	)

	if err != nil {
		return nil, err
	}

	return &TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
