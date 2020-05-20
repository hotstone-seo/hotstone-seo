package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// RoleType represented role_type entity
type RoleType struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Modules   JSONMap   `json:"modules"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// RoleTypeRepo to handle role_types entity
// @mock
type RoleTypeRepo interface {
	FindOne(context.Context, int64) (*RoleType, error)
	Find(ctx context.Context, paginationParam PaginationParam) ([]*RoleType, error)
	Insert(ctx context.Context, roleType RoleType) (lastInsertID int64, err error)
	Update(ctx context.Context, roleType RoleType) error
	Delete(ctx context.Context, id int64) error
	FindOneByName(context.Context, string) (*RoleType, error)
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

// Validate role_type
func (roleType RoleType) Validate() error {
	return validator.New().Struct(roleType)
}

// FindOne role_type
func (r *RoleTypeRepoImpl) FindOne(ctx context.Context, id int64) (e *RoleType, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "name", "modules", "updated_at", "created_at").
		From("role_type").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.DB(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(RoleType)
		if err = rows.Scan(&e.ID, &e.Name, &e.Modules, &e.UpdatedAt, &e.CreatedAt); err != nil {
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
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.DB(ctx, r))
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

// Insert role_type
func (r *RoleTypeRepoImpl) Insert(ctx context.Context, e RoleType) (lastInsertID int64, err error) {
	if e.Modules == nil {
		e.Modules = make(map[string]interface{}, 0)
	}

	builder := sq.
		Insert("role_type").
		Columns(
			"name",
			"modules",
		).
		Values(e.Name, e.Modules).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&e.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = e.ID
	return
}

// Update role_type
func (r *RoleTypeRepoImpl) Update(ctx context.Context, e RoleType) (err error) {
	if e.Modules == nil {
		e.Modules = make(map[string]interface{}, 0)
	}

	builder := sq.
		Update("role_type").
		Set("name", e.Name).
		Set("modules", e.Modules).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": e.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}

// Delete role_type
func (r *RoleTypeRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete("role_type").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

// FindOneByName role_type
func (r *RoleTypeRepoImpl) FindOneByName(ctx context.Context, name string) (e *RoleType, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id").
		From("role_type").
		Where(sq.Eq{"name": name}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.DB(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(RoleType)
		if err = rows.Scan(&e.ID); err != nil {
			dbtxn.SetError(ctx, err)
			return nil, err
		}
	}
	return
}
