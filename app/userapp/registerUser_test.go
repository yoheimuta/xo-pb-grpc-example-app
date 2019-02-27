package userapp_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"

	"github.com/yoheimuta/xo-example-app/app/userapp"

	"github.com/yoheimuta/xo-example-app/infra/expdep_test"
)

func TestApp_RegisterUser(t *testing.T) {
	t.Parallel()

	dep, err := expdep_test.NewDep()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer dep.Close()

	userID := "UUIDv4"
	now := time.Now().UTC().Truncate(time.Second)

	for _, test := range []struct {
		name string
		req  *userapp.RegisterUserRequest
	}{
		{
			name: "Register a user",
			req: &userapp.RegisterUserRequest{
				User: &expmodels.User{
					UserID:    userID,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Auth: &expmodels.UserAuth{
					UserID:       userID,
					Email:        "exp@example.com",
					PasswordHash: "hash",
					CreatedAt:    now,
					UpdatedAt:    now,
				},
			},
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := dep.UserApp().RegisterUser(
				context.Background(),
				test.req,
			)
			if err != nil {
				t.Errorf("got err %v", err)
				return
			}

			db, err := dep.OpenRawDB()
			if err != nil {
				t.Errorf("got err %v", err)
				return
			}
			defer func() {
				_ = db.Close()
			}()

			gotUser, err := expmodels.UserByUserID(db, test.req.User.UserID)
			if err != nil {
				t.Errorf("got err %v", err)
			}
			if !reflect.DeepEqual(gotUser, test.req.User) {
				t.Errorf("got %v, but want %v", gotUser, test.req.User)
			}

			gotAuth, err := expmodels.UserAuthByUserID(db, test.req.User.UserID)
			if err != nil {
				t.Errorf("got err %v", err)
			}
			if !reflect.DeepEqual(gotAuth, test.req.Auth) {
				t.Errorf("got %v, but want %v", gotAuth, test.req.Auth)
			}
		})
	}

}
