package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"go.uber.org/dig"
)

// MetricsUnmatchedService contain logic for MetricsUnmatchedController
type MetricsUnmatchedService interface {
	repository.MetricsUnmatchedRepo

	Record(ctx context.Context, requestPath string) error
}

// MetricsUnmatchedServiceImpl is implementation of MetricsUnmatchedService
type MetricsUnmatchedServiceImpl struct {
	dig.In
	repository.MetricsUnmatchedRepo
}

// NewMetricsUnmatchedService return new instance of MetricsUnmatchedService
func NewMetricsUnmatchedService(impl MetricsUnmatchedServiceImpl) MetricsUnmatchedService {
	return &impl
}

func (s *MetricsUnmatchedServiceImpl) Record(ctx context.Context, requestPath string) (err error) {
	_, err = s.MetricsUnmatchedRepo.Insert(ctx, repository.MetricsUnmatched{RequestPath: requestPath})
	return
}
