package metric

import (
	"go.uber.org/dig"
)

// MetricsRuleMatchingService contain logic for MetricsUnmatchedController [mock]
type MetricsRuleMatchingService interface {
	MetricsRuleMatchingRepo
}

// MetricsRuleMatchingServiceImpl is implementation of MetricsRuleMatchingService
type MetricsRuleMatchingServiceImpl struct {
	dig.In
	MetricsRuleMatchingRepo
}

// NewMetricsRuleMatchingService return new instance of MetricsRuleMatchingService [constructor]
func NewMetricsRuleMatchingService(impl MetricsRuleMatchingServiceImpl) MetricsRuleMatchingService {
	return &impl
}
