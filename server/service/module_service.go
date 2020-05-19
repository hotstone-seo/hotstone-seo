package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// ModuleService contain logic for Module Controller
// @mock
type ModuleService interface {
	FindOne(ctx context.Context, id int64) (*repository.Module, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.Module, error)
	Insert(ctx context.Context, module repository.Module) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, module repository.Module) error
}

// ModuleServiceImpl is implementation of ModuleService
type ModuleServiceImpl struct {
	dig.In
	ModuleRepo repository.ModuleRepo
	dbtxn.Transactional
	AuditTrailService
	HistoryService
}

// NewModuleService return new instance of ModuleService
// @constructor
func NewModuleService(impl ModuleServiceImpl) ModuleService {
	return &impl
}

// FindOne module
func (r *ModuleServiceImpl) FindOne(ctx context.Context, id int64) (module *repository.Module, err error) {
	return r.ModuleRepo.FindOne(ctx, id)
}

// Find module
func (r *ModuleServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) (list []*repository.Module, err error) {
	return r.ModuleRepo.Find(ctx, paginationParam)
}

// Insert module
func (r *ModuleServiceImpl) Insert(ctx context.Context, module repository.Module) (newUserID int64, err error) {
	defer r.CommitMe(&ctx)()
	if newUserID, err = r.ModuleRepo.Insert(ctx, module); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	newModule, err := r.ModuleRepo.FindOne(ctx, newUserID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "modules", newUserID, repository.Insert, nil, newModule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return newUserID, nil
}

// Delete module
func (r *ModuleServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer r.CommitMe(&ctx)()
	oldModule, err := r.ModuleRepo.FindOne(ctx, id)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.HistoryService.RecordHistory(ctx, "modules", id, oldModule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.ModuleRepo.Delete(ctx, id); err != nil {
		r.CancelMe(ctx, err)
		return
	}

	if _, err = r.AuditTrailService.RecordChanges(ctx, "modules", id, repository.Delete, oldModule, nil); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}

// Update module
func (r *ModuleServiceImpl) Update(ctx context.Context, module repository.Module) (err error) {
	defer r.CommitMe(&ctx)()
	oldModule, err := r.ModuleRepo.FindOne(ctx, module.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.ModuleRepo.Update(ctx, module); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	newUser, err := r.ModuleRepo.FindOne(ctx, module.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "modules", module.ID, repository.Update, oldModule, newUser); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}
