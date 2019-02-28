package expmysql

import (
	"context"

	"github.com/yoheimuta/xo-pb-example-app/infra/expmysql/expmodels"
)

// ListUserProductsByUserID gets a list of user's products using a userID.
func (c *Client) ListUserProductsByUserID(
	ctx context.Context,
	userID string,
) ([]*expmodels.UserProduct, error) {
	return expmodels.UserProductsByUserID(ctx, c.db, userID)
}
