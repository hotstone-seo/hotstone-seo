package service

import (
	"context"

	"github.com/go-redis/redis"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// ProviderService contain logic for provider api [mock]
type ProviderService interface {
	Match(context.Context, MatchRequest) (*MatchResponse, error)
	FetchTags(ctx context.Context, id int64, locale string) ([]*ITag, error)
}

// ProviderServiceImpl is implementation of Provider
type ProviderServiceImpl struct {
	dig.In
	MetricsRuleMatchingService
	repository.DataSourceRepo
	repository.RuleRepo
	repository.TagRepo
	URLService

	Redis *redis.Client
}

// ITag is tag after interpolate with data
type ITag repository.Tag

// IDataSource is datasource after interpolate with data
type IDataSource repository.DataSource

// NewProviderService return new instance of ProviderService [constructor]
func NewProviderService(impl ProviderServiceImpl) ProviderService {
	return &impl
}
