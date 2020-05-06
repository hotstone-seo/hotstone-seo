package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
)

// RoleType represented role_type entity
type RoleType struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// RoleTypeRepo to handle role_types entity
// @mock
type RoleTypeRepo interface {
	FindOne(context.Context, int64) (*RoleType, error)
	Find(ctx context.Context, paginationParam PaginationParam) ([]*RoleType, error)
}

// RoleTypeRepoImpl is implementation role_type repository
type RoleTypeRepoImpl struct {
	dig.In
	*sql.DB
}

// NewRoleTypeRepo return new instance of RoleTypeRepo
// @constructor
func NewRoleTypeRepo(impl RoleTypeRepoImpl) RoleTypeRepo {
	return &impl
}

// FindOne role_type
func (r *RoleTypeRepoImpl) FindOne(ctx context.Context, id int64) (e *RoleType, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "name", "updated_at", "created_at").
		From("role_type").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(RoleType)
		if err = rows.Scan(&e.ID, &e.Name, &e.UpdatedAt, &e.CreatedAt); err != nil {
			dbtxn.SetError(ctx, err)
			return nil, err
		}
	}
	return
}

// Find role_type
func (r *RoleTypeRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*RoleType, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "name", "updated_at", "created_at").
		From("role_type").
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*RoleType, 0)
	for rows.Next() {
		var e0 RoleType
		if err = rows.Scan(&e0.ID, &e0.Name, &e0.UpdatedAt, &e0.CreatedAt); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, &e0)
	}
	return
}
