package controller_test

import (
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service_mock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestModuleController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	moduleSvcMock := service_mock.NewMockModuleService(ctrl)
	moduleCntrl := controller.ModuleCntrl{
		ModuleService: moduleSvcMock,
	}
	t.Run("WHEN retrieved error", func(t *testing.T) {
		moduleSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(moduleCntrl.Find, "/", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		moduleSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(
			[]*repository.Module{
				&repository.Module{
					ID:      100,
					Name:    "Rule",
					Path:    "rules",
					Pattern: "rules*",
					Label:   "Rules",
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(moduleCntrl.Find, "/", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":100,\"name\":\"Rule\",\"path\":\"rules\",\"pattern\":\"rules*\",\"label\":\"Rules\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}]\n", rr.Body.String())
	})
}

func TestModuleController_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	moduleSvcMock := service_mock.NewMockModuleService(ctrl)
	moduleCntrl := controller.ModuleCntrl{
		ModuleService: moduleSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(moduleCntrl.FindOne, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=422, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		moduleSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, errors.New("find error"))
		_, err := echotest.DoGET(moduleCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN user is not found", func(t *testing.T) {
		moduleSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, sql.ErrNoRows)
		_, err := echotest.DoGET(moduleCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=404, message=Not Found")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		moduleSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(
			&repository.Module{
				ID:      100,
				Name:    "Rule",
				Path:    "rules",
				Pattern: "rules*",
				Label:   "Rules",
			},
			nil,
		)
		rr, err := echotest.DoGET(moduleCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"id\":100,\"name\":\"Rule\",\"path\":\"rules\",\"pattern\":\"rules*\",\"label\":\"Rules\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}
