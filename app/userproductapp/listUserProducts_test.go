package userproductapp_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/yoheimuta/xo-example-app/app/userproductapp"
	"github.com/yoheimuta/xo-example-app/infra/expdep_test"
	"github.com/yoheimuta/xo-example-app/infra/expmysql/expmodels"
)

func TestApp_ListUserProducts(t *testing.T) {
	t.Parallel()

	dep, err := expdep_test.NewDep()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer dep.Close(t)

	userID := "testUserID"
	prepared, err := prepareListUserProducts(
		userID,
		dep,
	)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	for _, test := range []struct {
		name        string
		inputUserID string
		wantResp    *userproductapp.ListUserProductsResponse
	}{
		{
			name:        "Got an empty list",
			inputUserID: "NotFoundID",
			wantResp: &userproductapp.ListUserProductsResponse{
				UserProducts: []*expmodels.UserProduct{},
			},
		},
		{
			name:        "Got a list containing 2 entities",
			inputUserID: userID,
			wantResp: &userproductapp.ListUserProductsResponse{
				UserProducts: prepared,
			},
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := dep.UserProductApp().ListUserProducts(
				context.Background(),
				test.inputUserID,
			)
			if err != nil {
				t.Errorf("got err %v", err)
				return
			}

			if !reflect.DeepEqual(got, test.wantResp) {
				t.Errorf("got %v, but want %v", got, test.wantResp)
			}
		})
	}
}

func prepareListUserProducts(
	userID string,
	dep *expdep_test.Dep,
) (
	[]*expmodels.UserProduct,
	error,
) {
	now := dep.Now()
	user := &expmodels.User{
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	ctx := context.Background()
	err := user.Insert(ctx, dep.RawDB())
	if err != nil {
		return nil, err
	}

	userProducts := []*expmodels.UserProduct{
		{
			UserProductID: "id1",
			UserID:        userID,
			Title:         "title1",
			Description:   "desc1",
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			UserProductID: "id2",
			UserID:        userID,
			Title:         "title2",
			Description:   "desc2",
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}
	for _, p := range userProducts {
		err = p.Insert(ctx, dep.RawDB())
		if err != nil {
			return nil, err
		}
	}
	return userProducts, nil
}
