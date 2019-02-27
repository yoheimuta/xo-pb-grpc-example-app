package expmysql

import (
	"context"

	"github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"
)

// ListUserProductsByUserID gets a list of user's products using a userID.
func (c *Client) ListUserProductsByUserID(
	_ context.Context,
	userID string,
) ([]*expmodels.UserProduct, error) {
	return expmodels.UserProductsByUserID(c.db, userID)
}
