package userapp

import (
	"context"
	"time"

	"github.com/yoheimuta/xo-pb-example-app/infra/expmysql/expmodels"
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

// Clock represents a source of current time.
type Clock interface {
	// Now returns a current time.
	Now() time.Time
}

// App represents an application managing a user's account.
type App struct {
	db    DBRepository
	clock Clock
}

// NewApp creates a new App.
func NewApp(
	db DBRepository,
	clock Clock,
) *App {
	return &App{
		db:    db,
		clock: clock,
	}
}
