package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbtxn"
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
		Menus     Strings   `json:"menus"`
		Paths     Strings   `json:"paths"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}
	// UserRoleRepo to handle user_roles entity
	// @mock
	UserRoleRepo interface {
		FindOne(context.Context, int64) (*UserRole, error)
		Find(ctx context.Context) ([]*UserRole, error)
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
			"menus",
			"paths",
			"updated_at",
			"created_at",
		).
		From(UserRoleTable).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(UserRole)
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Menus,
			&e.Paths,
			&e.UpdatedAt,
			&e.CreatedAt,
		); err != nil {
			return nil, err
		}
	}
	return
}

// Find role_type
func (r *UserRoleRepoImpl) Find(ctx context.Context) (list []*UserRole, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"id",
			"name",
			"menus",
			"paths",
			"updated_at",
			"created_at",
		).
		From(UserRoleTable).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
	list = make([]*UserRole, 0)
	for rows.Next() {
		var e UserRole
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Menus,
			&e.Paths,
			&e.UpdatedAt,
			&e.CreatedAt,
		); err != nil {
			return
		}
		list = append(list, &e)
	}
	return
}

// Insert role_type
func (r *UserRoleRepoImpl) Insert(ctx context.Context, e UserRole) (lastInsertID int64, err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return -1, err
	}
	builder := sq.
		Insert(UserRoleTable).
		Columns(
			"name",
			"menus",
			"paths",
		).
		Values(
			e.Name,
			e.Menus,
			e.Paths,
		).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	if err = builder.QueryRowContext(ctx).Scan(&e.ID); err != nil {
		txn.SetError(err)
		return
	}
	lastInsertID = e.ID
	return
}

// Update role_type
func (r *UserRoleRepoImpl) Update(ctx context.Context, e UserRole) (err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return err
	}
	builder := sq.
		Update(UserRoleTable).
		Set("name", e.Name).
		Set("menus", e.Menus).
		Set("paths", e.Paths).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": e.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	if _, err = builder.ExecContext(ctx); err != nil {
		txn.SetError(err)
		return
	}
	return
}

// Delete role_type
func (r *UserRoleRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return err
	}
	builder := sq.
		Delete(UserRoleTable).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	if _, err = builder.ExecContext(ctx); err != nil {
		txn.SetError(err)
	}
	return
}

// FindOneByName role_type
func (r *UserRoleRepoImpl) FindOneByName(ctx context.Context, name string) (e *UserRole, err error) {
	builder := sq.
		Select("id").
		From(UserRoleTable).
		Where(sq.Eq{"name": name}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(UserRole)
		if err = rows.Scan(&e.ID); err != nil {
			return nil, err
		}
	}
	return
}
