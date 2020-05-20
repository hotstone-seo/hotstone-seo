package provider_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/provider"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore_mock"

	"github.com/hotstone-seo/hotstone-seo/internal/analyt_mock"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

type matchTestCase struct {
	testName    string
	vals        url.Values
	pre         func()
	expected    *provider.MatchResponse
	expectedErr string
}

func TestProvider_Match(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := urlstore_mock.NewMockStore(ctrl)
	mockRuleMatching := analyt_mock.NewMockRuleMatchingRepo(ctrl)

	svc := provider.ServiceImpl{
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
			expected: &provider.MatchResponse{},
		},
		{
			testName: "match",
			pre: func() {
				mockStore.EXPECT().Get("some-path").Return("1", urlstore.NewParameter([]string{"hello"}, []string{"world"}))
				mockRuleMatching.EXPECT().Insert(gomock.Any(), gomock.Any())
			},
			vals: parseQuery("_path=some-path"),
			expected: &provider.MatchResponse{
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
