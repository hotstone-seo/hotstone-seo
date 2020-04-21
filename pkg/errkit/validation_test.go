package errkit_test

import (
	"errors"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/pkg/errkit"
	"github.com/stretchr/testify/require"
)

func TestValidationError(t *testing.T) {
	testcases := []struct {
		desc             string
		err              error
		expectedTrue     bool
		expectedMessages []string
	}{
		{
			desc:             "create validation error from messages",
			err:              errkit.ValidationErr("message-1", "message-2", "message-3"),
			expectedTrue:     true,
			expectedMessages: []string{"message-1", "message-2", "message-3"},
		},
		{
			desc:             "wrap error to validaton error ",
			err:              errkit.WrapValidation(errors.New("some-error")),
			expectedTrue:     true,
			expectedMessages: []string{"some-error"},
		},
		{
			desc:             "nil error",
			err:              nil,
			expectedTrue:     false,
			expectedMessages: nil,
		},
		{
			desc:             "not validation error",
			err:              errors.New("some-error"),
			expectedTrue:     false,
			expectedMessages: nil,
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expectedTrue, errkit.IsValidationErr(tt.err), tt.desc)
		require.Equal(t, tt.expectedMessages, errkit.ValidationMessages(tt.err), tt.desc)
	}

}
