package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
)

type (
	onSettingRepo func(sqlmock.Sqlmock)

	settingFind struct {
		testName      string
		onSettingRepo onSettingRepo
		opts          []dbkit.SelectOption
		expected      []*repository.Setting
		expectedErr   string
	}

	settingUpdate struct {
		testName      string
		onSettingRepo onSettingRepo
		setting       *repository.Setting
		opt           dbkit.UpdateOption
		expectedErr   string
	}
)

func createSettingRepo(fn onSettingRepo) (repository.SettingRepo, *sql.DB) {
	db, mock, _ := sqlmock.New()
	if fn != nil {
		fn(mock)
	}
	return &repository.SettingRepoImpl{DB: db}, db
}

func TestSetting_Find(t *testing.T) {
	now := time.Now()
	testcases := []settingFind{
		{
			onSettingRepo: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT key, value, updated_at FROM settings").
					WillReturnError(errors.New("some-error"))
			},
			expectedErr: "some-error",
		},
		{
			onSettingRepo: func(mock sqlmock.Sqlmock) {

				mock.ExpectQuery("SELECT key, value, updated_at FROM settings").
					WillReturnRows(sqlmock.
						NewRows([]string{"key", "value", "updated_at"}).
						AddRow("key-1", "value-1", now).
						AddRow("key-2", "value-2", now),
					)
			},
			expected: []*repository.Setting{
				{Key: "key-1", Value: "value-1", UpdatedAt: now},
				{Key: "key-2", Value: "value-2", UpdatedAt: now},
			},
		},
		{
			testName: "with select option",
			opts: []dbkit.SelectOption{
				dbkit.Equal("key", "some-key"),
				dbkit.Equal("value", "some-value"),
			},
			onSettingRepo: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT key, value, updated_at FROM settings WHERE key = $1 AND value = $2`)).
					WithArgs("some-key", "some-value").
					WillReturnRows(sqlmock.
						NewRows([]string{"key", "value", "updated_at"}).
						AddRow("some-key", "some-value", now),
					)
			},
			expected: []*repository.Setting{
				{Key: "some-key", Value: "some-value", UpdatedAt: now},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			repo, db := createSettingRepo(tt.onSettingRepo)
			defer db.Close()

			settings, err := repo.Find(context.Background(), tt.opts...)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expected, settings)
		})
	}
}

func TestSettingUpdate(t *testing.T) {
	testcases := []settingUpdate{
		{
			setting: &repository.Setting{
				Value: "new-value",
			},
			opt: dbkit.Equal(repository.SettingCols.Key, "some-key"),
			onSettingRepo: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE settings SET value = $1, updated_at = $2 WHERE key = $3`)).
					WithArgs("new-value", sqlmock.AnyArg(), "some-key").
					WillReturnError(errors.New("some-error"))
			},
			expectedErr: "some-error",
		},
		{
			setting: &repository.Setting{
				Value: "new-value",
			},
			opt: dbkit.Equal(repository.SettingCols.Key, "some-key"),
			onSettingRepo: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE settings SET value = $1, updated_at = $2 WHERE key = $3`)).
					WithArgs("new-value", sqlmock.AnyArg(), "some-key").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			repo, db := createSettingRepo(tt.onSettingRepo)
			defer db.Close()

			err := repo.Update(context.Background(), tt.setting, tt.opt)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				return
			}
			require.NoError(t, err)
		})
	}

}
