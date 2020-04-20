package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/hotstone-seo/hotstone-seo/server/mock_metric"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/hotstone-seo/hotstone-seo/server/service"
)

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

	t.Run("WHEN no match", func(t *testing.T) {
		urlServiceMock.EXPECT().
			Match("some-path").
			Return(int64(-1), nil)
		ruleMatchingMock.EXPECT().
			Insert(gomock.Any(), gomock.Any())

		_, err := svc.Match(ctx, service.MatchRequest{Path: "some-path"})
		require.EqualError(t, err, "No rule match: some-path")

		time.Sleep(time.Millisecond) // waiting to submit the metric
	})

	t.Run("WHEN match", func(t *testing.T) {
		urlServiceMock.EXPECT().
			Match("some-path").
			Return(int64(1), map[string]string{"hello": "world"})
		ruleMatchingMock.EXPECT().
			Insert(gomock.Any(), gomock.Any())

		resp, err := svc.Match(ctx, service.MatchRequest{Path: "some-path"})
		require.NoError(t, err)
		require.Equal(t, &service.MatchResponse{
			RuleID:    1,
			PathParam: map[string]string{"hello": "world"},
		}, resp)

		time.Sleep(time.Millisecond) // waiting to submit the metri	})
}
