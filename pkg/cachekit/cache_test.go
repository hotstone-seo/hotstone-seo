package cachekit_test

import (
	"errors"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/stretchr/testify/require"
)

func TestCache_CacheNoAvailable(t *testing.T) {
	var target bean

	testRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer testRedis.Close()

	client := redis.NewClient(&redis.Options{Addr: testRedis.Addr()})

	t.Run("WHEN refresh failed", func(t *testing.T) {
		cache := cachekit.New("key", func() (interface{}, error) {
			return nil, errors.New("some-error")
		})
		require.EqualError(t, cache.Execute(client, &target, pragmaWithCacheControl("")), "some-error")
	})
	t.Run("WHEN marshal failed", func(t *testing.T) {
		cache := cachekit.New("key", func() (interface{}, error) {
			return make(chan int), nil
		})
		require.EqualError(t, cache.Execute(client, &target, pragmaWithCacheControl("")), "Cache: json: unsupported type: chan int")
	})
	t.Run("WHEN failed to save to redis", func(t *testing.T) {
		badClient := redis.NewClient(&redis.Options{Addr: "wrong-addr"})
		cache := cachekit.New("key", func() (interface{}, error) {
			return &bean{Name: "new-name"}, nil
		})
		require.EqualError(t, cache.Execute(badClient, &target, pragmaWithCacheControl("")), "Cache: dial tcp: address wrong-addr: missing port in address")
	})
	t.Run("", func(t *testing.T) {
		// monkey patch time.Now
		defer monkey.Patch(time.Now, func() time.Time {
			return time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
		}).Unpatch()

		cache := cachekit.New("key", func() (interface{}, error) {
			return &bean{Name: "new-name"}, nil
		})

		pragma := pragmaWithCacheControl("")
		require.NoError(t, cache.Execute(client, &target, pragma))

		// check target
		require.Equal(t, bean{Name: "new-name"}, target)

		// check data in redis
		require.Equal(t, `{"Name":"new-name"}`, client.Get("key").Val())
		require.Equal(t, 30*time.Second, client.TTL("key").Val())

		// check pragma
		require.Equal(t, "Thu, 16 Feb 2017 00:00:30 GMT", pragma.ResponseHeaders()[cachekit.HeaderExpires])
		require.Equal(t, "Thu, 16 Feb 2017 00:00:00 GMT", pragma.ResponseHeaders()[cachekit.HeaderLastModified])
	})
}

func TestCache_CacheAvailable(t *testing.T) {

	var target bean

	testRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer testRedis.Close()

	client := redis.NewClient(&redis.Options{Addr: testRedis.Addr()})

	t.Run("", func(t *testing.T) {
		// monkey patching time.Now
		defer monkey.Patch(time.Now, func() time.Time {
			return time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
		}).Unpatch()

		// set cache n redis
		testRedis.Set("key", `{"name":"cached"}`)
		testRedis.Set("key:time", "Wed, 15 Feb 2017 23:55:00 GMT")
		testRedis.SetTTL("key", 10*time.Second)

		cache := cachekit.New("key", func() (interface{}, error) {
			return &bean{Name: "new-name"}, nil
		})

		pragma := pragmaWithCacheControl("")
		require.NoError(t, cache.Execute(client, &target, pragma))

		// Check target
		require.Equal(t, bean{Name: "cached"}, target)

		// check pragma
		require.Equal(t, "Thu, 16 Feb 2017 00:00:10 GMT", pragma.ResponseHeaders()[cachekit.HeaderExpires])
		require.Equal(t, "Wed, 15 Feb 2017 23:55:00 GMT", pragma.ResponseHeaders()[cachekit.HeaderLastModified])
	})
	t.Run("WHEN cache-control: no-cache", func(t *testing.T) {
		testRedis.Set("key", `{"name":"cached"}`)
		cache := cachekit.New("key", func() (interface{}, error) {
			return &bean{Name: "new-name"}, nil
		})

		require.NoError(t, cache.Execute(client, &target, pragmaWithCacheControl("no-cache")))
		require.Equal(t, bean{Name: "new-name"}, target)

		require.Equal(t, `{"Name":"new-name"}`, client.Get("key").Val())
		require.Equal(t, 30*time.Second, client.TTL("key").Val())
	})
}

func TestCache_IfModifiedSince(t *testing.T) {

	testRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer testRedis.Close()

	var target bean
	client := redis.NewClient(&redis.Options{Addr: testRedis.Addr()})

	testcases := []struct {
		lastModified       string
		ifModifiedSince    string
		expectedNoModified bool
	}{
		{
			lastModified:       "Wed, 15 Feb 2017 23:55:00 GMT",
			ifModifiedSince:    "Wed, 15 Feb 2017 23:58:00 GMT",
			expectedNoModified: true,
		},
		{
			lastModified:       "Wed, 15 Feb 2017 23:55:00 GMT",
			ifModifiedSince:    "Wed, 15 Feb 2017 23:50:00 GMT",
			expectedNoModified: false,
		},
	}

	for _, tt := range testcases {
		// set cache n redis
		testRedis.Set("key", `{"name":"cached"}`)
		testRedis.Set("key:time", tt.lastModified)
		testRedis.SetTTL("key", 10*time.Second)
		testRedis.SetTTL("key:time", 10*time.Second)

		cache := cachekit.New("key", func() (interface{}, error) {
			return &bean{Name: "new-name"}, nil
		})

		err := cache.Execute(client, &target, pragmaWithIfModifiedSince(tt.ifModifiedSince))
		require.Equal(t, tt.expectedNoModified, cachekit.NotModifiedError(err))
	}

}

type bean struct {
	Name string
}
