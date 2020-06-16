package controller_test

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/internal/api/controller"
	"github.com/stretchr/testify/require"
)

func TestIsRoleAllow(t *testing.T) {
	testcases := []struct {
		testName string
		path     string
		claims   jwt.MapClaims
		expected bool
	}{
		{
			path: "/api/something",
			claims: map[string]interface{}{
				"paths": []interface{}{
					"/api/something",
				},
			},
			expected: true,
		},
		{
			path: "/api/something",
			claims: map[string]interface{}{
				"paths": []interface{}{
					"/api/something-else",
					"/api/something",
				},
			},
			expected: true,
		},
		{
			path: "/api/missing",
			claims: map[string]interface{}{
				"paths": []interface{}{
					"/api/something",
				},
			},
			expected: false,
		},
		{
			path: "/api/something",
			claims: map[string]interface{}{
				"paths": []interface{}{
					".*",
				},
			},
			expected: true,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, controller.IsRoleAllow(tt.path, tt.claims))
		})
	}
}
