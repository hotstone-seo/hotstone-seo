package service

import (
	"github.com/hotstone-seo/hotstone-seo/internal/analyt"
	"go.uber.org/dig"
)

// MetricService contain logic for MetricsUnmatchedController
// @mock
type MetricService interface {
	analyt.ReportRepo
	analyt.ClientKeyAnalytRepo
}

// MetricServiceImpl is implementation of MetricsRuleMatchingService
type MetricServiceImpl struct {
	dig.In
	analyt.ReportRepo
	analyt.ClientKeyAnalytRepo
}

// NewMetricService return new instance of MetricsRuleMatchingService
// @constructor
func NewMetricService(impl MetricServiceImpl) MetricService {
	return &impl
}
