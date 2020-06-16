package service

import (
	"context"
	"encoding/json"

	"github.com/fatih/structs"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// RuleService contains logic for Rule Controller
// @mock
type RuleService interface {
	FindOne(ctx context.Context, id int64) (*repository.Rule, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.Rule, error)
	Insert(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule repository.Rule) error
	Patch(ctx context.Context, ruleID int64, fields map[string]interface{}) error
}

// RuleServiceImpl is an implementation of RuleService
type RuleServiceImpl struct {
	dig.In
	RuleRepo repository.RuleRepo
	SyncRepo urlstore.SyncRepo
	AuditTrailService
	HistoryService
	dbtxn.Transactional
}

// NewRuleService creates and returns new instance of RuleService
// @constructor
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
func (r *RuleServiceImpl) Insert(ctx context.Context, rule repository.Rule) (newID int64, err error) {
	defer func() {
		if err != nil {
			return
		}
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			Record{
				EntityName: "rules",
				EntityID:   newID,
				Operation:  InsertOp,
				PrevData:   nil,
				NextData:   rule,
			},
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return r.RuleRepo.Insert(ctx, rule)
}

// Update replaces the values of an existing Rule in the persistent storage by a new Rule
// TODO: Make updating URL store clearer
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
	go func() {
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			Record{
				EntityName: "rules",
				EntityID:   rule.ID,
				Operation:  UpdateOp,
				PrevData:   oldRule,
				NextData:   rule,
			},
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
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
			Record{
				EntityName: "rules",
				EntityID:   id,
				Operation:  DeleteOp,
				PrevData:   oldRule,
				NextData:   nil,
			},
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
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
