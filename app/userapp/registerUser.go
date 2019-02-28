package userapp

import "context"

// RegisterUser registers a user account.
func (a *App) RegisterUser(
	ctx context.Context,
	req *RegisterUserRequest,
) (string, error) {
	err := a.db.RegisterUser(
		ctx,
		req.User,
		req.Auth,
	)
	if err != nil {
		return "", err
	}
	return a.authTokenGenerator.Generate(
		req.userID(),
		a.clock.Now(),
	)
}
