package cachekit_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
)

func TestCreateCacheControl(t *testing.T) {
	testcases := []struct {
		req      *http.Request
		expected *cachekit.Pragma
	}{
		{
			req:      &http.Request{},
			expected: cachekit.NewPragma(),
		},
		{
			req: func() *http.Request {
				req, _ := http.NewRequest("", "", nil)
				req.Header.Add("Cache-Control", "a, b, c")
				return req
			}(),
			expected: cachekit.NewPragma("a", "b", "c"),
		},
	}

	for _, tt := range testcases {
		require.EqualValues(t, tt.expected, cachekit.CreatePragma(tt.req))
	}
}

func TestCacheContro_NoCache(t *testing.T) {
	testcases := []struct {
		*cachekit.Pragma
		expected bool
	}{
		{
			Pragma: cachekit.NewPragma(),
		},
		{
			Pragma:   cachekit.NewPragma("no-cache"),
			expected: true,
		},
		{
			Pragma:   cachekit.NewPragma("No-Cache"),
			expected: true,
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.NoCache())
	}
}

func TestCacheContro_MaxAge(t *testing.T) {
	testcases := []struct {
		*cachekit.Pragma
		expected time.Duration
	}{
		{
			Pragma:   cachekit.NewPragma(),
			expected: cachekit.DefaultMaxAge,
		},
		{
			Pragma:   cachekit.NewPragma().WithDefaultMaxAge(1),
			expected: 1,
		},
		{
			Pragma:   cachekit.NewPragma("max-age=100").WithDefaultMaxAge(1),
			expected: 100 * time.Second,
		},
		{
			Pragma:   cachekit.NewPragma("max-age=invalid"),
			expected: cachekit.DefaultMaxAge,
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.MaxAge())
	}
}
