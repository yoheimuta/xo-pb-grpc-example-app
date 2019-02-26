package expmysql

import "github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"

// CreateUser creates a user.
func (c *Client) CreateUser(
	user *expmodels.User,
) error {
	return user.Insert(c.db)
}
