package oauth2google_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestModule(t *testing.T) {
	t.Run("SHOULD implement configurer", func(t *testing.T) {
		var _ typcfg.Configurer = oauth2google.New()
	})
	t.Run("SHOULD implement provider", func(t *testing.T) {
		var _ typapp.Provider = oauth2google.New()
	})

}
