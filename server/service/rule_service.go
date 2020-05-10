package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// RuleService contain logic for Rule Controller
// @mock
type RuleService interface {
	FindOne(ctx context.Context, id int64) (*repository.Rule, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.Rule, error)
	Insert(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule repository.Rule) error
}

// RuleServiceImpl is implementation of RuleService
type RuleServiceImpl struct {
	dig.In
	RuleRepo    repository.RuleRepo
	URLSyncRepo repository.URLSyncRepo
	AuditTrailService
	HistoryService
	dbtxn.Transactional
}

// NewRuleService return new instance of RuleService
// @constructor
func NewRuleService(impl RuleServiceImpl) RuleService {
	return &impl
}

// FindOne rule
func (r *RuleServiceImpl) FindOne(ctx context.Context, id int64) (rule *repository.Rule, err error) {
	return r.RuleRepo.FindOne(ctx, id)
}

// Find rule
func (r *RuleServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) (list []*repository.Rule, err error) {
	return r.RuleRepo.Find(ctx, paginationParam)
}

// Insert rule
func (r *RuleServiceImpl) Insert(ctx context.Context, rule repository.Rule) (newRuleID int64, err error) {
	if rule.ID, err = r.RuleRepo.Insert(ctx, rule); err != nil {
		return
	}
	go func() {
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"rules",
			rule.ID,
			repository.Insert,
			nil,
			rule,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return rule.ID, nil
}

// Delete rule
func (r *RuleServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer r.CommitMe(&ctx)()
	oldRule, err := r.RuleRepo.FindOne(ctx, id)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.RuleRepo.Delete(ctx, id); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.URLSyncRepo.Insert(ctx, repository.URLSync{
		Operation:        "DELETE",
		RuleID:           id,
		LatestURLPattern: nil,
	}); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	go func() {
		if _, histErr := r.HistoryService.RecordHistory(
			ctx,
			"rules",
			id,
			oldRule,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"rules",
			id,
			repository.Delete,
			oldRule,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

// Update rule
func (r *RuleServiceImpl) Update(ctx context.Context, rule repository.Rule) (err error) {
	defer r.CommitMe(&ctx)()
	var (
		urlSync *repository.URLSync
	)
	oldRule, err := r.RuleRepo.FindOne(ctx, rule.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.RuleRepo.Update(ctx, rule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if urlSync, err = r.URLSyncRepo.FindRule(ctx, rule.ID); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if syncOP := syncOperation(rule, urlSync); syncOP != "" {
		if _, err = r.URLSyncRepo.Insert(ctx, repository.URLSync{
			Operation:        syncOP,
			RuleID:           rule.ID,
			LatestURLPattern: &rule.URLPattern,
		}); err != nil {
			r.CancelMe(ctx, err)
			return
		}
	}
	go func() {
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"rules",
			rule.ID,
			repository.Update,
			oldRule,
			rule,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

func syncOperation(rule repository.Rule, lastURLSync *repository.URLSync) string {
	if lastURLSync == nil || lastURLSync.Operation == "DELETE" {
		if rule.Status == "start" {
			return "INSERT"
		}
		return ""
	}
	if rule.Status == "stop" {
		return "DELETE"
	}
	return "UPDATE"
}
