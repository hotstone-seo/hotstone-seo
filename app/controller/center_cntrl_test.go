package controller_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-server/app/controller"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echokit"
)

func TestCenterCntrl_AddMetaTag(t *testing.T) {
	centerCntrl := controller.CenterCntrl{}
	_, err := echokit.DoPOST(centerCntrl.AddMetaTag, "/", `invalid`)
	require.EqualError(t, err, "code=501, message=Not implemented")
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
