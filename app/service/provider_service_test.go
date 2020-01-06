package service_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hotstone-seo/hotstone-server/app/repository"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/hotstone-seo/hotstone-server/mock"
)

func TestProvider_RetrieveData(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dataSourceRepo := mock.NewMockDataSourceRepo(ctrl)
	svc := service.ProviderServiceImpl{DataSourceRepo: dataSourceRepo}

	t.Run("Success", func(t *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			w.Write([]byte("some-data"))
		}
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		dataSourceRepo.EXPECT().FindOne(ctx, int64(99999)).Return(&repository.DataSource{
			Url: server.URL,
		}, nil)
		resp, err := svc.RetrieveData(ctx, service.RetrieveDataRequest{
			DataSourceID: 99999,
		})
		require.NoError(t, err)
		body, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, 500, resp.StatusCode)
		require.Equal(t, []byte("some-data"), body)
	})

	t.Run("WHEN FindOne returns Error", func(t *testing.T) {
		dataSourceRepo.EXPECT().FindOne(ctx, int64(99999)).Return(nil, errors.New("some-error"))
		resp, err := svc.RetrieveData(ctx, service.RetrieveDataRequest{
			DataSourceID: 99999,
		})
		require.EqualError(t, err, "some-error")
		require.Nil(t, resp)
	})

	t.Run("WHEN server errors", func(t *testing.T) {
		dataSourceRepo.EXPECT().FindOne(ctx, int64(99999)).Return(&repository.DataSource{
			Url: "non-existent",
		}, nil)
		resp, err := svc.RetrieveData(ctx, service.RetrieveDataRequest{
			DataSourceID: 99999,
		})
		require.EqualError(t, err, "Get non-existent: unsupported protocol scheme \"\"")
		require.Nil(t, resp)
	})
}
