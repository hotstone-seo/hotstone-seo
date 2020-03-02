package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestRuleController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ruleSvcMock := mock_service.NewMockRuleService(ctrl)
	ruleCntrl := controller.RuleCntrl{
		RuleService: ruleSvcMock,
	}
	t.Run("WHEN invalid rule request", func(t *testing.T) {
		_, err := echotest.DoPOST(ruleCntrl.Create, "/", `{ "name": "", "url_pattern": ""}`)
		require.EqualError(t, err, "code=400, message=Key: 'Rule.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Rule.UrlPattern' Error:Field validation for 'UrlPattern' failed on the 'required' tag")
	})
	t.Run("WHEN invalid json format", func(t *testing.T) {
		_, err := echotest.DoPOST(ruleCntrl.Create, "/", `invalid`)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN insert error", func(t *testing.T) {
		ruleSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("some-insert-error"))
		_, err := echotest.DoPOST(ruleCntrl.Create, "/", `{ "name": "some-name", "url_pattern": "some-pattern", "data_source_id":1}`)
		require.EqualError(t, err, "code=422, message=some-insert-error")
	})
	t.Run("WHEN insert success", func(t *testing.T) {
		ruleSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(999), nil)
		rr, err := echotest.DoPOST(ruleCntrl.Create, "/", `{ "name": "some-name", "url_pattern": "some-pattern", "data_source_id":1}`)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":999,\"name\":\"some-name\",\"url_pattern\":\"some-pattern\",\"data_source_id\":1,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}
