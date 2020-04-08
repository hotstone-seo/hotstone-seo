package cachekit_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/stretchr/testify/require"
)

func TestNotModifiedError(t *testing.T) {
	testcases := []struct {
		desc     string
		err      error
		expected bool
	}{
		{
			desc:     "predefined no-modified error",
			err:      cachekit.ErrNoModified,
			expected: true,
		},
		{
			desc:     "predefined no-modified error with prefix",
			err:      fmt.Errorf("Prefix: %w", cachekit.ErrNoModified),
			expected: true,
		},
		{
			desc:     "random error",
			err:      errors.New("random-error"),
			expected: false,
		},
		{
			desc:     "nil error",
			err:      nil,
			expected: false,
		},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.expected, cachekit.NoModifiedError(tt.err), tt.desc)
	}
}
