package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typmodule"

	"github.com/hotstone-seo/hotstone-server/app"
)

func TestModule(t *testing.T) {
	a := app.Module()
	require.True(t, typmodule.IsActionable(a))
}
