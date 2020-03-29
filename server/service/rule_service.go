package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// RuleService contain logic for Rule Controller [mock]
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
	repository.Transactional
}

// NewRuleService return new instance of RuleService [constructor]
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
	defer r.CommitMe(&ctx)()
	if newRuleID, err = r.RuleRepo.Insert(ctx, rule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.URLSyncRepo.Insert(ctx, repository.URLSync{
		Operation:        "INSERT",
		RuleID:           newRuleID,
		LatestURLPattern: &rule.UrlPattern,
	}); err != nil {
		r.CancelMe(ctx, err)
		return newRuleID, err
	}
	newRule, err := r.RuleRepo.FindOne(ctx, newRuleID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "rules", newRuleID, repository.Insert, nil, newRule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return newRuleID, nil
}

// Delete rule
func (r *RuleServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer r.CommitMe(&ctx)()
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
	oldRule, err := r.RuleRepo.FindOne(ctx, id)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "rules", id, repository.Delete, oldRule, nil); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}

// Update rule
func (r *RuleServiceImpl) Update(ctx context.Context, rule repository.Rule) (err error) {
	defer r.CommitMe(&ctx)()
	oldRule, err := r.RuleRepo.FindOne(ctx, rule.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.RuleRepo.Update(ctx, rule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.URLSyncRepo.Insert(ctx, repository.URLSync{
		Operation:        "UPDATE",
		RuleID:           rule.ID,
		LatestURLPattern: &rule.UrlPattern,
	}); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	newRule, err := r.RuleRepo.FindOne(ctx, rule.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "rules", rule.ID, repository.Update, oldRule, newRule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}
