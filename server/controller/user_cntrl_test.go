package controller_test

import (
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestUserController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userSvcMock := mock_service.NewMockUserService(ctrl)
	userCntrl := controller.UserCntrl{
		UserService: userSvcMock,
	}
	t.Run("WHEN invalid user request", func(t *testing.T) {
		_, err := echotest.DoPOST(userCntrl.Create, "/", `{ "email": ""}`, nil)
		require.EqualError(t, err, "code=400, message=Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag")
	})
	t.Run("WHEN invalid json format", func(t *testing.T) {
		_, err := echotest.DoPOST(userCntrl.Create, "/", `invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN insert error", func(t *testing.T) {
		userSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("some-insert-error"))
		_, err := echotest.DoPOST(userCntrl.Create, "/", `{ "email": "some-name", "role_type_id":1}`, nil)
		require.EqualError(t, err, "code=422, message=some-insert-error")
	})
	t.Run("WHEN insert success", func(t *testing.T) {
		userSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(100), nil)
		rr, err := echotest.DoPOST(userCntrl.Create, "/", `{ "email": "some-name", "role_type_id":1}`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":100,\"email\":\"some-name\",\"role_type_id\":1,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestUserController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userSvcMock := mock_service.NewMockUserService(ctrl)
	userCntrl := controller.UserCntrl{
		UserService: userSvcMock,
	}
	t.Run("WHEN retrieved error", func(t *testing.T) {
		userSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(userCntrl.Find, "/", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		userSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(
			[]*repository.User{
				&repository.User{
					ID:    100,
					Email: "test@tiket.com",
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(userCntrl.Find, "/", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":100,\"email\":\"test@tiket.com\",\"role_type_id\":0,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}]\n", rr.Body.String())
	})
}

func TestUserController_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userSvcMock := mock_service.NewMockUserService(ctrl)
	userCntrl := controller.UserCntrl{
		UserService: userSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(userCntrl.FindOne, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=422, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		userSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, errors.New("find error"))
		_, err := echotest.DoGET(userCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN user is not found", func(t *testing.T) {
		userSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, sql.ErrNoRows)
		_, err := echotest.DoGET(userCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=404, message=Not Found")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		userSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(
			&repository.User{
				ID:    100,
				Email: "test@tiket.com",
			},
			nil,
		)
		rr, err := echotest.DoGET(userCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"id\":100,\"email\":\"test@tiket.com\",\"role_type_id\":0,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestUserController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userSvcMock := mock_service.NewMockUserService(ctrl)
	userCntrl := controller.UserCntrl{
		UserService: userSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoDELETE(userCntrl.Delete, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		userSvcMock.EXPECT().Delete(gomock.Any(), int64(100)).Return(errors.New("find error"))
		_, err := echotest.DoDELETE(userCntrl.Delete, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		userSvcMock.EXPECT().Delete(gomock.Any(), int64(100)).Return(nil)
		rr, err := echotest.DoDELETE(userCntrl.Delete, "/", map[string]string{"id": "100"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success delete user #100\"}\n", rr.Body.String())
	})
}

func TestUserController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userSvcMock := mock_service.NewMockUserService(ctrl)
	userCntrl := controller.UserCntrl{
		UserService: userSvcMock,
	}
	t.Run("WHEN body is malformed", func(t *testing.T) {
		_, err := echotest.DoPUT(userCntrl.Update, "/", "invalid", nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN id is invalid", func(t *testing.T) {
		_, err := echotest.DoPUT(userCntrl.Update, "/", `{ "id": -1, "email": "test@tiket.com"}`, nil)
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error after updating", func(t *testing.T) {
		userSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
		_, err := echotest.DoPUT(userCntrl.Update, "/", `{ "id": 100, "email": "test@tiket.com" }`, nil)
		require.EqualError(t, err, "code=500, message=update error")
	})
	t.Run("WHEN successfully update user", func(t *testing.T) {
		userSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		rr, err := echotest.DoPUT(userCntrl.Update, "/", `{ "id": 100, "email": "test@tiket.com" }`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success update user #100\"}\n", rr.Body.String())
	})
}
