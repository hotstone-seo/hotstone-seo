package service

import (
	"context"
	"encoding/json"

	"github.com/fatih/structs"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
)

type (
	// RuleService contains logic for Rule Controller
	// @mock
	RuleService interface {
		FindOne(ctx context.Context, id int64) (*repository.Rule, error)
		Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.Rule, error)
		Insert(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, rule repository.Rule) error
		Patch(ctx context.Context, ruleID int64, fields map[string]interface{}) error
	}
	// RuleServiceImpl is an implementation of RuleService
	RuleServiceImpl struct {
		dig.In
		RuleRepo   repository.RuleRepo
		SyncRepo   urlstore.SyncRepo
		AuditTrail AuditTrailService
		dbtxn.Transactional
	}
)

// NewRuleService creates and returns new instance of RuleService
// @ctor
func NewRuleService(impl RuleServiceImpl) RuleService {
	return &impl
}

// FindOne returns a single Rule by its ID
func (r *RuleServiceImpl) FindOne(ctx context.Context, id int64) (rule *repository.Rule, err error) {
	return r.RuleRepo.FindOne(ctx, id)
}

// Find returns multiple Rule based on pagination criterion
func (r *RuleServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) (list []*repository.Rule, err error) {
	return r.RuleRepo.Find(ctx, paginationParam)
}

// Insert creates a new Rule on the persistent storage configured for the service
func (r *RuleServiceImpl) Insert(ctx context.Context, rule repository.Rule) (int64, error) {
	lastInsertedID, err := r.RuleRepo.Insert(ctx, rule)
	if err != nil {
		return -1, err
	}

	r.AuditTrail.RecordInsert(ctx, "rules", lastInsertedID, rule)
	return lastInsertedID, nil
}

// Update replaces the values of an existing Rule in the persistent storage by a new Rule
func (r *RuleServiceImpl) Update(ctx context.Context, rule repository.Rule) (err error) {
	defer r.BeginTxn(&ctx)()
	var (
		Sync *urlstore.Sync
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
	if Sync, err = r.SyncRepo.FindRule(ctx, rule.ID); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if syncOP := syncOperation(rule, Sync); syncOP != "" {
		if _, err = r.SyncRepo.Insert(ctx, urlstore.Sync{
			Operation:        syncOP,
			RuleID:           rule.ID,
			LatestURLPattern: &rule.URLPattern,
		}); err != nil {
			r.CancelMe(ctx, err)
			return
		}
	}

	r.AuditTrail.RecordUpdate(ctx, "rules", rule.ID, oldRule, rule)
	return nil
}

// Patch updates only selected fields to an existing Rule in the persistent storage
func (r *RuleServiceImpl) Patch(ctx context.Context, ruleID int64, fields map[string]interface{}) (err error) {
	existingRule, err := r.RuleRepo.FindOne(ctx, ruleID)
	if err != nil {
		return
	}
	targetMap := structs.Map(existingRule)
	for k, v := range fields {
		targetMap[k] = v
	}
	j, err := json.Marshal(targetMap)
	if err != nil {
		return
	}
	var newRule repository.Rule
	if err = json.Unmarshal(j, &newRule); err != nil {
		return
	}
	return r.Update(ctx, newRule)
}

// Delete removes the Rule entry from persistent storage configured for the service
func (r *RuleServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer r.BeginTxn(&ctx)()
	oldRule, err := r.RuleRepo.FindOne(ctx, id)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.RuleRepo.Delete(ctx, id); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.SyncRepo.Insert(ctx, urlstore.Sync{
		Operation:        "DELETE",
		RuleID:           id,
		LatestURLPattern: nil,
	}); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	r.AuditTrail.RecordDelete(ctx, "rules", id, oldRule)
	return nil
}

func syncOperation(rule repository.Rule, lastSync *urlstore.Sync) string {
	if lastSync == nil || lastSync.Operation == "DELETE" {
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
