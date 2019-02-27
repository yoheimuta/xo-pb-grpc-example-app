package expmysql

import (
	"github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"
	"github.com/yoheimuta/xo-example-app/infra/expsql"
)

// RegisterUser registers a user.
func (c *Client) RegisterUser(
	user *expmodels.User,
	auth *expmodels.UserAuth,
) error {
	return c.db.WithTx(func(tx expsql.Tx) error {
		err := user.Insert(tx)
		if err != nil {
			return err
		}

		err = auth.Insert(tx)
		if err != nil {
			return err
		}
		return nil
	})
}
