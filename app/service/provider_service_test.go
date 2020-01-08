package service_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/app/service"
	"github.com/hotstone-seo/hotstone-seo/mock"
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

func TestProvider_Tags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagRepo := mock.NewMockTagRepo(ctrl)
	svc := service.ProviderServiceImpl{TagRepo: tagRepo}
	ctx := context.Background()
	t.Run("WHEN can't find tag by rule and locale", func(t *testing.T) {
		tagRepo.EXPECT().FindByRuleAndLocale(ctx, int64(999), int64(888)).
			Return(nil, errors.New("some-error"))
		tags, err := svc.Tags(ctx, service.ProvideTagsRequest{
			RuleID:   999,
			LocaleID: 888,
		})
		require.EqualError(t, err, "some-error")
		require.Nil(t, tags)
	})
	t.Run("WHEN success", func(t *testing.T) {
		tagRepo.EXPECT().FindByRuleAndLocale(ctx, int64(999), int64(888)).
			Return([]*repository.Tag{
				{
					ID:         1,
					RuleID:     1,
					LocaleID:   1,
					Type:       "some-type",
					Attributes: dbkit.JSON(`{"key1": "value1 {{.Data1}}", "key2{{.Data2}}": "value2"}`),
					Value:      "some-value{{.Data3}}",
				},
			}, nil)
		tags, err := svc.Tags(ctx, service.ProvideTagsRequest{
			RuleID:   999,
			LocaleID: 888,
			Data: struct {
				Data1 string
				Data2 string
				Data3 string
			}{"some-data-1", "some-data-2", "some-data-3"},
		})
		require.NoError(t, err)
		require.EqualValues(t, []*service.InterpolatedTag{
			{
				ID:         1,
				RuleID:     1,
				LocaleID:   1,
				Type:       "some-type",
				Attributes: dbkit.JSON(`{"key1": "value1 some-data-1", "key2some-data-2": "value2"}`),
				Value:      "some-valuesome-data-3",
			},
		}, tags)

	})
}
