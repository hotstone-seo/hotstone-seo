package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// User Entity
type User struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email" validate:"required"`
	RoleTypeID *int64    `json:"role_type_id"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

// UserRepo is user repository [mock]
type UserRepo interface {
	FindOne(ctx context.Context, id int64) (*User, error)
	Find(ctx context.Context, paginationParam PaginationParam) ([]*User, error)
	Insert(ctx context.Context, user User) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user User) error
}

// UserRepoImpl is implementation user repository
type UserRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// NewUserRepo return new instance of UserRepo [constructor]
func NewUserRepo(impl UserRepoImpl) UserRepo {
	return &impl
}

// Validate user
func (user User) Validate() error {
	return validator.New().Struct(user)
}

// FindOne user
func (r *UserRepoImpl) FindOne(ctx context.Context, id int64) (*User, error) {
	row := sq.StatementBuilder.
		Select(
			"id",
			"email",
			"role_type_id",
			"updated_at",
			"created_at",
		).
		From("role_user").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r)).
		QueryRowContext(ctx)

	user := new(User)
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.RoleTypeID,
		&user.UpdatedAt,
		&user.CreatedAt,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}

	return user, nil
}

// Find user
func (r *UserRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*User, err error) {
	var (
		rows *sql.Rows
	)

	builder := sq.StatementBuilder.
		Select(
			"id",
			"email",
			"role_type_id",
			"updated_at",
			"created_at",
		).
		From("role_user").
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	builder = ComposePagination(builder, paginationParam)

	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()

	list = make([]*User, 0)

	for rows.Next() {
		user := new(User)
		if err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.RoleTypeID,
			&user.UpdatedAt,
			&user.CreatedAt,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, user)
	}
	return
}

// Insert user
func (r *UserRepoImpl) Insert(ctx context.Context, user User) (lastInsertID int64, err error) {

	query := sq.Insert("role_user").
		Columns(
			"role_type_id",
			"email",
		).
		Values(
			user.RoleTypeID,
			user.Email,
		).
		Suffix("RETURNING \"id\"").
		RunWith(dbtxn.BaseRunner(ctx, r)).
		PlaceholderFormat(sq.Dollar).
		QueryRowContext(ctx)

	if err = query.Scan(&user.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}

	lastInsertID = user.ID
	return
}

// Delete user
func (r *UserRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.StatementBuilder.
		Delete("role_user").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}

// Update user
func (r *UserRepoImpl) Update(ctx context.Context, user User) (err error) {
	builder := sq.StatementBuilder.
		Update("role_user").
		Set("role_type_id", user.RoleTypeID).
		Set("email", user.Email).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": user.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}