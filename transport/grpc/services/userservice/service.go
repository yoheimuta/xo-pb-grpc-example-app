package userservice

import (
	"context"

	"github.com/yoheimuta/xo-pb-grpc-example-app/app/userapp"
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expgenproto/userpb"
)

// Service represents a service about an user's account.
type Service struct {
	app *userapp.App
}

// NewService creates a new Service.
func NewService(
	app *userapp.App,
) *Service {
	return &Service{
		app: app,
	}
}

// RegisterUser registers an user.
func (s *Service) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	appReq := s.app.NewRegisterUserRequestFromPB(req)
	token, err := s.app.RegisterUser(
		ctx,
		appReq,
	)
	if err != nil {
		return nil, err
	}
	return &userpb.RegisterUserResponse{
		AuthToken: token,
	}, nil
}
