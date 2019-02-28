package userapp

import (
	"context"
	"time"

	"github.com/yoheimuta/xo-pb-grpc-example-app/domain/authtoken"

	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expmysql/expmodels"
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
	db                 DBRepository
	clock              Clock
	authTokenGenerator *authtoken.Generator
}

// NewApp creates a new App.
func NewApp(
	db DBRepository,
	clock Clock,
	authTokenGenerator *authtoken.Generator,
) *App {
	return &App{
		db:                 db,
		clock:              clock,
		authTokenGenerator: authTokenGenerator,
	}
}
