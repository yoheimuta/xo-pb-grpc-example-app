package userapp

import (
	"context"

	"github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"
)

// DBRepository represents a DB repository.
type DBRepository interface {
	// RegisterUser registers a user.
	RegisterUser(
		ctx context.Context,
		user *expmodels.User,
		auth *expmodels.UserAuth,
	) error
}

// App represents an application managing a user's account.
type App struct {
	db DBRepository
}

// NewApp creates a new App.
func NewApp(
	db DBRepository,
) *App {
	return &App{
		db: db,
	}
}
