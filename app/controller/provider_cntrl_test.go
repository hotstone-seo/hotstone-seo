package controller_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/mock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/app/controller"
	"github.com/hotstone-seo/hotstone-seo/app/service"
)

func TestProviderCntrl_MatchRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockProviderService(ctrl)
	cntrl := controller.ProviderCntrl{
		ProviderService: svc,
	}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.MatchRule, "/", `{invalid`)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=2, error=invalid character 'i' looking for beginning of object key string")
	})
	t.Run("WHEN error match rule", func(t *testing.T) {
		svc.EXPECT().MatchRule(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().MatchRule(gomock.Any(), gomock.Any()).Return(&service.MatchRuleResponse{RuleID: 12345, PathParam: map[string]string{"param01": "value01"}}, nil)
		rec, err := echotest.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`)
		require.NoError(t, err)
		require.Equal(t, 200, rec.Code)
		require.Equal(t, "{\"rule_id\":12345,\"path_param\":{\"param01\":\"value01\"}}\n", rec.Body.String())
	})
}

func TestProviderCntrl_RetrieveData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockProviderService(ctrl)
	cntrl := controller.ProviderCntrl{ProviderService: svc}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.RetrieveData, "/", `{invalid`)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=2, error=invalid character 'i' looking for beginning of object key string")
	})
	t.Run("WHEN error match rule", func(t *testing.T) {
		svc.EXPECT().RetrieveData(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.RetrieveData, "/", `{"rule_id": 99999}`)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().RetrieveData(gomock.Any(), gomock.Any()).
			Return(&http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader("some-string"))}, nil)
		rec, err := echotest.DoPOST(cntrl.RetrieveData, "/", `{"rule_id": 99999}`)
		require.NoError(t, err)
		require.Equal(t, 200, rec.Code)
		require.Equal(t, "some-string", rec.Body.String())
	})
}
