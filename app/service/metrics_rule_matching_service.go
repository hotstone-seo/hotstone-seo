package service

import (
	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"go.uber.org/dig"
)

// MetricsRuleMatchingService contain logic for MetricsUnmatchedController
type MetricsRuleMatchingService interface {
	repository.MetricsRuleMatchingRepo

	SetMatched(m *repository.MetricsRuleMatching)
	SetMismatched(m *repository.MetricsRuleMatching, mismatchedURL string)
}

// MetricsRuleMatchingServiceImpl is implementation of MetricsRuleMatchingService
type MetricsRuleMatchingServiceImpl struct {
	dig.In
	repository.MetricsRuleMatchingRepo
}

// NewMetricsRuleMatchingService return new instance of MetricsRuleMatchingService
func NewMetricsRuleMatchingService(impl MetricsRuleMatchingServiceImpl) MetricsRuleMatchingService {
	return &impl
}

func (s *MetricsRuleMatchingServiceImpl) SetMatched(m *repository.MetricsRuleMatching) {
	m.IsMatched = 1
}

func (s *MetricsRuleMatchingServiceImpl) SetMismatched(m *repository.MetricsRuleMatching, mismatchedURL string) {
	m.IsMatched = 0
	m.URLMismatched = &mismatchedURL
}
