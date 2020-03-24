package server_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestModule(t *testing.T) {
	t.Run("SHOULD implement Configurer", func(t *testing.T) {
		var _ typcfg.Configurer = server.New()
	})
	t.Run("SHOULD implement Provider", func(t *testing.T) {
		var _ typapp.Provider = server.New()
	})
	t.Run("SHOULD implement Destroyer", func(t *testing.T) {
		var _ typapp.Destroyer = server.New()
	})
	t.Run("SHOULD implement EntryPointer", func(t *testing.T) {
		var _ typapp.EntryPointer = server.New()
	})
}
