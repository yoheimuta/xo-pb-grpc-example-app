package userproductapp

import "github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"

type ListUserProductsResponse struct {
	UserProducts []*expmodels.UserProduct
}
