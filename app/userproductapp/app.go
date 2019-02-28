package userproductapp

import (
	"context"

	"github.com/yoheimuta/xo-pb-example-app/infra/expmysql/expmodels"
)

// DBRepository represents a DB repository.
type DBRepository interface {
	// ListUserProductsByUserID gets a list of user's products using a userID.
	ListUserProductsByUserID(
		_ context.Context,
		userID string,
	) ([]*expmodels.UserProduct, error)
}

// App represents an application managing a user's products.
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
