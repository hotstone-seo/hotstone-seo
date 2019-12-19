package service

import (
	"context"
	"database/sql"

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

	err = repository.WithTransaction(r.RuleRepo.DB(), func(tx *sql.Tx) error {
		rule, err = r.RuleRepo.Find(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// List rule
func (r *RuleServiceImpl) List(ctx context.Context) (list []*repository.Rule, err error) {
	err = repository.WithTransaction(r.RuleRepo.DB(), func(tx *sql.Tx) error {
		list, err = r.RuleRepo.List(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// Insert rule
func (r *RuleServiceImpl) Insert(ctx context.Context, rule repository.Rule) (lastInsertID int64, err error) {
	err = repository.WithTransaction(r.RuleRepo.DB(), func(tx *sql.Tx) error {
		lastInsertID, err = r.RuleRepo.Insert(ctx, tx, rule)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// Delete rule
func (r *RuleServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	err = repository.WithTransaction(r.RuleRepo.DB(), func(tx *sql.Tx) error {
		err = r.RuleRepo.Delete(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// Update rule
func (r *RuleServiceImpl) Update(ctx context.Context, rule repository.Rule) (err error) {

	err = repository.WithTransaction(r.RuleRepo.DB(), func(tx *sql.Tx) error {
		err = r.RuleRepo.Update(ctx, tx, rule)
		if err != nil {
			return err
		}

		return nil
	})

	return
}
