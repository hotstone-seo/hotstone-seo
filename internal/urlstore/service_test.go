package urlstore_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore_mock"
	"github.com/stretchr/testify/require"
	"github.com/xorcare/pointer"
)

func TestURLStoreServerImpl_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	list1And2Sync := []*urlstore.Sync{
		&urlstore.Sync{Version: 1, Operation: "INSERT", RuleID: 1, LatestURLPattern: pointer.String("/url/1")},
		&urlstore.Sync{Version: 2, Operation: "UPDATE", RuleID: 1, LatestURLPattern: pointer.String("/url/1update")},
	}

	list3And4Sync := []*urlstore.Sync{
		&urlstore.Sync{Version: 3, Operation: "INSERT", RuleID: 2, LatestURLPattern: pointer.String("/url/b")},
		&urlstore.Sync{Version: 4, Operation: "UPDATE", RuleID: 2, LatestURLPattern: pointer.String("/url/bupdate")},
	}

	mockRepo := urlstore_mock.NewMockSyncRepo(ctrl)

	urlStoreServer := &urlstore.Service{
		SyncRepo:   mockRepo,
		Store:         urlstore.NewStore(),
		LatestVersion: -1,
	}

	t.Run("WHEN first sync (s.LatestVersion < latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, -1, urlStoreServer.LatestVersion)
		require.Equal(t, 0, urlStoreServer.Store.Count())

		ctx := context.Background()
		mockRepo.EXPECT().GetLatestVersion(ctx).Return(int64(len(list1And2Sync)), nil)
		mockRepo.EXPECT().GetListDiff(ctx, gomock.Eq(int64(-1))).Return(list1And2Sync, nil)

		err := urlStoreServer.Sync(ctx)
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.Store.Count())
	})

	t.Run("WHEN second sync (s.LatestVersion == latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.Store.Count())

		ctx := context.Background()
		mockRepo.EXPECT().GetLatestVersion(ctx).Return(int64(2), nil)

		err := urlStoreServer.Sync(ctx)
		require.NoError(t, err)

		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.Store.Count())
	})

	t.Run("WHEN third sync (s.LatestVersion < latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.Store.Count())

		ctx := context.Background()
		mockRepo.EXPECT().GetLatestVersion(ctx).Return(int64(4), nil)
		mockRepo.EXPECT().GetListDiff(ctx, gomock.Eq(int64(2))).Return(list3And4Sync, nil)

		err := urlStoreServer.Sync(ctx)
		require.NoError(t, err)

		require.Equal(t, 4, urlStoreServer.LatestVersion)
		require.Equal(t, 2, urlStoreServer.Store.Count())
	})

	t.Run("WHEN outlier case (s.LatestVersion > latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 4, urlStoreServer.LatestVersion)
		require.Equal(t, 2, urlStoreServer.Store.Count())

		ctx := context.Background()
		mockRepo.EXPECT().GetLatestVersion(ctx).Return(int64(2), nil) // latestVersion from DB = 2 (somehow some rows has been deleted)
		mockRepo.EXPECT().Find(ctx).Return(list1And2Sync, nil)

		err := urlStoreServer.Sync(ctx)
		require.NoError(t, err)

		fmt.Println(urlStoreServer.Store.String())

		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.Store.Count())
	})

	t.Run("WHEN outlier case (no data in url_sync)", func(t *testing.T) {
		require.Equal(t, 2, urlStoreServer.LatestVersion)
		require.Equal(t, 1, urlStoreServer.Store.Count())

		ctx := context.Background()
		mockRepo.EXPECT().GetLatestVersion(ctx).Return(int64(0), nil) // latestVersion from DB = 0 (all data have been deleted)

		err := urlStoreServer.Sync(ctx)
		require.NoError(t, err)

		require.Equal(t, 0, urlStoreServer.LatestVersion)
		require.Equal(t, 0, urlStoreServer.Store.Count())
	})

}

func TestURLStoreImpl_Match(t *testing.T) {
	t.Run("WHEN static url not exist", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		id, varMap := svc.Get("/gopher/doc.jpg")
		require.Nil(t, id)
		require.True(t, varMap.Empty())
	})

	t.Run("WHEN static url exist", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		id, varMap := svc.Get("/gopher/doc.png")
		require.Equal(t, "6", id)
		require.True(t, varMap.Empty())
	})

	t.Run("WHEN param url not exist", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		id, varMap := svc.Get("/users/def/abc")
		require.Nil(t, id)
		require.True(t, varMap.Empty())
	})

	t.Run("WHEN param url exist", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		id, varMap := svc.Get("/users/def/123")
		require.Equal(t, "12", id)
		require.Equal(t, 2, len(varMap.Keys()))
		require.Equal(t, "def", varMap.Map()["id"])
		require.Equal(t, "123", varMap.Map()["accnt"])
	})

	t.Run("WHEN more than 1 param exist in a subpath", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		id, varMap := svc.Get("/flight/src-abc-dst-def")
		require.Equal(t, "15", id)
		require.Equal(t, 2, len(varMap.Keys()))
		require.Equal(t, "abc", varMap.Map()["src"])
		require.Equal(t, "def", varMap.Map()["dst"])
	})
}

func TestURLStoreImpl_AddURL(t *testing.T) {
	t.Run("WHEN new static url added AND id not exist before", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		url := "/gopher/doc.jpg"
		svc.Insert(20, url)
		id, varMap := svc.Get(url)
		require.Equal(t, "20", id)
		require.True(t, varMap.Empty())
		require.Equal(t, 11, svc.Count())
	})

	t.Run("WHEN new static url added AND id exist before THEN double data added (with same id)", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		svc.Insert(20, "/gopher/old.jpg")
		svc.Insert(20, "/gopher/new.img")

		id, varMap := svc.Get("/gopher/new.img")
		require.Equal(t, "20", id)
		require.True(t, varMap.Empty())

		id, varMap = svc.Get("/gopher/old.jpg")
		require.Equal(t, "20", id)
		require.True(t, varMap.Empty())
		require.Equal(t, 12, svc.Count())
	})
}

func TestURLStoreImpl_UpdateURL(t *testing.T) {
	t.Run("WHEN existing static url updated with different url", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		svc.Update(6, "/gopher/updated.bmp")

		id, varMap := svc.Get("/gopher/old.png")
		require.Nil(t, id)
		require.True(t, varMap.Empty())

		id, varMap = svc.Get("/gopher/updated.bmp")
		require.Equal(t, "6", id)
		require.Equal(t, 0, len(varMap.Keys()))
		require.Equal(t, 10, svc.Count())
	})
}

func TestURLStoreImpl_DeleteURL(t *testing.T) {
	t.Run("WHEN existing static url deleted", func(t *testing.T) {
		svc := &urlstore.Service{Store: buildStore()}
		require.Equal(t, true, svc.Delete(6))

		id, varMap := svc.Get("/gopher/doc.png")
		require.Nil(t, id)
		require.True(t, varMap.Empty())
		require.Equal(t, false, svc.Delete(6))
		require.Equal(t, 9, svc.Count())
	})
}

func buildStore() urlstore.Store {
	pairs := []struct {
		id         int64
		key, value string
	}{
		{6, "/gopher/doc.png", "6"},
		{7, "/gopher/doc", "7"},
		{8, "/users/<id>", "8"},
		{9, "/users/<id>/profile", "9"},
		{10, "/users/<id>/<accnt:\\d+>/address", "10"},
		{11, "/users/<id>/age", "11"},
		{12, "/users/<id>/<accnt:\\d+>", "12"},
		{13, "/users/<id>/test/<name>", "13"},
		{14, "/users/abc/<id>/<name>", "14"},
		{15, "/flight/src-<src>-dst-<dst>", "15"},
	}

	store := urlstore.NewStore()
	for _, pair := range pairs {
		store.Add(pair.id, pair.key, pair.value)
	}

	return store
}
