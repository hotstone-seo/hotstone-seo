package repository

import (
	"context"
	"time"
)

// RoleType represented  role_type entity
type RoleType struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// RoleTypeRepo to handle role_type entity
// @mock
type RoleTypeRepo interface {
	FindOne(context.Context, int64) (*RoleType, error)
	Find(ctx context.Context, paginationParam PaginationParam) ([]*RoleType, error)
	// Find(context.Context) ([]*RoleType, error)
}

// NewRoleTypeRepo return new instance of RoleTypeRepo
// @constructor
func NewRoleTypeRepo(impl RoleTypeRepoImpl) RoleTypeRepo {
	return &impl
}
