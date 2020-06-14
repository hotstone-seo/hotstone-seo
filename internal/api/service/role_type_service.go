package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// RoleTypeService contain logic for RoleType Controller
// @mock
type RoleTypeService interface {
	FindOne(ctx context.Context, id int64) (*repository.RoleType, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.RoleType, error)
	Insert(ctx context.Context, req RoleTypeRequest) (lastInsertID int64, err error)
	Update(ctx context.Context, req RoleTypeRequest) error
	Delete(ctx context.Context, id int64) error
	FindOneByName(ctx context.Context, name string) (*repository.RoleType, error)
}

// RoleTypeServiceImpl is implementation of RoleTypeService
type RoleTypeServiceImpl struct {
	dig.In
	RoleTypeRepo repository.RoleTypeRepo
	AuditTrailService
	HistoryService
	dbtxn.Transactional
	ModuleRepo repository.ModuleRepo
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
func (r *RoleTypeServiceImpl) Insert(ctx context.Context, req RoleTypeRequest) (newID int64, err error) {
	var data repository.RoleType

	data = repository.RoleType{
		Name: req.Name,
		Modules: map[string]interface{}{
			"modules": mapModules(ctx, req.Modules, r),
		},
		Menus:     strings.Split(req.Menus, "\n"),
		Paths:     strings.Split(req.Paths, "\n"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
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
func (r *RoleTypeServiceImpl) Update(ctx context.Context, req RoleTypeRequest) (err error) {
	var oldData *repository.RoleType
	if oldData, err = r.RoleTypeRepo.FindOne(ctx, req.ID); err != nil {
		return
	}
	var data repository.RoleType
	data = repository.RoleType{
		ID:   req.ID,
		Name: req.Name,
		Modules: map[string]interface{}{
			"modules": mapModules(ctx, req.Modules, r),
		},
		Menus:     strings.Split(req.Menus, "\n"),
		Paths:     strings.Split(req.Paths, "\n"),
		UpdatedAt: time.Now(),
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

// FindOneByName RoleType
func (r *RoleTypeServiceImpl) FindOneByName(ctx context.Context, name string) (roleType *repository.RoleType, err error) {
	return r.RoleTypeRepo.FindOneByName(ctx, name)
}

func mapModules(ctx context.Context, mItem []ModuleItem, r *RoleTypeServiceImpl) []map[string]interface{} {
	faqsMap := make([]map[string]interface{}, len(mItem))
	for index, tempMod := range mItem {
		moduleMs, err := r.ModuleRepo.FindOneByName(ctx, tempMod.Module)
		if err == sql.ErrNoRows {
			log.Error(err)
		}
		/*var APIPathMaps []interface{}
		for k, v := range moduleMs.APIPath {
			switch vv := v.(type) {
			case []interface{}:
				APIPathMaps = vv
				log.Info("index:", k)
				break
			}
		}*/
		faqsMap[index] = map[string]interface{}{
			"path":    moduleMs.Path,
			"name":    tempMod.Module,
			"pattern": moduleMs.Pattern,
			"label":   moduleMs.Label,
		}
	}
	return faqsMap
}

func mapMenus(mItem []string) []map[string]interface{} {
	menusMap := make([]map[string]interface{}, len(mItem))
	for index, temp := range mItem {
		menusMap[index] = map[string]interface{}{
			"menu": temp,
		}
	}
	return menusMap
}

func mapPaths(mItem []string) []map[string]interface{} {
	pathsMap := make([]map[string]interface{}, len(mItem))
	for index, temp := range mItem {
		pathsMap[index] = map[string]interface{}{
			"path": temp,
		}
	}
	return pathsMap
}
