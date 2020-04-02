package typical_test

import (
	"context"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/typical"
	"github.com/stretchr/testify/require"
)

func TestCommandExist(t *testing.T) {
	testcases := []struct {
		name     string
		expected bool
	}{

		{"cd", true},
		{"bash", true},
		{"", false},
		{"invalid-command", false},
	}

	ctx := context.Background()
	for _, tt := range testcases {
		require.Equal(t, tt.expected, typical.AvailableCommand(ctx, tt.name))
	}

}
