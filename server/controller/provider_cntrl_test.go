package controller_test

import (
	"errors"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"

	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/hotstone-seo/hotstone-seo/server/service"
)

func TestProviderCntrl_MatchRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_service.NewMockProviderService(ctrl)
	cntrl := controller.ProviderCntrl{
		ProviderService: svc,
	}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.MatchRule, "/", `{invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=2, error=invalid character 'i' looking for beginning of object key string")
	})
	t.Run("WHEN error match rule", func(t *testing.T) {
		svc.EXPECT().Match(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().Match(gomock.Any(), gomock.Any()).Return(&service.MatchResponse{RuleID: 12345, PathParam: map[string]string{"param01": "value01"}}, nil)
		rec, err := echotest.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 200, rec.Code)
		require.Equal(t, "{\"rule_id\":12345,\"path_param\":{\"param01\":\"value01\"}}\n", rec.Body.String())
	})
}

func TestProviderController_FetchTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockProviderService(ctrl)
	cntrl := controller.ProviderCntrl{
		ProviderService: mockService,
	}

	t.Run("", func(t *testing.T) {
		mockService.EXPECT().FetchTagsWithCache(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
		_, err := echotest.DoGET(cntrl.FetchTag, "/?locale=id-id", map[string]string{
			"id": "1",
		})
		require.NoError(t, err)
	})

	t.Run("GIVEN no url param id", func(t *testing.T) {
		_, err := echotest.DoGET(cntrl.FetchTag, "/", nil)
		require.EqualError(t, err, "code=422, message=Missing url param for `ID`")
	})

	t.Run("GIVEN non-integer url param id", func(t *testing.T) {
		_, err := echotest.DoGET(cntrl.FetchTag, "/", map[string]string{
			"id": "not-integer",
		})
		require.EqualError(t, err, "code=422, message=strconv.ParseInt: parsing \"not-integer\": invalid syntax")
	})

	t.Run("GIVEN no locale", func(t *testing.T) {
		_, err := echotest.DoGET(cntrl.FetchTag, "/", map[string]string{
			"id": "1",
		})
		require.EqualError(t, err, "code=422, message=Missing query param for `Locale`")
	})

	t.Run("GIVEN cache is not modified", func(t *testing.T) {
		mockService.EXPECT().
			FetchTagsWithCache(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, cachekit.ErrNotModified)
		_, err := echotest.DoGET(cntrl.FetchTag, "/?locale=en_US", map[string]string{
			"id": "1",
		})
		require.EqualError(t, err, "code=304, message=Not Modified")
	})
}
