package cachekit_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
)

func TestCreateCacheControl(t *testing.T) {
	testcases := []struct {
		req      *http.Request
		expected *cachekit.CacheControl
	}{
		{
			req:      &http.Request{},
			expected: cachekit.NewCacheControl(),
		},
		{
			req: func() *http.Request {
				req, _ := http.NewRequest("", "", nil)
				req.Header.Add("Cache-Control", "a, b, c")
				return req
			}(),
			expected: cachekit.NewCacheControl("a", "b", "c"),
		},
	}

	for _, tt := range testcases {
		require.EqualValues(t, tt.expected, cachekit.CreateCacheControl(tt.req))
	}
}

func TestCacheContro_NoCache(t *testing.T) {
	testcases := []struct {
		*cachekit.CacheControl
		expected bool
	}{
		{
			CacheControl: cachekit.NewCacheControl(),
		},
		{
			CacheControl: cachekit.NewCacheControl("no-cache"),
			expected:     true,
		},
		{
			CacheControl: cachekit.NewCacheControl("No-Cache"),
			expected:     true,
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.NoCache())
	}
}

func TestCacheContro_MaxAge(t *testing.T) {
	testcases := []struct {
		*cachekit.CacheControl
		expected int
	}{
		{
			CacheControl: cachekit.NewCacheControl(),
			expected:     cachekit.DefaultMaxAge,
		},
		{
			CacheControl: cachekit.NewCacheControl().WithDefaultMaxAge(1),
			expected:     1,
		},
		{
			CacheControl: cachekit.NewCacheControl("max-age=100").WithDefaultMaxAge(1),
			expected:     100,
		},
		{
			CacheControl: cachekit.NewCacheControl("max-age=invalid"),
			expected:     cachekit.DefaultMaxAge,
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.MaxAge())
	}
}
