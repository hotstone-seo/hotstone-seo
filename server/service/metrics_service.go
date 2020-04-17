package service

import (
	"github.com/hotstone-seo/hotstone-seo/server/metric"
	"go.uber.org/dig"
)

// MetricService contain logic for MetricsUnmatchedController [mock]
type MetricService interface {
	metric.MetricsRuleMatchingRepo
}

// MetricServiceImpl is implementation of MetricsRuleMatchingService
type MetricServiceImpl struct {
	dig.In
	metric.MetricsRuleMatchingRepo
}

// NewMetricService return new instance of MetricsRuleMatchingService [constructor]
func NewMetricService(impl MetricServiceImpl) MetricService {
	return &impl
}
