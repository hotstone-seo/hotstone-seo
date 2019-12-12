package service

import (
	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// RuleService contain logic for Rule Controller
type RuleService interface {
	repository.RuleRepo
}

// RuleServiceImpl is implementation of RuleService
type RuleServiceImpl struct {
	dig.In
	repository.RuleRepo
}

// NewRuleService return new instance of RuleService
func NewRuleService(impl RuleServiceImpl) RuleService {
	return &impl
}
