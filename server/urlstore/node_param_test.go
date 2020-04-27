package urlstore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindFirstParam(t *testing.T) {
	testcases := []struct {
		s        string
		expected *nodeParam
	}{
		{
			s: "/<name1>",
			expected: &nodeParam{
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
			expected: &nodeParam{
				Start:     0,
				End:       6,
				Raw:       "<name1>",
				Name:      "name1",
				AtLastPos: true,
			},
		},
		{
			s: "<name1:pattern1>",
			expected: &nodeParam{
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
			expected: &nodeParam{
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
			expected: &nodeParam{
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
		require.Equal(t, tt.expected, findFirstParam(tt.s))
	}
}
