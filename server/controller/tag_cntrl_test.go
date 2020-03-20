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
	"github.com/hotstone-seo/hotstone-seo/server/repository"
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
	t.Run("WHEN successfully insert new tag", func(t *testing.T) {
		tagSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(999), nil)
		rr, err := echotest.DoPOST(tagCntrl.Create, "/", `{ "rule_id": 999, "locale": "en_US", "type": "title", "value": "Page Title" }`)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":999,\"rule_id\":999,\"locale\":\"en_US\",\"type\":\"title\",\"attributes\":null,\"value\":\"Page Title\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestTagController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagSvcMock := mock_service.NewMockTagService(ctrl)
	tagCntrl := controller.TagCntrl{
		TagService: tagSvcMock,
	}
	t.Run("WHEN rule_id is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(tagCntrl.Find, "/tags?rule_id=invalid", nil)
		require.EqualError(t, err, "code=400, message=Invalid Rule ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		tagSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(tagCntrl.Find, "/tags?rule_id=999&locale=en_US", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		tagSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(
			[]*repository.Tag{
				&repository.Tag{
					ID:         999,
					RuleID:     999,
					Locale:     "en_US",
					Type:       "title",
					Value:      "Page Title",
					Attributes: []byte("{}"),
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(tagCntrl.Find, "/tags?rule_id=999&locale=en_US", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":999,\"rule_id\":999,\"locale\":\"en_US\",\"type\":\"title\",\"attributes\":{},\"value\":\"Page Title\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}]\n", rr.Body.String())
	})
}

func TestTagController_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagSvcMock := mock_service.NewMockTagService(ctrl)
	tagCntrl := controller.TagCntrl{
		TagService: tagSvcMock,
	}
	t.Run("WHEN ID param is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		tagSvcMock.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(nil, errors.New("find error"))
		_, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN tag not found", func(t *testing.T) {
		tagSvcMock.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(nil, nil)
		_, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=404, message=Tag#999 not found")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		tagSvcMock.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(
			&repository.Tag{
				ID:         999,
				RuleID:     999,
				Locale:     "en_US",
				Type:       "title",
				Value:      "Page Title",
				Attributes: []byte("{}"),
			},
			nil,
		)
		rr, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"id\":999,\"rule_id\":999,\"locale\":\"en_US\",\"type\":\"title\",\"attributes\":{},\"value\":\"Page Title\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}
