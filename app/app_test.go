package app_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcli"

	"github.com/stretchr/testify/require"

	"github.com/hotstone-seo/hotstone-server/app"
)

func TestModule(t *testing.T) {
	a := app.Module()
	require.True(t, typcli.IsAppCommander(a))
	require.True(t, typcfg.IsConfigurer(a))
}
