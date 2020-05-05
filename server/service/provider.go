package service

import (
	"context"
	"net/url"

	"github.com/go-redis/redis"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/hotstone-seo/hotstone-seo/server/metric"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/urlstore"
	"go.uber.org/dig"
)

// ProviderService contain logic for provider api
// @mock
type ProviderService interface {
	Match(context.Context, url.Values) (*MatchResponse, error)
	FetchTags(context.Context, url.Values) ([]*ITag, error)
	FetchTagsWithCache(context.Context, url.Values, *cachekit.Pragma) ([]*ITag, error)
}

// ProviderServiceImpl is implementation of Provider
type ProviderServiceImpl struct {
	dig.In
	metric.RuleMatchingRepo
	repository.DataSourceRepo
	repository.RuleRepo
	repository.TagRepo
	repository.StructuredDataRepo

	Redis *redis.Client
	urlstore.Store
}

// ITag is tag after interpolate with data
type ITag repository.Tag

// IDataSource is datasource after interpolate with data
type IDataSource repository.DataSource

// NewProviderService return new instance of ProviderService
// @constructor
func NewProviderService(impl ProviderServiceImpl) ProviderService {
	return &impl
}
