package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestDataSourceController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dataSourceSvcMock := mock_service.NewMockDataSourceService(ctrl)
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
