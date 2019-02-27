package userproductapp

import (
	"context"
)

// ListUserProducts gets a list of user's products.
func (a *App) ListUserProducts(
	ctx context.Context,
	userID string,
) (*ListUserProductsResponse, error) {
	userProducts, err := a.db.ListUserProductsByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}
	return &ListUserProductsResponse{
		UserProducts: userProducts,
	}, nil
}
