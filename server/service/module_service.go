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
}

// ModuleServiceImpl is implementation of ModuleService
type ModuleServiceImpl struct {
	dig.In
	ModuleRepo repository.ModuleRepo
	dbtxn.Transactional
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
