package dbtxn_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/stretchr/testify/require"
)

func TestErrCtx(t *testing.T) {
	ctx := dbtxn.WithContext(context.Background())
	dbtxn.SetError(ctx, errors.New("some-error"))
	require.EqualError(t, dbtxn.Error(ctx), "some-error")
}
