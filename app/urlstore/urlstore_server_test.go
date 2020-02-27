package urlstore_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/app/mock_urlstore"
	"github.com/hotstone-seo/hotstone-seo/app/urlstore"
	"github.com/stretchr/testify/require"
	"github.com/xorcare/pointer"
)

func TestURLStoreServerImpl_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	list1And2URLStoreSync := []*urlstore.URLStoreSync{
		&urlstore.URLStoreSync{Version: 1, Operation: "INSERT", RuleID: 1, LatestURLPattern: pointer.String("/url/1")},
		&urlstore.URLStoreSync{Version: 2, Operation: "UPDATE", RuleID: 1, LatestURLPattern: pointer.String("/url/1update")},
	}

	list3And4URLStoreSync := []*urlstore.URLStoreSync{
		&urlstore.URLStoreSync{Version: 3, Operation: "INSERT", RuleID: 2, LatestURLPattern: pointer.String("/url/b")},
		&urlstore.URLStoreSync{Version: 4, Operation: "UPDATE", RuleID: 2, LatestURLPattern: pointer.String("/url/bupdate")},
	}

	urlStoreSyncSvcMock := mock_urlstore.NewMockURLStoreSyncService(ctrl)

	urlStoreServer := &urlstore.URLStoreServerImpl{
		URLStoreSyncService: urlStoreSyncSvcMock,
		URLStore:            urlstore.InitURLStore(),
		LatestVersion:       -1,
	}

	t.Run("WHEN first sync (s.LatestVersion < latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, -1, urlStoreServer.LatestVersion)
		require.Equal(t, 0, urlStoreServer.URLStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(len(list1And2URLStoreSync)), nil)
		urlStoreSyncSvcMock.EXPECT().GetListDiff(ctx, gomock.Eq(int64(-1))).Return(list1And2URLStoreSync, nil)

		err := urlStoreServer.Sync()
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.URLStore.Count())
	})

	t.Run("WHEN second sync (s.LatestVersion == latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.URLStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(2), nil)

		err := urlStoreServer.Sync()
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.URLStore.Count())
	})

	t.Run("WHEN third sync (s.LatestVersion < latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.URLStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(4), nil)
		urlStoreSyncSvcMock.EXPECT().GetListDiff(ctx, gomock.Eq(int64(2))).Return(list3And4URLStoreSync, nil)

		err := urlStoreServer.Sync()
		require.NoError(t, err)

		require.Equal(t, 4, urlStoreServer.LatestVersion)
		require.Equal(t, 2, urlStoreServer.URLStore.Count())
	})

	t.Run("WHEN outlier case (s.LatestVersion > latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 4, urlStoreServer.LatestVersion)
		require.Equal(t, 2, urlStoreServer.URLStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(2), nil) // latestVersion from DB = 2 (somehow some rows has been deleted)
		urlStoreSyncSvcMock.EXPECT().Find(ctx).Return(list1And2URLStoreSync, nil)

		err := urlStoreServer.Sync()
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.URLStore.Count())
	})

	t.Run("WHEN outlier case (no data in urlstore_sync)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.URLStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(0), nil) // latestVersion from DB = 0 (all data have been deleted)

		err := urlStoreServer.Sync()
		require.NoError(t, err)

		require.Equal(t, 0, urlStoreServer.LatestVersion)
		require.Equal(t, 0, urlStoreServer.URLStore.Count())
	})

}
