package app_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-server/app"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typobj"
)

func TestModule(t *testing.T) {
	a := app.Module()
	require.True(t, typobj.IsActionable(a))
	require.True(t, typobj.IsConfigurer(a))
}
