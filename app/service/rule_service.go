package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/app/urlstore"
	"go.uber.org/dig"
)

// RuleService contain logic for Rule Controller
type RuleService interface {
	// repository.RuleRepo
	FindOne(ctx context.Context, id int64) (*repository.Rule, error)
	Find(ctx context.Context) ([]*repository.Rule, error)
	Insert(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule repository.Rule) error
}

// RuleServiceImpl is implementation of RuleService
type RuleServiceImpl struct {
	dig.In
	RuleRepo    repository.RuleRepo
	URLSyncRepo urlstore.URLStoreSyncRepo
	repository.Transactional
}

// NewRuleService return new instance of RuleService
func NewRuleService(impl RuleServiceImpl) RuleService {
	return &impl
}

// Find rule
func (r *RuleServiceImpl) FindOne(ctx context.Context, id int64) (rule *repository.Rule, err error) {
	return r.RuleRepo.FindOne(ctx, id)
}

// List rule
func (r *RuleServiceImpl) Find(ctx context.Context) (list []*repository.Rule, err error) {
	return r.RuleRepo.Find(ctx)
}

// Insert rule
func (r *RuleServiceImpl) Insert(ctx context.Context, rule repository.Rule) (newRuleID int64, err error) {
	defer r.CommitMe(&ctx)()
	if newRuleID, err = r.RuleRepo.Insert(ctx, rule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.URLSyncRepo.Insert(ctx, urlstore.URLStoreSync{
		Operation:        "INSERT",
		RuleID:           newRuleID,
		LatestURLPattern: &rule.UrlPattern,
	}); err != nil {
		r.CancelMe(ctx, err)
		return newRuleID, err
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
	if _, err = r.URLSyncRepo.Insert(ctx, urlstore.URLStoreSync{
		Operation:        "DELETE",
		RuleID:           id,
		LatestURLPattern: nil,
	}); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}

// Update rule
func (r *RuleServiceImpl) Update(ctx context.Context, rule repository.Rule) (err error) {
	defer r.CommitMe(&ctx)()
	if err = r.RuleRepo.Update(ctx, rule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.URLSyncRepo.Insert(ctx, urlstore.URLStoreSync{
		Operation:        "UPDATE",
		RuleID:           rule.ID,
		LatestURLPattern: &rule.UrlPattern,
	}); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}
