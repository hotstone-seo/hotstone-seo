package service_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtype"
	"github.com/hotstone-seo/hotstone-seo/server/mock_repository"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
)

var (
	ds666_id int64 = 666

	rule999 = &repository.Rule{
		ID:           999,
		DataSourceID: &ds666_id,
	}
	tags_rule999_en_US = []*repository.Tag{
		{
			ID:    91,
			Type:  "title",
			Value: "Page for {{name}}",
		},
		{
			ID:         92,
			Type:       "meta",
			Attributes: dbtype.JSON(`{"name":"description", "content": "This year is {{year}}"}`),
		},
	}
	ds666_response      string = `{"name":"covid19", "year": 2020}`
	itags_rule999_en_US        = []*service.ITag{
		{
			ID:    91,
			Type:  "title",
			Value: "Page for covid19",
		},
		{
			ID:         92,
			Type:       "meta",
			Attributes: dbtype.JSON(`{"name":"description", "content": "This year is 2020"}`),
		},
	}

	rule777_noDS       = &repository.Rule{ID: 777}
	tags_rule777_en_US = []*repository.Tag{
		{ID: 71},
		{ID: 72},
	}
)

func TestProviderService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dsmock := mock_repository.NewMockDataSourceRepo(ctrl)
	rulemock := mock_repository.NewMockRuleRepo(ctrl)
	tagmock := mock_repository.NewMockTagRepo(ctrl)

	svc := service.ProviderServiceImpl{
		DataSourceRepo: dsmock,
		RuleRepo:       rulemock,
		TagRepo:        tagmock,
	}

	ts := dummyServer(ds666_response)
	defer ts.Close()

	ctx := context.Background()

	rulemock.EXPECT().FindOne(ctx, int64(999)).
		Return(rule999, nil)

	tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").
		Return(tags_rule999_en_US, nil)

	dsmock.EXPECT().FindOne(ctx, ds666_id).
		Return(&repository.DataSource{
			ID:  ds666_id,
			URL: ts.URL,
		}, nil)

	tags, err := svc.FetchTags(ctx, int64(999), "en_US")
	require.NoError(t, err)
	require.Equal(t, itags_rule999_en_US, tags)

}

func TestProviderService_FetchTags_When(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dsmock := mock_repository.NewMockDataSourceRepo(ctrl)
	rulemock := mock_repository.NewMockRuleRepo(ctrl)
	tagmock := mock_repository.NewMockTagRepo(ctrl)

	svc := service.ProviderServiceImpl{
		DataSourceRepo: dsmock,
		RuleRepo:       rulemock,
		TagRepo:        tagmock,
	}

	ctx := context.Background()

	t.Run("WHEN rule not exist", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(999)).
			Return(nil, sql.ErrNoRows)

		_, err := svc.FetchTags(ctx, int64(999), "en_US")

		// THEN return error
		require.Equal(t, sql.ErrNoRows, err)
	})

	t.Run("WHEN database error in get tags ", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(999)).
			Return(rule999, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").
			Return(nil, errors.New("some-error"))

		_, err := svc.FetchTags(ctx, int64(999), "en_US")

		// THEN return error
		require.EqualError(t, err, "Find-Tags: some-error")
	})

	t.Run("WHEN no tags", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(777)).
			Return(rule777_noDS, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(777), "en_US").
			Return([]*repository.Tag{}, nil)

		tags, err := svc.FetchTags(ctx, int64(777), "en_US")
		require.NoError(t, err)

		// THEN return empty list
		require.Empty(t, tags)
	})

	t.Run("WHEN rule has no datasource", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(777)).
			Return(rule777_noDS, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(777), "en_US").
			Return(tags_rule777_en_US, nil)

		tags, err := svc.FetchTags(ctx, int64(777), "en_US")
		require.NoError(t, err)

		// THEN return tags same with the original without interpolation
		for i, tag := range tags {
			og := service.ITag(*tags_rule777_en_US[i])
			require.Equal(t, &og, tag)
		}
	})

	t.Run("WHEN data source is not exist", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(999)).
			Return(rule999, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").
			Return(tags_rule999_en_US, nil)
		dsmock.EXPECT().FindOne(ctx, ds666_id).
			Return(nil, fmt.Errorf("some-error"))

		_, err := svc.FetchTags(ctx, int64(999), "en_US")

		// THEN return error
		require.EqualError(t, err, "DataSource: some-error")
	})

	t.Run("WHEN can't call to data source", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(999)).
			Return(rule999, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").
			Return(tags_rule999_en_US, nil)
		dsmock.EXPECT().FindOne(ctx, ds666_id).
			Return(&repository.DataSource{URL: "bad-url"}, nil)

		_, err := svc.FetchTags(ctx, int64(999), "en_US")

		// THEN return error
		require.EqualError(t, err, "Call: Get \"bad-url\": unsupported protocol scheme \"\"")
	})

	t.Run("WHEN data source return no json", func(t *testing.T) {
		ts := dummyServer(`{bad-json`)
		defer ts.Close()

		rulemock.EXPECT().FindOne(ctx, int64(999)).
			Return(rule999, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").
			Return(tags_rule999_en_US, nil)
		dsmock.EXPECT().FindOne(ctx, ds666_id).
			Return(&repository.DataSource{URL: ts.URL}, nil)

		_, err := svc.FetchTags(ctx, int64(999), "en_US")

		// THEN return error
		require.EqualError(t, err, "JSON: invalid character 'b' looking for beginning of object key string")
	})

}

func dummyServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, response)
	}))
}
