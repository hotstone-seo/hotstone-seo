package cachekit_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
)

func TestCacheStore(t *testing.T) {
	testRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer testRedis.Close()

	client := redis.NewClient(&redis.Options{Addr: testRedis.Addr()})
	store := cachekit.New(client)
	ctx := context.Background()

	t.Run("GIVEN cache not available", func(t *testing.T) {
		t.Run("WHEN refresh failed", func(t *testing.T) {
			var b bean
			_, err := store.Retrieve(ctx, "key", &b, func() (interface{}, error) {
				return nil, errors.New("some-refresh-error")
			})
			require.EqualError(t, err, "some-refresh-error")
		})
		t.Run("WHEN marshal failed", func(t *testing.T) {
			var b bean
			_, err := store.Retrieve(ctx, "key", &b, func() (interface{}, error) {
				return make(chan int), nil
			})
			require.EqualError(t, err, "json: unsupported type: chan int")
		})
		t.Run("WHEN failed to save to redis", func(t *testing.T) {
			var b bean
			broken := cachekit.New(redis.NewClient(&redis.Options{Addr: "wrong-addr"}))
			_, err := broken.Retrieve(ctx, "key", &b, func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})
			require.EqualError(t, err, "dial tcp: address wrong-addr: missing port in address")
		})
		t.Run("", func(t *testing.T) {
			var b bean
			fromCache, err := store.Retrieve(ctx, "key", &b, func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})
			require.NoError(t, err)
			require.Equal(t, bean{Name: "new-name"}, b)
			require.False(t, fromCache)
		})
	})
	t.Run("GIVEN cache available", func(t *testing.T) {
		testRedis.Set("key", `{"name":"cached"}`)
		var b bean
		fromCache, err := store.Retrieve(ctx, "key", &b, func() (interface{}, error) {
			return &bean{Name: "new-name"}, nil
		})
		require.NoError(t, err)
		require.Equal(t, bean{Name: "cached"}, b)
		require.True(t, fromCache)
	})
}

type bean struct {
	Name string
}
