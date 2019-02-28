package userproductapp

import (
	"github.com/yoheimuta/xo-pb-example-app/infra/expgenproto/userproductpb"
	"github.com/yoheimuta/xo-pb-example-app/infra/expmysql/expmodels"
)

// ListUserProductsResponse represents a response for ListUserProducts method.
type ListUserProductsResponse struct {
	UserProducts []*expmodels.UserProduct
}

// ToListUserProductResponsePB converts to a ListUserProductResponse pb.
func (r *ListUserProductsResponse) ToListUserProductResponsePB() *userproductpb.ListUserProductResponse {
	var products []*userproductpb.ListUserProductResponse_Product
	for _, p := range r.UserProducts {
		products = append(products, &userproductpb.ListUserProductResponse_Product{
			UserProductId: p.UserProductID,
			Title:         p.Title,
			Description:   p.Description,
		})
	}
	return &userproductpb.ListUserProductResponse{
		Products: products,
	}
}
