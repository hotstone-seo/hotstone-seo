package service_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/hotstone-seo/hotstone-seo/server/mock_metric"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
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

	urlServiceMock := mock_service.NewMockURLService(ctrl)
	ruleMatchingMock := mock_metric.NewMockRuleMatchingRepo(ctrl)

	svc := service.ProviderServiceImpl{
		URLService:       urlServiceMock,
		RuleMatchingRepo: ruleMatchingMock,
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
				urlServiceMock.EXPECT().Match("some-path").Return(int64(-1), nil)
				ruleMatchingMock.EXPECT().Insert(gomock.Any(), gomock.Any())
			},
			vals:        parseQuery("_path=some-path"),
			expectedErr: "No rule match: some-path",
		},
		{
			testName: "match",
			pre: func() {
				urlServiceMock.EXPECT().Match("some-path").Return(int64(1), map[string]string{"hello": "world"})
				ruleMatchingMock.EXPECT().Insert(gomock.Any(), gomock.Any())
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
