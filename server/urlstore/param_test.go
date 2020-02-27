package urlstore_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/server/urlstore"
	"github.com/stretchr/testify/require"
)

func TestParam(t *testing.T) {
	testcases := []struct {
		s        string
		expected *urlstore.Param
	}{
		{
			s: "/<name1>",
			expected: &urlstore.Param{
				Start:        1,
				End:          7,
				StringBefore: "/",
				Raw:          "<name1>",
				Name:         "name1",
				Pattern:      "",
				AtLastPos:    true,
			},
		},
		{
			s: "<name1>",
			expected: &urlstore.Param{
				Start:     0,
				End:       6,
				Raw:       "<name1>",
				Name:      "name1",
				AtLastPos: true,
			},
		},
		{
			s: "<name1:pattern1>",
			expected: &urlstore.Param{
				Start:     0,
				End:       15,
				Raw:       "<name1:pattern1>",
				Name:      "name1",
				Pattern:   "pattern1",
				AtLastPos: true,
			},
		},
		{
			s: "/<name1>/some-string",
			expected: &urlstore.Param{
				Start:        1,
				End:          7,
				StringBefore: "/",
				StringAfter:  "/some-string",
				Raw:          "<name1>",
				Name:         "name1",
			},
		},
		{
			s: "<:.*>",
			expected: &urlstore.Param{
				Start:     0,
				End:       4,
				Raw:       "<:.*>",
				Pattern:   ".*",
				AtLastPos: true,
			},
		},
		{
			s: "no-param",
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, urlstore.CreateParam(tt.s))
	}
}
