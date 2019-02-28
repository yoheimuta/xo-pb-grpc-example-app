package userapp

import (
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expgenproto/userpb"
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expmysql/expmodels"
)

// RegisterUserRequest represents a request for RegisterUser method.
type RegisterUserRequest struct {
	User *expmodels.User
	Auth *expmodels.UserAuth
}

// NewRegisterUserRequestFromPB creates a new RegisterUserRequest.
func (a *App) NewRegisterUserRequestFromPB(
	pb *userpb.RegisterUserRequest,
) *RegisterUserRequest {
	now := a.clock.Now()
	user := &expmodels.User{
		UserID:    pb.UserId,
		CreatedAt: now,
		UpdatedAt: now,
	}
	auth := &expmodels.UserAuth{
		UserID:       pb.UserId,
		Email:        pb.EmailAddress,
		PasswordHash: pb.Password,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	return &RegisterUserRequest{
		User: user,
		Auth: auth,
	}
}

func (r *RegisterUserRequest) userID() string {
	return r.User.UserID
}
