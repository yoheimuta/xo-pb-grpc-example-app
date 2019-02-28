package userproductapp

import "github.com/yoheimuta/xo-pb-example-app/infra/expmysql/expmodels"

// ListUserProductsResponse represents a response for ListUserProducts method.
type ListUserProductsResponse struct {
	UserProducts []*expmodels.UserProduct
}
