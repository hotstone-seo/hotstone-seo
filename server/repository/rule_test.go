package repository_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/stretchr/testify/require"
)

func TestRuleValidation(t *testing.T) {
	t.Run("WHEN rule is valid", func(t *testing.T) {
		rule := repository.Rule{Name: "My Rule", URLPattern: "/some/url"}
		require.NoError(t, rule.Validate())
	})
	t.Run("WHEN name is missing", func(t *testing.T) {
		rule := repository.Rule{URLPattern: "/some/url"}
		require.EqualError(t, rule.Validate(),
			"Key: 'Rule.Name' Error:Field validation for 'Name' failed on the 'required' tag")
	})
	t.Run("WHEN UrlPattern is missing", func(t *testing.T) {
		rule := repository.Rule{Name: "My Rule"}
		require.EqualError(t, rule.Validate(),
			"Key: 'Rule.URLPattern' Error:Field validation for 'URLPattern' failed on the 'required' tag")
	})
}
