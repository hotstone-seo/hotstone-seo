package urlstore_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-server/app/urlstore"
	"github.com/stretchr/testify/require"
)

func TestURLStoreImpl_GetURL(t *testing.T) {
	t.Run("WHEN static url not exist", func(t *testing.T) {
		store := buildStore(t)
		id, varMap := store.Get("/gopher/doc.jpg")
		require.Equal(t, -1, id)
		require.Empty(t, varMap)
	})

	t.Run("WHEN static url exist", func(t *testing.T) {
		store := buildStore(t)
		id, varMap := store.Get("/gopher/doc.png")
		require.Equal(t, 6, id)
		require.Empty(t, varMap)
	})

	t.Run("WHEN param url not exist", func(t *testing.T) {
		store := buildStore(t)
		id, varMap := store.Get("/users/def/abc")
		require.Equal(t, -1, id)
		require.Empty(t, varMap)
	})

	t.Run("WHEN param url exist", func(t *testing.T) {
		store := buildStore(t)
		id, varMap := store.Get("/users/def/123")
		require.Equal(t, 12, id)
		require.Equal(t, 2, len(varMap))
		require.Equal(t, "def", varMap["id"])
		require.Equal(t, "123", varMap["accnt"])
	})
}

func TestURLStoreImpl_AddURL(t *testing.T) {

	t.Run("WHEN new static url added AND id not exist before", func(t *testing.T) {
		store := buildStore(t)
		url := "/gopher/doc.jpg"
		store.Add(20, url)
		id, varMap := store.Get(url)
		require.Equal(t, 20, id)
		require.Empty(t, varMap)
		require.Equal(t, 10, store.Count())
	})

	t.Run("WHEN new static url added AND id exist before THEN double data added (with same id)", func(t *testing.T) {
		store := buildStore(t)
		store.Add(20, "/gopher/old.jpg")
		store.Add(20, "/gopher/new.img")

		id, varMap := store.Get("/gopher/new.img")
		require.Equal(t, 20, id)
		require.Empty(t, varMap)

		id, varMap = store.Get("/gopher/old.jpg")
		require.Equal(t, 20, id)
		require.Empty(t, varMap)
		require.Equal(t, 11, store.Count())
	})
}

func TestURLStoreImpl_UpdateURL(t *testing.T) {
	t.Run("WHEN existing static url updated with different url", func(t *testing.T) {
		store := buildStore(t)
		store.Update(6, "/gopher/updated.bmp")

		id, varMap := store.Get("/gopher/old.png")
		require.Equal(t, -1, id)
		require.Empty(t, varMap)

		id, varMap = store.Get("/gopher/updated.bmp")
		require.Equal(t, 6, id)
		require.Equal(t, 0, len(varMap))
		require.Equal(t, 9, store.Count())
	})
}

func TestURLStoreImpl_DeleteURL(t *testing.T) {
	t.Run("WHEN existing static url deleted", func(t *testing.T) {
		store := buildStore(t)
		require.Equal(t, true, store.Delete(6))

		id, varMap := store.Get("/gopher/doc.png")
		require.Equal(t, -1, id)
		require.Empty(t, varMap)
		require.Equal(t, false, store.Delete(6))
		require.Equal(t, 8, store.Count())
	})
}

func buildStore(t *testing.T) urlstore.URLStore {
	t.Helper()

	pairs := []struct {
		id         int
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
	}

	store := urlstore.NewURLStoreTree()

	for _, pair := range pairs {
		store.Add(pair.id, pair.key, pair.value)
	}

	require.Equal(t, 9, store.Count())

	return &urlstore.URLStoreImpl{URLStoreTree: store}
}
