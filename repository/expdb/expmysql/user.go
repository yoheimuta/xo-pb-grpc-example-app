package expmysql

import (
	"context"

	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expmysql/expmodels"
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expsql"
)

// RegisterUser registers a user.
func (c *Client) RegisterUser(
	ctx context.Context,
	user *expmodels.User,
	auth *expmodels.UserAuth,
) error {
	return c.db.WithTx(func(tx expsql.Tx) error {
		err := user.Insert(ctx, tx)
		if err != nil {
			return err
		}

		err = auth.Insert(ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
}
