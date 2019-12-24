package task

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/app/urlstore"
	"github.com/hotstone-seo/hotstone-server/mock"
	"github.com/stretchr/testify/require"
)

func TestURLStoreSchedulerImpl_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	list1And2URLStoreSync := []*repository.URLStoreSync{
		&repository.URLStoreSync{Version: 1, Operation: "INSERT", RuleID: 1, LatestURLPattern: "/url/1"},
		&repository.URLStoreSync{Version: 2, Operation: "UPDATE", RuleID: 1, LatestURLPattern: "/url/1update"},
	}

	list3And4URLStoreSync := []*repository.URLStoreSync{
		&repository.URLStoreSync{Version: 3, Operation: "INSERT", RuleID: 2, LatestURLPattern: "/url/b"},
		&repository.URLStoreSync{Version: 4, Operation: "UPDATE", RuleID: 2, LatestURLPattern: "/url/bupdate"},
	}

	urlStoreSyncSvcMock := mock.NewMockURLStoreSyncService(ctrl)

	urlStoreScheduler := &URLStoreSchedulerImpl{
		URLStoreSyncService: urlStoreSyncSvcMock,
		urlStore:            urlstore.InitURLStore(),
		latestVersion:       -1,
	}

	t.Run("WHEN first sync (s.latestVersion < latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, -1, urlStoreScheduler.latestVersion)
		require.Equal(t, 0, urlStoreScheduler.urlStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(len(list1And2URLStoreSync)), nil)
		urlStoreSyncSvcMock.EXPECT().GetListDiff(ctx, gomock.Eq(int64(-1))).Return(list1And2URLStoreSync, nil)

		err := urlStoreScheduler.Sync()
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreScheduler.latestVersion)
		require.Equal(t, 1, urlStoreScheduler.urlStore.Count())
	})

	t.Run("WHEN second sync (s.latestVersion == latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreScheduler.latestVersion)
		require.Equal(t, 1, urlStoreScheduler.urlStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(2), nil)

		err := urlStoreScheduler.Sync()
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreScheduler.latestVersion)
		require.Equal(t, 1, urlStoreScheduler.urlStore.Count())
	})

	t.Run("WHEN third sync (s.latestVersion < latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreScheduler.latestVersion)
		require.Equal(t, 1, urlStoreScheduler.urlStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(4), nil)
		urlStoreSyncSvcMock.EXPECT().GetListDiff(ctx, gomock.Eq(int64(2))).Return(list3And4URLStoreSync, nil)

		err := urlStoreScheduler.Sync()
		require.NoError(t, err)

		require.Equal(t, 4, urlStoreScheduler.latestVersion)
		require.Equal(t, 2, urlStoreScheduler.urlStore.Count())
	})

	t.Run("WHEN outlier case (s.latestVersion > latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 4, urlStoreScheduler.latestVersion)
		require.Equal(t, 2, urlStoreScheduler.urlStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(2), nil) // latestVersion from DB = 2 (somehow some rows has been deleted)
		urlStoreSyncSvcMock.EXPECT().List(ctx).Return(list1And2URLStoreSync, nil)

		err := urlStoreScheduler.Sync()
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreScheduler.latestVersion)
		require.Equal(t, 1, urlStoreScheduler.urlStore.Count())
	})

}
