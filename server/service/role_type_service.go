package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// RoleTypeService contain logic for RoleType Controller
// @mock
type RoleTypeService interface {
	FindOne(ctx context.Context, id int64) (*repository.RoleType, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.RoleType, error)
}

// RoleTypeServiceImpl is implementation of RoleTypeService
type RoleTypeServiceImpl struct {
	dig.In
	RoleTypeRepo repository.RoleTypeRepo
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
