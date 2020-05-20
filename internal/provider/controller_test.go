package provider_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/internal/provider"
	"github.com/hotstone-seo/hotstone-seo/internal/provider_mock"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
)

func TestController_Match(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := provider_mock.NewMockService(ctrl)
	cntrl := provider.Controller{
		Service: svc,
	}

	t.Run("WHEN error match rule", func(t *testing.T) {
		svc.EXPECT().Match(gomock.Any(), gomock.Any()).Return(nil, errvalid.New("some-error"))
		_, err := echotest.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})

	t.Run("WHEN error match rule", func(t *testing.T) {
		svc.EXPECT().Match(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`, nil)
		require.EqualError(t, err, "code=500, message=some-error")
	})

	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().Match(gomock.Any(), gomock.Any()).Return(&provider.MatchResponse{RuleID: 12345, PathParam: map[string]string{"param01": "value01"}}, nil)
		rec, err := echotest.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 200, rec.Code)
		require.Equal(t, "{\"rule_id\":12345,\"path_param\":{\"param01\":\"value01\"}}\n", rec.Body.String())
	})
}

func TestProviderController_FetchTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := provider_mock.NewMockService(ctrl)
	cntrl := provider.Controller{
		Service: mockService,
	}

	t.Run("", func(t *testing.T) {
		mockService.EXPECT().FetchTagsWithCache(gomock.Any(), gomock.Any(), gomock.Any())
		_, err := echotest.DoGET(cntrl.FetchTag, "/?locale=id_ID", map[string]string{
			"id": "1",
		})
		require.NoError(t, err)
	})

	t.Run("GIVEN validation error", func(t *testing.T) {
		mockService.EXPECT().
			FetchTagsWithCache(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errvalid.New("some error"))

		_, err := echotest.DoGET(cntrl.FetchTag, "/", nil)
		require.EqualError(t, err, "code=422, message=some error")
	})

	t.Run("GIVEN cache is not modified", func(t *testing.T) {
		mockService.EXPECT().
			FetchTagsWithCache(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, cachekit.ErrNotModified)
		_, err := echotest.DoGET(cntrl.FetchTag, "/?locale=en_US", map[string]string{
			"id": "1",
		})
		require.EqualError(t, err, "code=304, message=Not Modified")
	})
}
