package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/service_mock"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestDataSourceController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dataSourceSvcMock := service_mock.NewMockDataSourceService(ctrl)
	dataSourceCntrl := controller.DataSourceCntrl{
		DataSourceService: dataSourceSvcMock,
	}
	t.Run("WHEN invalid json format", func(t *testing.T) {
		_, err := echotest.DoPOST(dataSourceCntrl.Create, "/", `invalid json`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN insert error", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("some-insert-error"))
		_, err := echotest.DoPOST(dataSourceCntrl.Create, "/", `{ "name": "data-source-name", "url": "http://data-source-url"}`, nil)
		require.EqualError(t, err, "code=422, message=some-insert-error")
	})
	t.Run("WHEN insert success", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(100), nil)
		rr, err := echotest.DoPOST(dataSourceCntrl.Create, "/", `{ "name": "data-source-name", "url": "http://data-source-url"}`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":100,\"name\":\"data-source-name\",\"url\":\"http://data-source-url\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestDataSourceController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dataSourceSvcMock := service_mock.NewMockDataSourceService(ctrl)
	dataSourceCntrl := controller.DataSourceCntrl{
		DataSourceService: dataSourceSvcMock,
	}
	t.Run("WHEN retrieved error", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Find(gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(dataSourceCntrl.Find, "/", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Find(gomock.Any()).Return(
			[]*repository.DataSource{
				&repository.DataSource{
					ID:   100,
					Name: "Airport Data Source",
					URL:  "dssource1.com",
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(dataSourceCntrl.Find, "/", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":100,\"name\":\"Airport Data Source\",\"url\":\"dssource1.com\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}]\n", rr.Body.String())
	})
}

func TestDataSourceController_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dataSourceSvcMock := service_mock.NewMockDataSourceService(ctrl)
	dataSourceCntrl := controller.DataSourceCntrl{
		DataSourceService: dataSourceSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(dataSourceCntrl.FindOne, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, errors.New("find error"))
		_, err := echotest.DoGET(dataSourceCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN Data Source is not found", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(nil, nil)
		_, err := echotest.DoGET(dataSourceCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=404, message=DataSource#100 not found")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().FindOne(gomock.Any(), int64(100)).Return(
			&repository.DataSource{
				ID:   100,
				Name: "Airport Data Source",
				URL:  "dssource1.com",
			},
			nil,
		)
		rr, err := echotest.DoGET(dataSourceCntrl.FindOne, "/", map[string]string{"id": "100"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"id\":100,\"name\":\"Airport Data Source\",\"url\":\"dssource1.com\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestDataSourceController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dataSourceSvcMock := service_mock.NewMockDataSourceService(ctrl)
	dataSourceCntrl := controller.DataSourceCntrl{
		DataSourceService: dataSourceSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoDELETE(dataSourceCntrl.Delete, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Delete(gomock.Any(), int64(100)).Return(errors.New("find error"))
		_, err := echotest.DoDELETE(dataSourceCntrl.Delete, "/", map[string]string{"id": "100"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Delete(gomock.Any(), int64(100)).Return(nil)
		rr, err := echotest.DoDELETE(dataSourceCntrl.Delete, "/", map[string]string{"id": "100"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success delete data_source #100\"}\n", rr.Body.String())
	})
}

func TestDataSourceController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dataSourceSvcMock := service_mock.NewMockDataSourceService(ctrl)
	dataSourceCntrl := controller.DataSourceCntrl{
		DataSourceService: dataSourceSvcMock,
	}
	t.Run("WHEN body is malformed", func(t *testing.T) {
		_, err := echotest.DoPUT(dataSourceCntrl.Update, "/", "invalid", nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN id is invalid", func(t *testing.T) {
		_, err := echotest.DoPUT(dataSourceCntrl.Update, "/", `{ "id": -1, "name": "Airport Rule", "url_pattern": "/airport"}`, nil)
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error after updating", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
		_, err := echotest.DoPUT(dataSourceCntrl.Update, "/", `{ "id": 100, "name": "Airport Data Source", "url": "dssource1.com" }`, nil)
		require.EqualError(t, err, "code=500, message=update error")
	})
	t.Run("WHEN successfully update data source", func(t *testing.T) {
		dataSourceSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		rr, err := echotest.DoPUT(dataSourceCntrl.Update, "/", `{ "id": 100, "name": "Airport Data Source", "url": "dssource1.com" }`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success update data_source #100\"}\n", rr.Body.String())
	})
}
