package dbtxn_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
)

func TestRetrieveContext(t *testing.T) {
	testcases := []struct {
		testName        string
		ctx             context.Context
		expectedContext *dbtxn.Context
		expectedErr     string
	}{
		{
			ctx:         nil,
			expectedErr: "context.Context is nil",
		},
		{
			ctx:             context.Background(),
			expectedContext: nil,
		},
		{
			ctx:         context.WithValue(context.Background(), dbtxn.ContextKey, "meh"),
			expectedErr: "bad txn context",
		},
		{
			ctx:             context.WithValue(context.Background(), dbtxn.ContextKey, &dbtxn.Context{}),
			expectedContext: &dbtxn.Context{},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			context, err := dbtxn.RetrieveContext(tt.ctx)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedContext, context)
			}
		})
	}
}

func TestUse_When_Error(t *testing.T) {
	t.Run("Rollback when error", func(t *testing.T) {
		ctx := context.Background()
		defer dbtxn.Begin(&ctx)()

		db, mock, _ := sqlmock.New()
		mock.ExpectBegin()
		mock.ExpectRollback()

		txn, err := dbtxn.Use(ctx, db)
		require.NoError(t, err)
		require.True(t, txn.Txn())
		txn.SetError(errors.New("some-message"))
	})

	t.Run("Commit when no error", func(t *testing.T) {
		ctx := context.Background()
		defer dbtxn.Begin(&ctx)()

		db, mock, _ := sqlmock.New()
		mock.ExpectBegin()
		mock.ExpectCommit()

		txn, err := dbtxn.Use(ctx, db)
		require.NoError(t, err)
		require.True(t, txn.Txn())
	})
}
