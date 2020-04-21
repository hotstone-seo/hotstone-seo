package service_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
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
			ID:   92,
			Type: "meta",
			Attributes: map[string]string{
				"name":    "description",
				"content": "This year is {{year}}",
			},
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
			ID:   92,
			Type: "meta",
			Attributes: map[string]string{
				"name":    "description",
				"content": "This year is 2020",
			},
		},
	}

	rule777_noDS       = &repository.Rule{ID: 777}
	tags_rule777_en_US = []*repository.Tag{
		{ID: 71},
		{ID: 72},
	}
)

type fetchTestCase struct {
	testName    string
	server      *httptest.Server
	pre         func(fetchTestCase)
	vals        url.Values
	expected    []*service.ITag
	expectedErr string
}

func TestProviderService_FetchTagsWithCache(t *testing.T) {
	testRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer testRedis.Close()

	b, _ := json.Marshal(itags_rule999_en_US)

	key := "_locale=en_US&_rule=999"
	testRedis.Set(key, string(b))
	testRedis.Set(key+":time", time.Now().UTC().Add(10*time.Second).Format(time.RFC1123))
	testRedis.SetTTL(key, 10*time.Second)
	testRedis.SetTTL(key+":time", 10*time.Second)

	svc := service.ProviderServiceImpl{
		Redis: redis.NewClient(&redis.Options{Addr: testRedis.Addr()}),
	}

	tags, err := svc.FetchTagsWithCache(context.Background(), parseQuery(`_rule=999&_locale=en_US`), &cachekit.Pragma{})
	require.NoError(t, err)
	require.Equal(t, itags_rule999_en_US, tags)
}

func TestProviderService2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dsmock := mock_repository.NewMockDataSourceRepo(ctrl)
	rulemock := mock_repository.NewMockRuleRepo(ctrl)
	tagmock := mock_repository.NewMockTagRepo(ctrl)
	ctx := context.Background()

	svc := service.ProviderServiceImpl{
		DataSourceRepo: dsmock,
		RuleRepo:       rulemock,
		TagRepo:        tagmock,
	}

	testCases := []fetchTestCase{
		{
			testName:    "WHEN no ID",
			vals:        parseQuery(`_locale=en_US`),
			expectedErr: "Validation: Missing url param for `ID`",
		},
		{
			testName:    "WHEN ID is not number",
			vals:        parseQuery(`_rule=qwery&_locale=en_US`),
			expectedErr: "Validation: Missing url param for `ID`",
		},
		{
			testName:    "WHEN no locale",
			vals:        parseQuery(`_rule=999`),
			expectedErr: "Validation: Missing query param for `Locale`",
		},
		{
			testName: "WHEN rule not exist",
			pre: func(fetchTestCase) {
				rulemock.EXPECT().FindOne(ctx, int64(999)).Return(nil, sql.ErrNoRows)
			},
			vals:        parseQuery(`_rule=999&_locale=en_US`),
			expectedErr: sql.ErrNoRows.Error(),
		},
		{
			testName: "WHEN database error in get tags",
			pre: func(fetchTestCase) {
				rulemock.EXPECT().FindOne(ctx, int64(999)).Return(rule999, nil)
				tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").Return(nil, errors.New("some-error"))
			},
			vals:        parseQuery(`_rule=999&_locale=en_US`),
			expectedErr: "Find-Tags: some-error",
		},
		{
			testName: "WHEN no tags",
			pre: func(fetchTestCase) {
				rulemock.EXPECT().FindOne(ctx, int64(777)).Return(rule777_noDS, nil)
				tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(777), "en_US").Return([]*repository.Tag{}, nil)
			},
			vals:     parseQuery(`_rule=777&_locale=en_US`),
			expected: []*service.ITag{},
		},
		{
			testName: "WHEN rule has no datasource",
			pre: func(fetchTestCase) {
				rulemock.EXPECT().FindOne(ctx, int64(777)).Return(rule777_noDS, nil)
				tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(777), "en_US").Return(tags_rule777_en_US, nil)
			},
			vals: parseQuery(`_rule=777&_locale=en_US`),
			expected: func() (tags []*service.ITag) {
				for _, tag := range tags_rule777_en_US {
					itag := service.ITag(*tag)
					tags = append(tags, &itag)
				}
				return
			}(),
		},
		{
			testName: "WHEN data source is not exist",
			pre: func(fetchTestCase) {
				rulemock.EXPECT().FindOne(ctx, int64(999)).Return(rule999, nil)
				tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").Return(tags_rule999_en_US, nil)
				dsmock.EXPECT().FindOne(ctx, ds666_id).Return(nil, fmt.Errorf("some-error"))
			},
			vals:        parseQuery(`_rule=999&_locale=en_US`),
			expectedErr: "DataSource: some-error",
		},
		{
			testName: "WHEN can't call to data source",
			pre: func(fetchTestCase) {
				rulemock.EXPECT().FindOne(ctx, int64(999)).Return(rule999, nil)
				tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").Return(tags_rule999_en_US, nil)
				dsmock.EXPECT().FindOne(ctx, ds666_id).Return(&repository.DataSource{URL: "bad-url"}, nil)
			},
			vals:        parseQuery(`_rule=999&_locale=en_US`),
			expectedErr: "Call: Get \"bad-url\": unsupported protocol scheme \"\"",
		},
		{
			testName: "WHEN data source return no json",
			server:   dummyServer(`{bad-json`),
			pre: func(tt fetchTestCase) {
				rulemock.EXPECT().FindOne(ctx, int64(999)).Return(rule999, nil)
				tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").Return(tags_rule999_en_US, nil)
				dsmock.EXPECT().FindOne(ctx, ds666_id).Return(&repository.DataSource{URL: tt.server.URL}, nil)
			},
			vals:        parseQuery(`_rule=999&_locale=en_US`),
			expectedErr: "JSON: invalid character 'b' looking for beginning of object key string",
		},
		{
			testName: "success",
			server:   dummyServer(ds666_response),
			pre: func(tt fetchTestCase) {
				ds := &repository.DataSource{
					ID:  ds666_id,
					URL: tt.server.URL,
				}

				rulemock.EXPECT().FindOne(ctx, int64(999)).Return(rule999, nil)
				tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").Return(tags_rule999_en_US, nil)
				dsmock.EXPECT().FindOne(ctx, ds666_id).Return(ds, nil)
			},
			vals:     parseQuery(`_rule=999&_locale=en_US`),
			expected: itags_rule999_en_US,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			defer func() {
				if tt.server != nil {
					tt.server.Close()
				}
			}()

			if tt.pre != nil {
				tt.pre(tt)
			}

			tags, err := svc.FetchTags(ctx, tt.vals)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, tags)
		})
	}

}

func parseQuery(s string) url.Values {
	vals, _ := url.ParseQuery(s)
	return vals
}

func dummyServer(msg string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	}))
}

func TestConvertToParam(t *testing.T) {
	testcases := []struct {
		desc     string
		query    string
		expected map[string]string
	}{
		{
			desc:  "normal query string",
			query: "locale=en_US&some-field=some-value",
			expected: map[string]string{
				"locale":     "en_US",
				"some-field": "some-value",
			},
		},
		{
			desc:  "some field is empty",
			query: "locale=en_US&some-field=",
			expected: map[string]string{
				"locale":     "en_US",
				"some-field": "",
			},
		},
	}

	for _, tt := range testcases {
		vals, _ := url.ParseQuery(tt.query)
		require.Equal(t, tt.expected, service.ConvertToParams(vals))
	}
}

func TestUnmarshalData(t *testing.T) {
	testcases := []struct {
		desc        string
		data        string
		expected    interface{}
		expectedErr string
	}{
		{
			desc: "json object",
			data: `{"hello": "world"}`,
			expected: map[string]interface{}{
				"hello": "world",
			},
		},
		{
			desc: "json array",
			data: `[{"hello": "world"}]`,
			expected: map[string]interface{}{
				"hello": "world",
			},
		},
		{
			desc:     "empty json array",
			data:     `[]`,
			expected: map[string]interface{}{},
		},
	}

	for _, tt := range testcases {
		v, err := service.UnmarshalData([]byte(tt.data))
		if tt.expectedErr != "" {
			require.EqualError(t, err, tt.expectedErr)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, tt.expected, v)
	}
}
