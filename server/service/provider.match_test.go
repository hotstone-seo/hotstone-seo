package service_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/hotstone-seo/hotstone-seo/mock_urlstore"
	"github.com/hotstone-seo/hotstone-seo/urlstore"

	"github.com/hotstone-seo/hotstone-seo/server/mock_metric"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/service"
)

type matchTestCase struct {
	testName    string
	vals        url.Values
	pre         func()
	expected    *service.MatchResponse
	expectedErr string
}

func TestProvider_Match(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_urlstore.NewMockStore(ctrl)
	mockRuleMatching := mock_metric.NewMockRuleMatchingRepo(ctrl)

	svc := service.ProviderServiceImpl{
		Store:            mockStore,
		RuleMatchingRepo: mockRuleMatching,
	}

	ctx := context.Background()

	testcases := []matchTestCase{
		{
			testName:    "validate path",
			expectedErr: "Validation: _path can't empty",
		},
		{
			testName: "not match",
			pre: func() {
				mockStore.EXPECT().Get("some-path").Return(nil, nil)
				mockRuleMatching.EXPECT().Insert(gomock.Any(), gomock.Any())
			},
			vals:     parseQuery("_path=some-path"),
			expected: &service.MatchResponse{},
		},
		{
			testName: "match",
			pre: func() {
				mockStore.EXPECT().Get("some-path").Return("1", urlstore.NewParameter([]string{"hello"}, []string{"world"}))
				mockRuleMatching.EXPECT().Insert(gomock.Any(), gomock.Any())
			},
			vals: parseQuery("_path=some-path"),
			expected: &service.MatchResponse{
				RuleID:    1,
				PathParam: map[string]string{"hello": "world"},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			if tt.pre != nil {
				tt.pre()
			}
			res, err := svc.Match(ctx, tt.vals)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, res)

			time.Sleep(time.Millisecond) // waiting to submit the metric
		})
	}

}
