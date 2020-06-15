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

type (
	// UserRoleSvc contain logic for UserRole Controller
	// @mock
	UserRoleSvc interface {
		FindOne(ctx context.Context, id int64) (*repository.UserRole, error)
		Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.UserRole, error)
		Insert(ctx context.Context, req UserRoleRequest) (lastInsertID int64, err error)
		Update(ctx context.Context, req UserRoleRequest) error
		Delete(ctx context.Context, id int64) error
		FindOneByName(ctx context.Context, name string) (*repository.UserRole, error)
	}
	// UserRoleSvcImpl is implementation of UserRoleService
	UserRoleSvcImpl struct {
		dig.In
		UserRoleRepo repository.UserRoleRepo
		AuditTrailService
		HistoryService
		dbtxn.Transactional
		ModuleRepo repository.ModuleRepo
	}
	// UserRoleRequest is request model for UserRole related method
	UserRoleRequest struct {
		ID      int64        `json:"id"`
		Name    string       `json:"name"`
		Menus   string       `json:"menus"`
		Paths   string       `json:"paths"`
		Modules []ModuleItem `json:"modules"`
	}
	// ModuleItem contain module TODO: remove this
	ModuleItem struct {
		Module string `json:"name"`
	}
)

// NewUserRoleSvc return new instance of UserRoleService
// @ctor
func NewUserRoleSvc(impl UserRoleSvcImpl) UserRoleSvc {
	return &impl
}

// FindOne UserRole
func (r *UserRoleSvcImpl) FindOne(ctx context.Context, id int64) (UserRole *repository.UserRole, err error) {
	return r.UserRoleRepo.FindOne(ctx, id)
}

// Find UserRole
func (r *UserRoleSvcImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) (list []*repository.UserRole, err error) {
	return r.UserRoleRepo.Find(ctx, paginationParam)
}

// Insert UserRole
func (r *UserRoleSvcImpl) Insert(ctx context.Context, req UserRoleRequest) (newID int64, err error) {
	var data repository.UserRole

	data = repository.UserRole{
		Name: req.Name,
		Modules: map[string]interface{}{
			"modules": mapModules(ctx, req.Modules, r),
		},
		Menus:     strings.Split(req.Menus, "\n"),
		Paths:     strings.Split(req.Paths, "\n"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if data.ID, err = r.UserRoleRepo.Insert(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"UserRole",
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

// Update UserRole
func (r *UserRoleSvcImpl) Update(ctx context.Context, req UserRoleRequest) (err error) {
	var oldData *repository.UserRole
	if oldData, err = r.UserRoleRepo.FindOne(ctx, req.ID); err != nil {
		return
	}
	var data repository.UserRole
	data = repository.UserRole{
		ID:   req.ID,
		Name: req.Name,
		Modules: map[string]interface{}{
			"modules": mapModules(ctx, req.Modules, r),
		},
		Menus:     strings.Split(req.Menus, "\n"),
		Paths:     strings.Split(req.Paths, "\n"),
		UpdatedAt: time.Now(),
	}
	if err = r.UserRoleRepo.Update(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"UserRole",
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

// Delete UserRole
func (r *UserRoleSvcImpl) Delete(ctx context.Context, id int64) (err error) {
	var oldData *repository.UserRole
	if oldData, err = r.UserRoleRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = r.UserRoleRepo.Delete(ctx, id); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	go func() {
		if _, histErr := r.HistoryService.RecordHistory(
			ctx,
			"UserRole",
			id,
			oldData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := r.AuditTrailService.RecordChanges(
			ctx,
			"UserRole",
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

// FindOneByName UserRole
func (r *UserRoleSvcImpl) FindOneByName(ctx context.Context, name string) (UserRole *repository.UserRole, err error) {
	return r.UserRoleRepo.FindOneByName(ctx, name)
}

func mapModules(ctx context.Context, mItem []ModuleItem, r *UserRoleSvcImpl) []map[string]interface{} {
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
