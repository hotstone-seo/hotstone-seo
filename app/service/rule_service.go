package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// RuleService contain logic for Rule Controller
type RuleService interface {
	repository.RuleRepo
	InsertToDBAndStore(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error)
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

func (s *RuleServiceImpl) InsertToDBAndStore(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error) {
	return 0, nil
}
