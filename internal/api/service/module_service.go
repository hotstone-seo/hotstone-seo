package service

import (
	"context"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
)

// ModuleService contain logic for Module Controller
// @mock
type ModuleService interface {
	FindOne(ctx context.Context, id int64) (*repository.Module, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.Module, error)
	Insert(ctx context.Context, req ModuleRequest) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, req ModuleRequest) error
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
// @ctor
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
//func (r *ModuleServiceImpl) Insert(ctx context.Context, module repository.Module) (newUserID int64, err error) {
func (r *ModuleServiceImpl) Insert(ctx context.Context, req ModuleRequest) (newID int64, err error) {
	var data repository.Module
	data = repository.Module{
		Name:      req.Name,
		Path:      req.Path,
		Label:     req.Label,
		Pattern:   req.Pattern,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	defer r.BeginTxn(&ctx)()
	if data.ID, err = r.ModuleRepo.Insert(ctx, data); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	newModule, err := r.ModuleRepo.FindOne(ctx, data.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "modules", data.ID, repository.Insert, nil, newModule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return data.ID, nil
}

// Delete module
func (r *ModuleServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer r.BeginTxn(&ctx)()
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
func (r *ModuleServiceImpl) Update(ctx context.Context, req ModuleRequest) (err error) {
	defer r.BeginTxn(&ctx)()
	oldModule, err := r.ModuleRepo.FindOne(ctx, req.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	var data repository.Module
	data = repository.Module{
		ID:        req.ID,
		Name:      req.Name,
		Path:      req.Path,
		Label:     req.Label,
		Pattern:   req.Pattern,
		UpdatedAt: time.Now(),
	}

	if err = r.ModuleRepo.Update(ctx, data); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	newModule, err := r.ModuleRepo.FindOne(ctx, req.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "modules", req.ID, repository.Update, oldModule, newModule); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}
