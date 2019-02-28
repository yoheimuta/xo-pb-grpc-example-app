package userproductapp

import (
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expgenproto/userproductpb"
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expmysql/expmodels"
)

// ListUserProductsResponse represents a response for ListUserProducts method.
type ListUserProductsResponse struct {
	UserProducts []*expmodels.UserProduct
}

// ToListUserProductResponsePB converts to a ListUserProductResponse pb.
func (r *ListUserProductsResponse) ToListUserProductResponsePB() *userproductpb.ListUserProductsResponse {
	var products []*userproductpb.ListUserProductsResponse_Product
	for _, p := range r.UserProducts {
		products = append(products, &userproductpb.ListUserProductsResponse_Product{
			UserProductId: p.UserProductID,
			Title:         p.Title,
			Description:   p.Description,
		})
	}
	return &userproductpb.ListUserProductsResponse{
		Products: products,
	}
}
