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

var (
	// UserRoleTable is table name for user role
	UserRoleTable = "user_roles"
)

type (
	// UserRole represented user role entity
	UserRole struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Modules   JSONMap   `json:"modules"`
		Menus     Strings   `json:"menus"`
		Paths     Strings   `json:"paths"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}
	// UserRoleRepo to handle role_types entity
	// @mock
	UserRoleRepo interface {
		FindOne(context.Context, int64) (*UserRole, error)
		Find(ctx context.Context, paginationParam PaginationParam) ([]*UserRole, error)
		Insert(ctx context.Context, UserRole UserRole) (lastInsertID int64, err error)
		Update(ctx context.Context, UserRole UserRole) error
		Delete(ctx context.Context, id int64) error
		FindOneByName(context.Context, string) (*UserRole, error)
	}
	// UserRoleRepoImpl is implementation role_type repository
	UserRoleRepoImpl struct {
		dig.In
		*sql.DB
	}
)

// NewUserRoleRepo return new instance of UserRoleRepo
// @ctor
func NewUserRoleRepo(impl UserRoleRepoImpl) UserRoleRepo {
	return &impl
}

// Validate role_type
func (UserRole UserRole) Validate() error {
	return validator.New().Struct(UserRole)
}

// FindOne role_type
func (r *UserRoleRepoImpl) FindOne(ctx context.Context, id int64) (e *UserRole, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"id",
			"name",
			"modules",
			"menus",
			"paths",
			"updated_at",
			"created_at",
		).
		From(UserRoleTable).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.DB(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(UserRole)
		if err = rows.Scan(&e.ID, &e.Name, &e.Modules, &e.Menus, &e.Paths, &e.UpdatedAt, &e.CreatedAt); err != nil {
			dbtxn.SetError(ctx, err)
			return nil, err
		}
	}
	return
}

// Find role_type
func (r *UserRoleRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*UserRole, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"id",
			"name",
			"modules",
			"updated_at",
			"created_at",
		).
		From(UserRoleTable).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*UserRole, 0)
	for rows.Next() {
		var e0 UserRole
		if err = rows.Scan(
			&e0.ID,
			&e0.Name,
			&e0.Modules,
			&e0.UpdatedAt,
			&e0.CreatedAt,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, &e0)
	}
	return
}

// Insert role_type
func (r *UserRoleRepoImpl) Insert(ctx context.Context, e UserRole) (lastInsertID int64, err error) {
	if e.Modules == nil {
		e.Modules = make(map[string]interface{}, 0)
	}
	builder := sq.
		Insert(UserRoleTable).
		Columns(
			"name",
			"modules",
			"menus",
			"paths",
		).
		Values(
			e.Name,
			e.Modules,
			e.Menus,
			e.Paths,
		).
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
func (r *UserRoleRepoImpl) Update(ctx context.Context, e UserRole) (err error) {
	if e.Modules == nil {
		e.Modules = make(map[string]interface{}, 0)
	}
	builder := sq.
		Update(UserRoleTable).
		Set("name", e.Name).
		Set("modules", e.Modules).
		Set("menus", e.Menus).
		Set("paths", e.Paths).
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
func (r *UserRoleRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete(UserRoleTable).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

// FindOneByName role_type
func (r *UserRoleRepoImpl) FindOneByName(ctx context.Context, name string) (e *UserRole, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id").
		From(UserRoleTable).
		Where(sq.Eq{"name": name}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.DB(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(UserRole)
		if err = rows.Scan(&e.ID); err != nil {
			dbtxn.SetError(ctx, err)
			return nil, err
		}
	}
	return
}
