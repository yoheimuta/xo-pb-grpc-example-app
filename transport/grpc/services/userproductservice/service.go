package userproductservice

import (
	"context"

	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expgenproto/userproductpb"

	"github.com/yoheimuta/xo-pb-grpc-example-app/app/userproductapp"
)

// Service represents a service about user's products.
type Service struct {
	app *userproductapp.App
}

// NewService creates a new Service.
func NewService(
	app *userproductapp.App,
) *Service {
	return &Service{
		app: app,
	}
}

// ListUserProducts registers an user.
func (s *Service) ListUserProducts(ctx context.Context, _ *userproductpb.ListUserProductsRequest) (*userproductpb.ListUserProductsResponse, error) {
	appResp, err := s.app.ListUserProducts(
		ctx,
		"TODO",
	)
	if err != nil {
		return nil, err
	}
	return appResp.ToListUserProductResponsePB(), nil
}
