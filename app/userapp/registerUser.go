package userapp

import "context"

// RegisterUser registers a user account.
func (a *App) RegisterUser(
	ctx context.Context,
	req *RegisterUserRequest,
) error {
	return a.db.RegisterUser(
		ctx,
		req.user,
		req.auth,
	)
}
