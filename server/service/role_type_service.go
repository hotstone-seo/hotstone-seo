package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// RoleTypeService contain logic for RoleType Controller
// @mock
type RoleTypeService interface {
	FindOne(ctx context.Context, id int64) (*repository.RoleType, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.RoleType, error)
	Insert(ctx context.Context, roleType repository.RoleType) (lastInsertID int64, err error)
	Update(ctx context.Context, roleType repository.RoleType) error
	Delete(ctx context.Context, id int64) error
}

// RoleTypeServiceImpl is implementation of RoleTypeService
type RoleTypeServiceImpl struct {
	dig.In
	RoleTypeRepo repository.RoleTypeRepo
	AuditTrailService
	HistoryService
	dbtxn.Transactional
}

// NewRoleTypeService return new instance of RoleTypeService
// @constructor
func NewRoleTypeService(impl RoleTypeServiceImpl) RoleTypeService {
	return &impl
}

// FindOne RoleType
func (r *RoleTypeServiceImpl) FindOne(ctx context.Context, id int64) (roleType *repository.RoleType, err error) {
	return r.RoleTypeRepo.FindOne(ctx, id)
}

// Find RoleType
func (r *RoleTypeServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) (list []*repository.RoleType, err error) {
	return r.RoleTypeRepo.Find(ctx, paginationParam)
}

// Insert RoleType
func (r *RoleTypeServiceImpl) Insert(ctx context.Context, data repository.RoleType) (newID int64, err error) {
	if data.ID, err = r.RoleTypeRepo.Insert(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"roleType",
			data.ID,
			repository.Insert,
			nil,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return data.ID, nil
}

// Update RoleType
func (r *RoleTypeServiceImpl) Update(ctx context.Context, data repository.RoleType) (err error) {
	var oldData *repository.RoleType
	if oldData, err = r.RoleTypeRepo.FindOne(ctx, data.ID); err != nil {
		return
	}
	if err = r.RoleTypeRepo.Update(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"roleType",
			data.ID,
			repository.Update,
			oldData,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

// Delete RoleType
func (r *RoleTypeServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var oldData *repository.RoleType
	if oldData, err = r.RoleTypeRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = r.RoleTypeRepo.Delete(ctx, id); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	go func() {
		if _, histErr := r.HistoryService.RecordHistory(
			ctx,
			"roleType",
			id,
			oldData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"roleType",
			id,
			repository.Delete,
			oldData,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}
