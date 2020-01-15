package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// CachedTagRepoImpl is cached implementation of tag repository
type CachedTagRepoImpl struct {
	dig.In
	TagRepoImpl
	Redis *redis.Client
}

// FindOne tag
func (r *CachedTagRepoImpl) FindOne(ctx context.Context, id int64) (e *Tag, err error) {
	cacheKey := fmt.Sprintf("TAGS:FIND:%d", id)
	e = new(Tag)
	redisClient := r.Redis.WithContext(ctx)
	if err = dbkit.GetCache(redisClient, cacheKey, e); err == nil {
		log.Infof("Using cache %s", cacheKey)
		return
	}
	if e, err = r.TagRepoImpl.FindOne(ctx, id); err != nil {
		return
	}
	if err2 := dbkit.SetCache(redisClient, cacheKey, e, 20*time.Second); err2 != nil {
		log.Fatal(err2.Error())
	}
	return
}

// Find tags
func (r *CachedTagRepoImpl) Find(ctx context.Context, filter TagFilter) (list []*Tag, err error) {
	// TODO: Provide different key for each filter applied
	var (
		filterKey   []byte
		cacheKey    string
		redisClient = r.Redis.WithContext(ctx)
	)
	if filterKey, err = json.Marshal(filter); err == nil {
		cacheKey = fmt.Sprintf("TAGS:%s", string(filterKey))
		if err = dbkit.GetCache(redisClient, cacheKey, &list); err == nil {
			log.Infof("Using cache %s", cacheKey)
			return
		}
	}
	if list, err = r.TagRepoImpl.Find(ctx, filter); err != nil {
		return
	}
	if err2 := dbkit.SetCache(redisClient, cacheKey, list, 20*time.Second); err2 != nil {
		log.Fatal(err2.Error())
	}
	return
}
