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

func TestRoleTypeController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	roleTypeSvcMock := mock_service.NewMockRoleTypeService(ctrl)
	roleTypeCntrl := controller.RoleTypeCntrl{
		RoleTypeService: roleTypeSvcMock,
	}

	t.Run("WHEN invalid role type request", func(t *testing.T) {
		_, err := echotest.DoPOST(roleTypeCntrl.Create, "/", `{ "name": ""}`, nil)
		require.EqualError(t, err, "code=400, message=Key: 'RoleType.Name' Error:Field validation for 'Name' failed on the 'required' tag")
	})
	t.Run("WHEN invalid json format", func(t *testing.T) {
		_, err := echotest.DoPOST(roleTypeCntrl.Create, "/", `invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN insert error", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("some-insert-error"))
		_, err := echotest.DoPOST(roleTypeCntrl.Create, "/", `{ "name": "some-name", "modules": {} }`, nil)
		require.EqualError(t, err, "code=422, message=some-insert-error")
	})
	t.Run("WHEN insert success", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(100), nil)
		rr, err := echotest.DoPOST(roleTypeCntrl.Create, "/", `{ "name": "some-name" } `, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":100,\"name\":\"some-name\",\"modules\":null,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestRoleTypeController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	roleTypeSvcMock := mock_service.NewMockRoleTypeService(ctrl)
	roleTypeCntrl := controller.RoleTypeCntrl{
		RoleTypeService: roleTypeSvcMock,
	}
	t.Run("WHEN retrieved error", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(roleTypeCntrl.Find, "/", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(
			[]*repository.RoleType{
				&repository.RoleType{
					ID:      100,
					Name:    "admin",
					Modules: map[string]string{},
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(roleTypeCntrl.Find, "/", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":100,\"name\":\"admin\",\"modules\":{},\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}]\n", rr.Body.String())
	})
}

func TestRoleTypeController_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	roleTypeSvcMock := mock_service.NewMockRoleTypeService(ctrl)
	roleTypeCntrl := controller.RoleTypeCntrl{
		RoleTypeService: roleTypeSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(roleTypeCntrl.FindOne, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=422, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, errors.New("find error"))
		_, err := echotest.DoGET(roleTypeCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN role type is not found", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, sql.ErrNoRows)
		_, err := echotest.DoGET(roleTypeCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=404, message=Not Found")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(
			&repository.RoleType{
				ID:      100,
				Name:    "admin",
				Modules: map[string]string{},
			},
			nil,
		)
		rr, err := echotest.DoGET(roleTypeCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"id\":100,\"name\":\"admin\",\"modules\":{},\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestRoleTypeController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	roleTypeSvcMock := mock_service.NewMockRoleTypeService(ctrl)
	roleTypeCntrl := controller.RoleTypeCntrl{
		RoleTypeService: roleTypeSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoDELETE(roleTypeCntrl.Delete, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Delete(gomock.Any(), int64(100)).Return(errors.New("find error"))
		_, err := echotest.DoDELETE(roleTypeCntrl.Delete, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Delete(gomock.Any(), int64(100)).Return(nil)
		rr, err := echotest.DoDELETE(roleTypeCntrl.Delete, "/", map[string]string{"id": "100"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success delete role type #100\"}\n", rr.Body.String())
	})
}

func TestRoleTypeController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	roleTypeSvcMock := mock_service.NewMockRoleTypeService(ctrl)
	roleTypeCntrl := controller.RoleTypeCntrl{
		RoleTypeService: roleTypeSvcMock,
	}
	t.Run("WHEN body is malformed", func(t *testing.T) {
		_, err := echotest.DoPUT(roleTypeCntrl.Update, "/", "invalid", nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN id is invalid", func(t *testing.T) {
		_, err := echotest.DoPUT(roleTypeCntrl.Update, "/", `{ "id": -1, "name": "admin"}`, nil)
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error after updating", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
		_, err := echotest.DoPUT(roleTypeCntrl.Update, "/", `{ "id": 100, "name": "admin" }`, nil)
		require.EqualError(t, err, "code=500, message=update error")
	})
	t.Run("WHEN successfully update role type", func(t *testing.T) {
		roleTypeSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		rr, err := echotest.DoPUT(roleTypeCntrl.Update, "/", `{ "id": 100, "name": "admin" }`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success update role type #100\"}\n", rr.Body.String())
	})
}
