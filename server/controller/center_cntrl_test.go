package controller_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/hotstone-seo/hotstone-seo/server/repository"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

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
		svc.EXPECT().AddMetaTag(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddMetaTag, "/", `{"name":"some-name", "content":"some-content"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddMetaTag(gomock.Any(), gomock.Any()).Return(&repository.Tag{Attributes: map[string]string{}}, nil)
		rr, err := echotest.DoPOST(cntrl.AddMetaTag, "/", `{"name":"some-name", "content":"some-content"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"id\":0,\"rule_id\":0,\"locale\":\"\",\"type\":\"\",\"attributes\":{},\"value\":\"\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
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
		svc.EXPECT().AddTitleTag(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddTitleTag, "/", `{"title":"some-name"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddTitleTag(gomock.Any(), gomock.Any()).Return(&repository.Tag{Attributes: map[string]string{}}, nil)
		rr, err := echotest.DoPOST(cntrl.AddTitleTag, "/", `{"title":"some-name"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"id\":0,\"rule_id\":0,\"locale\":\"\",\"type\":\"\",\"attributes\":{},\"value\":\"\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestCenterCntrl_AddCanonicalTag(t *testing.T) {
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
		svc.EXPECT().AddCanonicalTag(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddCanonicalTag, "/", `{"canonical":"test","rule_id":1,"href":"http://localhost"}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddCanonicalTag(gomock.Any(), gomock.Any()).Return(&repository.Tag{Attributes: map[string]string{}}, nil)
		rr, err := echotest.DoPOST(cntrl.AddCanonicalTag, "/", `{"canonical":"test","rule_id":1,"href":"http://localhost"}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"id\":0,\"rule_id\":0,\"locale\":\"\",\"type\":\"\",\"attributes\":{},\"value\":\"\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
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
		svc.EXPECT().AddScriptTag(gomock.Any(), gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echotest.DoPOST(cntrl.AddScriptTag, "/", `{"type":"javascript","rule_ud":1,"datasource_id":1}`, nil)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddScriptTag(gomock.Any(), gomock.Any()).Return(&repository.Tag{Attributes: map[string]string{}}, nil)
		rr, err := echotest.DoPOST(cntrl.AddScriptTag, "/", `{"type":"javascript","rule_ud":1,"datasource_id":1}`, nil)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"id\":0,\"rule_id\":0,\"locale\":\"\",\"type\":\"\",\"attributes\":{},\"value\":\"\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestCenterCntrl_AddArticle(t *testing.T) {
	centerCntrl := controller.CenterCntrl{}
	_, err := echotest.DoPOST(centerCntrl.AddArticle, "/", `invalid`, nil)
	require.EqualError(t, err, "code=501, message=Not implemented")
}
