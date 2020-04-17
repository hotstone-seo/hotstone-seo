package metric

import (
	"go.uber.org/dig"
)

// MetricsRuleMatchingService contain logic for MetricsUnmatchedController [mock]
type MetricsRuleMatchingService interface {
	MetricsRuleMatchingRepo

	SetMatched(m *MetricsRuleMatching, matchedURL string, ruleID int64)
	SetMismatched(m *MetricsRuleMatching, mismatchedURL string)
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

// SetMatched to set matched
func (s *MetricsRuleMatchingServiceImpl) SetMatched(m *MetricsRuleMatching, matchedURL string, ruleID int64) {
	m.IsMatched = 1
	m.URL = &matchedURL
	m.RuleID = &ruleID
}

// SetMismatched to set mismatched
func (s *MetricsRuleMatchingServiceImpl) SetMismatched(m *MetricsRuleMatching, mismatchedURL string) {
	m.IsMatched = 0
	m.URL = &mismatchedURL
}
