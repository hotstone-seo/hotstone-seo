package repository_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/stretchr/testify/require"
)

func TestRuleValidation(t *testing.T) {
	t.Run("When rule is valid", func(t *testing.T) {
		rule := repository.Rule{Name: "My Rule", UrlPattern: "/some/url"}
		require.NoError(t, rule.Validate())
	})
	t.Run("When name is missing", func(t *testing.T) {
		rule := repository.Rule{UrlPattern: "/some/url"}
		require.EqualError(t, rule.Validate(),
			"Key: 'Rule.Name' Error:Field validation for 'Name' failed on the 'required' tag")
	})
	t.Run("When UrlPattern is missing", func(t *testing.T) {
		rule := repository.Rule{Name: "My Rule"}
		require.EqualError(t, rule.Validate(),
			"Key: 'Rule.UrlPattern' Error:Field validation for 'UrlPattern' failed on the 'required' tag")
	})
}
