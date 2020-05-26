package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/stretchr/testify/require"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository_mock"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
)

type (
	onSettingSvc func(*repository_mock.MockSettingRepo)

	settingFind struct {
		testName     string
		onSettingSvc onSettingSvc
		expected     []*repository.Setting
		expectedErr  string
	}

	settingFindOne struct {
		testName     string
		key          string
		onSettingSvc onSettingSvc
		expected     *repository.Setting
		expectedErr  string
	}

	settingUpdate struct {
		testName     string
		key          string
		onSettingSvc onSettingSvc
		setting      *repository.Setting
		expectedErr  string
	}
)

func createSettingSvc(t *testing.T, fn onSettingSvc) (*service.SettingSvcImpl, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	repo := repository_mock.NewMockSettingRepo(ctrl)
	if fn != nil {
		fn(repo)
	}

	return &service.SettingSvcImpl{
		SettingRepo: repo,
	}, ctrl

}

func TestSetting_Find(t *testing.T) {
	testcases := []settingFind{
		{
			onSettingSvc: func(repo *repository_mock.MockSettingRepo) {
				repo.EXPECT().Find(gomock.Any()).Return(nil, errors.New("some-error"))
			},
			expectedErr: "some-error",
		},
		{
			onSettingSvc: func(repo *repository_mock.MockSettingRepo) {
				repo.EXPECT().Find(gomock.Any()).Return([]*repository.Setting{
					{Key: "key-1", Value: "value-1"},
					{Key: "key-2", Value: "value-2"},
				}, nil)
			},
			expected: []*repository.Setting{
				{Key: "key-1", Value: "value-1"},
				{Key: "key-2", Value: "value-2"},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			svc, ctrl := createSettingSvc(t, tt.onSettingSvc)
			defer ctrl.Finish()

			settings, err := svc.Find(context.Background())
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, settings)
		})
	}
}

func TestSetting_FindOne(t *testing.T) {
	testcases := []settingFindOne{
		{
			testName:    "key is empty",
			key:         "",
			expectedErr: "Validation: Key is missing",
		},
		{
			key: "some-key",
			onSettingSvc: func(repo *repository_mock.MockSettingRepo) {
				repo.EXPECT().
					Find(gomock.Any(), dbkit.Equal("key", "some-key")).
					Return(nil, errors.New("some-error"))
			},
			expectedErr: "some-error",
		},
		{
			key: "some-key",
			onSettingSvc: func(repo *repository_mock.MockSettingRepo) {
				repo.EXPECT().
					Find(gomock.Any(), dbkit.Equal("key", "some-key")).
					Return([]*repository.Setting{
						{Key: "some-key", Value: "some-value"},
					}, nil)
			},
			expected: &repository.Setting{
				Key:   "some-key",
				Value: "some-value",
			},
		},
		{
			testName: "repo return empty list but no error",
			key:      "some-key",
			onSettingSvc: func(repo *repository_mock.MockSettingRepo) {
				repo.EXPECT().
					Find(gomock.Any(), dbkit.Equal("key", "some-key")).
					Return([]*repository.Setting{}, nil)
			},
			expectedErr: "sql: no rows in result set",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			svc, ctrl := createSettingSvc(t, tt.onSettingSvc)
			defer ctrl.Finish()

			settings, err := svc.FindOne(context.Background(), tt.key)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, settings)
		})
	}
}

func TestSetting_Update(t *testing.T) {
	testcases := []settingUpdate{
		{
			testName:    "key is missing",
			key:         "",
			expectedErr: "Validation: key is missing",
		},
		{
			testName:    "value is missing",
			key:         "some-key",
			setting:     &repository.Setting{},
			expectedErr: "Key: 'Setting.Value' Error:Field validation for 'Value' failed on the 'required' tag",
		},
		{
			key:     "some-key",
			setting: &repository.Setting{Key: "some-key", Value: "some-value"},
			onSettingSvc: func(repo *repository_mock.MockSettingRepo) {
				repo.EXPECT().
					Update(
						gomock.Any(),
						&repository.Setting{Key: "some-key", Value: "some-value"},
						dbkit.Equal("key", "some-key"),
					).
					Return(errors.New("some-error"))
			},
			expectedErr: "some-error",
		},
		{
			key:     "some-key",
			setting: &repository.Setting{Key: "some-key", Value: "some-value"},
			onSettingSvc: func(repo *repository_mock.MockSettingRepo) {
				repo.EXPECT().
					Update(
						gomock.Any(),
						&repository.Setting{Key: "some-key", Value: "some-value"},
						dbkit.Equal("key", "some-key"),
					).
					Return(nil)
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			svc, ctrl := createSettingSvc(t, tt.onSettingSvc)
			defer ctrl.Finish()

			err := svc.Update(context.Background(), tt.key, tt.setting)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				return
			}

			require.NoError(t, err)
		})
	}
}
