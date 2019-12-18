package service

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// RuleService contain logic for Rule Controller
type RuleService interface {
	// repository.RuleRepo
	Find(ctx context.Context, id int64) (*repository.Rule, error)
	List(ctx context.Context) ([]*repository.Rule, error)
	Insert(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule repository.Rule) error
}

// RuleServiceImpl is implementation of RuleService
type RuleServiceImpl struct {
	dig.In
	RuleRepo repository.RuleRepo
}

// NewRuleService return new instance of RuleService
func NewRuleService(impl RuleServiceImpl) RuleService {
	return &impl
}

// Find rule
func (r *RuleServiceImpl) Find(ctx context.Context, id int64) (rule *repository.Rule, err error) {
	return r.RuleRepo.Find(ctx, id)
}

// List rule
func (r *RuleServiceImpl) List(ctx context.Context) (list []*repository.Rule, err error) {
	return r.RuleRepo.List(ctx)
}

// Insert rule
func (r *RuleServiceImpl) Insert(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error) {
	log.Warn(">>> rule.INSERT")
	err = repository.WithTransaction(r.RuleRepo.DB(), func(tx *sql.Tx) error {
		lastInsertID, err = r.RuleRepo.Insert(ctx, rule)
		log.Warnf("#1 lastInsertID: %d", lastInsertID)
		if err != nil {
			return err
		}

		// if true {
		// 	return errors.New("DUMMY ERROR")
		// }

		lastInsertID, err = r.RuleRepo.Insert(repository.SetTx(ctx, tx), rule)
		// lastInsertID, err = r.RuleRepo.Insert(ctx, rule)
		log.Warnf("#2 lastInsertID: %d", lastInsertID)
		if err != nil {
			return err
		}

		return nil
	})

	// lastInsertID, err = r.RuleRepo.Insert(ctx, rule)
	// log.Warnf("#1 lastInsertID: %d", lastInsertID)
	// if err != nil {
	// 	return
	// }

	// if true {
	// 	err = errors.New("DUMMY ERROR")
	// 	return
	// }

	return
}

// Delete rule
func (r *RuleServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	return r.RuleRepo.Delete(ctx, id)
}

// Update rule
func (r *RuleServiceImpl) Update(ctx context.Context, rule repository.Rule) (err error) {
	return r.RuleRepo.Update(ctx, rule)
}
