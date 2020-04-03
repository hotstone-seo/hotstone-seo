package cachekit_test

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/stretchr/testify/require"
)

func ExampleCache() {
	var (
		data   string
		server *miniredis.Miniredis
		err    error
	)

	// run redis server
	if server, err = miniredis.Run(); err != nil {
		log.Fatal(err.Error())
	}
	defer server.Close()

	// create redis client
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})

	// define key and refresh function for your cache
	cache := cachekit.New("some-key", func() (interface{}, error) {
		return "fresh-data", nil
	})

	// execute cache to get the data
	if err = cache.Execute(client, &data, cachekit.NewCacheControl()); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(data)

	// Output:
	// fresh-data

}

func TestCache(t *testing.T) {
	testRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer testRedis.Close()

	client := redis.NewClient(&redis.Options{Addr: testRedis.Addr()})

	t.Run("GIVEN cache not available", func(t *testing.T) {
		t.Run("WHEN refresh failed", func(t *testing.T) {
			var b bean
			cache := cachekit.New("key", func() (interface{}, error) {
				return nil, errors.New("some-refresh-error")
			})
			require.EqualError(t, cache.Execute(client, &b, cachekit.NewCacheControl()), "some-refresh-error")
		})
		t.Run("WHEN marshal failed", func(t *testing.T) {
			var b bean
			cache := cachekit.New("key", func() (interface{}, error) {
				return make(chan int), nil
			})
			require.EqualError(t, cache.Execute(client, &b, cachekit.NewCacheControl()), "json: unsupported type: chan int")
		})
		t.Run("WHEN failed to save to redis", func(t *testing.T) {
			var b bean
			badClient := redis.NewClient(&redis.Options{Addr: "wrong-addr"})
			cache := cachekit.New("key", func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})
			require.EqualError(t, cache.Execute(badClient, &b, cachekit.NewCacheControl()), "dial tcp: address wrong-addr: missing port in address")
		})
		t.Run("", func(t *testing.T) {
			var b bean
			cache := cachekit.New("key", func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})
			require.NoError(t, cache.Execute(client, &b, cachekit.NewCacheControl()))
			require.Equal(t, bean{Name: "new-name"}, b)
		})
	})
	t.Run("GIVEN cache available", func(t *testing.T) {
		testRedis.Set("key", `{"name":"cached"}`)
		var b bean
		cache := cachekit.New("key", func() (interface{}, error) {
			return &bean{Name: "new-name"}, nil
		})
		require.NoError(t, cache.Execute(client, &b, cachekit.NewCacheControl()))
		require.Equal(t, bean{Name: "cached"}, b)
	})
}

type bean struct {
	Name string
}
