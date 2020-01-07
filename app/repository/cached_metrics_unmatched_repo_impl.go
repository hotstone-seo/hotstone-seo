package repository

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// CachedMetricsUnmatchedRepoImpl is cached implementation of metrics_unmatched repository
type CachedMetricsUnmatchedRepoImpl struct {
	dig.In
	MetricsUnmatchedRepoImpl
	Redis *redis.Client
}

// Find metrics_unmatched entity
func (r *CachedMetricsUnmatchedRepoImpl) Find(ctx context.Context, id int64) (e *MetricsUnmatched, err error) {
	cacheKey := fmt.Sprintf("METRICS_UNMATCHED:FIND:%d", id)
	e = new(MetricsUnmatched)
	redisClient := r.Redis.WithContext(ctx)
	if err = dbkit.GetCache(redisClient, cacheKey, e); err == nil {
		log.Infof("Using cache %s", cacheKey)
		return
	}
	if e, err = r.MetricsUnmatchedRepoImpl.Find(ctx, id); err != nil {
		return
	}
	if err2 := dbkit.SetCache(redisClient, cacheKey, e, 20*time.Second); err2 != nil {
		log.Fatal(err2.Error())
	}
	return
}

// List of metrics_unmatched entity
func (r *CachedMetricsUnmatchedRepoImpl) List(ctx context.Context) (list []*MetricsUnmatched, err error) {
	cacheKey := fmt.Sprintf("METRICS_UNMATCHED:LIST")
	redisClient := r.Redis.WithContext(ctx)
	if err = dbkit.GetCache(redisClient, cacheKey, &list); err == nil {
		log.Infof("Using cache %s", cacheKey)
		return
	}
	if list, err = r.MetricsUnmatchedRepoImpl.List(ctx); err != nil {
		return
	}
	if err2 := dbkit.SetCache(redisClient, cacheKey, list, 20*time.Second); err2 != nil {
		log.Fatal(err2.Error())
	}
	return
}
