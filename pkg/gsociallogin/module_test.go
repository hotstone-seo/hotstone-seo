package gsociallogin_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/pkg/gsociallogin"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestModule(t *testing.T) {
	t.Run("SHOULD implement configurer", func(t *testing.T) {
		var _ typcfg.Configurer = gsociallogin.New()
	})
	t.Run("SHOULD implement provider", func(t *testing.T) {
		var _ typapp.Provider = gsociallogin.New()
	})

}
