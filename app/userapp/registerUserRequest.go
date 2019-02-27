package userapp

import "github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"

// RegisterUserRequest represents a request for RegisterUser method.
type RegisterUserRequest struct {
	User *expmodels.User
	Auth *expmodels.UserAuth
}
