package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestCenterCntrl_AddTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_service.NewMockCenterService(ctrl)
	cntrl := controller.CenterCntrl{
		CenterService: svc,
	}
	t.Run("WHEN type is not recognized", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.AddTag, "/", "{}", map[string]string{"type": "invalid"})
		require.EqualError(t, err, "code=400, message=invalid is not a valid type")
	})
	t.Run("WHEN body is malformed", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.AddTag, "/", `{ "title" }`, map[string]string{"type": "title"})
		require.EqualError(t, err, "code=400, message=code=400, message=Syntax error: offset=11, error=invalid character '}' after object key")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		svc.EXPECT().AddTag(gomock.Any(), &service.AddTitleTagRequest{Title: "Page Title"}).Return(nil, errors.New("insert error"))
		_, err := echotest.DoPOST(cntrl.AddTag, "/", `{ "title": "Page Title" }`, map[string]string{"type": "title"})
		require.EqualError(t, err, "code=422, message=insert error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		svc.EXPECT().AddTag(gomock.Any(), &service.AddTitleTagRequest{Title: "Page Title"}).Return(
			&repository.Tag{Value: "Page Title", Attributes: []byte("{}")},
			nil,
		)
		rr, err := echotest.DoPOST(cntrl.AddTag, "/", `{ "title": "Page Title" }`, map[string]string{"type": "title"})
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":0,\"rule_id\":0,\"locale\":\"\",\"type\":\"\",\"attributes\":{},\"value\":\"Page Title\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestCenterCntrl_AddMetaTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_service.NewMockCenterService(ctrl)
	cntrl := controller.CenterCntrl{
		CenterService: svc,
	}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.AddMetaTag, "/", `invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddMetaTag(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddMetaTag, "/", `{"name":"some-name", "content":"some-content"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddMetaTag(gomock.Any(), gomock.Any()).Return(int64(999), nil)
		rr, err := echotest.DoPOST(cntrl.AddMetaTag, "/", `{"name":"some-name", "content":"some-content"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"message\":\"Success insert new meta tag #999\"}\n", rr.Body.String())
	})

}

func TestCenterCntrl_AddTitleTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_service.NewMockCenterService(ctrl)
	cntrl := controller.CenterCntrl{
		CenterService: svc,
	}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.AddTitleTag, "/", `invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddTitleTag(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddTitleTag, "/", `{"title":"some-name"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddTitleTag(gomock.Any(), gomock.Any()).Return(int64(100), nil)
		rr, err := echotest.DoPOST(cntrl.AddTitleTag, "/", `{"title":"some-name"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"message\":\"Success insert new title tag #100\"}\n", rr.Body.String())
	})
}

func TestCenterCntrl_AddCanoncicalTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_service.NewMockCenterService(ctrl)
	cntrl := controller.CenterCntrl{
		CenterService: svc,
	}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.AddCanonicalTag, "/", `invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddCanonicalTag(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddCanonicalTag, "/", `{"canonical":"test","rule_id":1,"href":"http://localhost"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddCanonicalTag(gomock.Any(), gomock.Any()).Return(int64(101), nil)
		rr, err := echotest.DoPOST(cntrl.AddCanonicalTag, "/", `{"canonical":"test","rule_id":1,"href":"http://localhost"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"message\":\"Success insert new canonical tag #101\"}\n", rr.Body.String())
	})
}

func TestCenterCntrl_AddScriptTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_service.NewMockCenterService(ctrl)
	cntrl := controller.CenterCntrl{
		CenterService: svc,
	}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echotest.DoPOST(cntrl.AddScriptTag, "/", `invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddScriptTag(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddScriptTag, "/", `{"type":"javascript","rule_ud":1,"datasource_id":1}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddScriptTag(gomock.Any(), gomock.Any()).Return(int64(102), nil)
		rr, err := echotest.DoPOST(cntrl.AddScriptTag, "/", `{"type":"javascript","rule_ud":1,"datasource_id":1}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"message\":\"Success insert new canonical tag #102\"}\n", rr.Body.String())
	})
}

func TestCenterCntrl_AddArticle(t *testing.T) {
	centerCntrl := controller.CenterCntrl{}
	_, err := echotest.DoPOST(centerCntrl.AddArticle, "/", `invalid`, nil)
	require.EqualError(t, err, "code=501, message=Not implemented")
}
