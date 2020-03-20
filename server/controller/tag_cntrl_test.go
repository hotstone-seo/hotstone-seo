package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"

	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
)

func TestTagController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagSvcMock := mock_service.NewMockTagService(ctrl)
	tagCntrl := controller.TagCntrl{
		TagService: tagSvcMock,
	}
	t.Run("WHEN body is malformed", func(t *testing.T) {
		_, err := echotest.DoPOST(tagCntrl.Create, "/", `{ "rule_id", "type": "title" }`)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=12, error=invalid character ',' after object key")
	})
	t.Run("WHEN tag is invalid", func(t *testing.T) {
		_, err := echotest.DoPOST(tagCntrl.Create, "/", `{ "rule_id": 999, "locale": "en_US", "type": "title" }`)
		require.EqualError(t, err, "code=400, message=Key: 'Tag.Value' Error:Field validation for 'Value' failed on the 'noempty' tag")
	})
	t.Run("WHEN received error after inserting", func(t *testing.T) {
		tagSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("insert error"))
		_, err := echotest.DoPOST(tagCntrl.Create, "/", `{ "rule_id": 999, "locale": "en_US", "type": "title", "value": "Page Title" }`)
		require.EqualError(t, err, "code=422, message=insert error")
	})
	t.Run("When successfully insert new tag", func(t *testing.T) {
		tagSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(999), nil)
		rr, err := echotest.DoPOST(tagCntrl.Create, "/", `{ "rule_id": 999, "locale": "en_US", "type": "title", "value": "Page Title" }`)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":999,\"rule_id\":999,\"locale\":\"en_US\",\"type\":\"title\",\"attributes\":null,\"value\":\"Page Title\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}
