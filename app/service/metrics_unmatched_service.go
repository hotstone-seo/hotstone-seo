package service

import (
	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// MetricsUnmatchedService contain logic for MetricsUnmatchedController
type MetricsUnmatchedService interface {
	repository.MetricsUnmatchedRepo
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
