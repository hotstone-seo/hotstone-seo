package cachekit_test

import (
	"fmt"
	"log"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
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
	if err = cache.Execute(client, &data, pragmaWithCacheControl("")); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(data)

	// Output:
	// fresh-data

}
