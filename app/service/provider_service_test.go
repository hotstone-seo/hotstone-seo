package service_test

import (
	"context"
	"errors"
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
		}, false)
		require.NoError(t, err)
		require.Equal(t, []byte("some-data"), resp)
	})

	t.Run("WHEN FindOne returns Error", func(t *testing.T) {
		dataSourceRepo.EXPECT().FindOne(ctx, int64(99999)).Return(nil, errors.New("some-error"))
		resp, err := svc.RetrieveData(ctx, service.RetrieveDataRequest{
			DataSourceID: 99999,
		}, false)
		require.EqualError(t, err, "some-error")
		require.Nil(t, resp)
	})

	t.Run("WHEN server errors", func(t *testing.T) {
		dataSourceRepo.EXPECT().FindOne(ctx, int64(99999)).Return(&repository.DataSource{
			Url: "non-existent",
		}, nil)
		resp, err := svc.RetrieveData(ctx, service.RetrieveDataRequest{
			DataSourceID: 99999,
		}, false)
		require.EqualError(t, err, "Get non-existent: unsupported protocol scheme \"\"")
		require.Nil(t, resp)
	})
}

func TestProvider_Tags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		tagRepo        = mock.NewMockTagRepo(ctrl)
		ruleRepo       = mock.NewMockRuleRepo(ctrl)
		dataSourceRepo = mock.NewMockDataSourceRepo(ctrl)

		svc = service.ProviderServiceImpl{
			TagRepo:        tagRepo,
			RuleRepo:       ruleRepo,
			DataSourceRepo: dataSourceRepo,
		}

		ctx = context.Background()
	)
	t.Run("WHEN can't find tag by rule and locale", func(t *testing.T) {
		tagRepo.EXPECT().Find(ctx, repository.TagFilter{RuleID: int64(999), Locale: "en-US"}).Return(nil, errors.New("some-error"))

		tags, err := svc.Tags(ctx, service.ProvideTagsRequest{RuleID: 999, Locale: "en-US"}, false)
		require.EqualError(t, err, "some-error")
		require.Nil(t, tags)
	})
	t.Run("WHEN requesting external data", func(t *testing.T) {
		tagRepo.EXPECT().Find(ctx, repository.TagFilter{RuleID: int64(999), Locale: "en-US"}).
			Return([]*repository.Tag{
				{ID: 1, RuleID: 999, Locale: "en-US", Type: "some-type", Attributes: dbkit.JSON(`{"key1": "value1 {{.Data1}}", "key2{{.Data2}}": "value2"}`), Value: "some-value{{.Data3}}"},
			}, nil)
		dataSourceID := new(int64)
		*dataSourceID = 555
		ruleRepo.EXPECT().FindOne(ctx, int64(999)).
			Return(&repository.Rule{
				ID:           1,
				Name:         "Some rule",
				UrlPattern:   "/some-pattern",
				DataSourceID: dataSourceID,
			}, nil)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte(`{"Data1":"some-data-1","Data2":"some-data-2","Data3":"some-data-3"}`))
		}))
		defer server.Close()
		dataSourceRepo.EXPECT().FindOne(ctx, int64(555)).
			Return(&repository.DataSource{
				ID:   int64(555),
				Name: "Some datasource",
				Url:  server.URL,
			}, nil)
		tags, err := svc.Tags(ctx, service.ProvideTagsRequest{
			RuleID: 999,
			Locale: "en-US",
		}, false)
		require.NoError(t, err)
		require.EqualValues(t, []*service.InterpolatedTag{
			{
				ID:         1,
				RuleID:     999,
				Locale:     "en-US",
				Type:       "some-type",
				Attributes: dbkit.JSON(`{"key1": "value1 some-data-1", "key2some-data-2": "value2"}`),
				Value:      "some-valuesome-data-3",
			},
		}, tags)
	})

}

func TestProvider_Tags_Success(t *testing.T) {
	testcases := []struct {
		attribute         dbkit.JSON
		value             string
		data              interface{}
		expectedAttribute dbkit.JSON
		expectedValue     string
	}{
		{
			attribute: dbkit.JSON(`{"key1": "value1 {{data1}}", "key2{{Data2}}": "value2"}`),
			value:     "some-value{{data3}}",
			data: struct {
				Data1 string
				Data2 string
				Data3 string
			}{
				Data1: "some-data-1",
				Data2: "some-data-2",
				Data3: "some-data-3",
			},
			expectedAttribute: dbkit.JSON(`{"key1": "value1 some-data-1", "key2some-data-2": "value2"}`),
			expectedValue:     "some-valuesome-data-3",
		},
		{
			attribute: dbkit.JSON(`{"key1": "value1 {{data1}}", "key2{{data2}}": "value2"}`),
			value:     "some-value{{data3}}",
			data: map[string]string{
				"data1": "some-data-1",
				"data2": "some-data-2",
				"data3": "some-data-3",
			},
			expectedAttribute: dbkit.JSON(`{"key1": "value1 some-data-1", "key2some-data-2": "value2"}`),
			expectedValue:     "some-valuesome-data-3",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		tagRepo        = mock.NewMockTagRepo(ctrl)
		ruleRepo       = mock.NewMockRuleRepo(ctrl)
		dataSourceRepo = mock.NewMockDataSourceRepo(ctrl)

		svc = service.ProviderServiceImpl{
			TagRepo:        tagRepo,
			RuleRepo:       ruleRepo,
			DataSourceRepo: dataSourceRepo,
		}
	)
	for i, tt := range testcases {
		ctx := context.Background()
		tagRepo.EXPECT().Find(ctx, repository.TagFilter{RuleID: int64(999), Locale: "en-US"}).
			Return([]*repository.Tag{{ID: 1, RuleID: 1, Locale: "en-US", Type: "some-type", Attributes: tt.attribute, Value: tt.value}}, nil)

		tags, err := svc.Tags(ctx, service.ProvideTagsRequest{RuleID: 999, Locale: "en-US", Data: tt.data}, false)
		require.NoError(t, err, i)
		require.EqualValues(t, []*service.InterpolatedTag{
			{ID: 1, RuleID: 1, Locale: "en-US", Type: "some-type", Attributes: tt.expectedAttribute, Value: tt.expectedValue},
		}, tags, i)
	}

}
