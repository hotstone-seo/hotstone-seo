package service

import (
	"github.com/hotstone-seo/hotstone-seo/analyt"
	"go.uber.org/dig"
)

// MetricService contain logic for MetricsUnmatchedController
// @mock
type MetricService interface {
	analyt.ReportRepo
}

// MetricServiceImpl is implementation of MetricsRuleMatchingService
type MetricServiceImpl struct {
	dig.In
	analyt.ReportRepo
}

// NewMetricService return new instance of MetricsRuleMatchingService
// @constructor
func NewMetricService(impl MetricServiceImpl) MetricService {
	return &impl
}
