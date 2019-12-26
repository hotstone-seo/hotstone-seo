package controller_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-server/app/controller"
	"github.com/hotstone-seo/hotstone-server/mock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echokit"
)

func TestCenterCntrl_AddMetaTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockCenterService(ctrl)
	cntrl := controller.CenterCntrl{
		CenterService: svc,
	}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echokit.DoPOST(cntrl.AddMetaTag, "/", `invalid`)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddMetaTag(gomock.Any()).Return(int64(0), errors.New("some-error"))
		_, err := echokit.DoPOST(cntrl.AddMetaTag, "/", `{"name":"some-name", "content":"some-content"}`)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().AddMetaTag(gomock.Any()).Return(int64(999), nil)
		rr, err := echokit.DoPOST(cntrl.AddMetaTag, "/", `{"name":"some-name", "content":"some-content"}`)
		require.NoError(t, err)
		require.Equal(t, 201, rr.Code)
		require.Equal(t, "{\"message\":\"Success insert new meta tag #999\"}\n", rr.Body.String())
	})

}

func TestCenterCntrl_AddTitleTag(t *testing.T) {
	centerCntrl := controller.CenterCntrl{}
	_, err := echokit.DoPOST(centerCntrl.AddTitleTag, "/", `invalid`)
	require.EqualError(t, err, "code=501, message=Not implemented")
}

func TestCenterCntrl_AddCanoncicalTag(t *testing.T) {
	centerCntrl := controller.CenterCntrl{}
	_, err := echokit.DoPOST(centerCntrl.AddCanoncicalTag, "/", `invalid`)
	require.EqualError(t, err, "code=501, message=Not implemented")
}

func TestCenterCntrl_AddScriptTag(t *testing.T) {
	centerCntrl := controller.CenterCntrl{}
	_, err := echokit.DoPOST(centerCntrl.AddScriptTag, "/", `invalid`)
	require.EqualError(t, err, "code=501, message=Not implemented")
}

func TestCenterCntrl_AddArticle(t *testing.T) {
	centerCntrl := controller.CenterCntrl{}
	_, err := echokit.DoPOST(centerCntrl.AddArticle, "/", `invalid`)
	require.EqualError(t, err, "code=501, message=Not implemented")
}
