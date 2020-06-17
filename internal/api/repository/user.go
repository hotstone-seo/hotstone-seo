package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

var (
	// UserTableName is table name of user
	UserTableName = "users"
	// UserTable is table column of user
	UserTable = struct {
		ID         string
		Email      string
		UserRoleID string
		UpdatedAt  string
		CreatedAt  string
	}{
		ID:         "id",
		Email:      "email",
		UserRoleID: "user_role_id",
		UpdatedAt:  "updated_at",
		CreatedAt:  "created_at",
	}
)

type (
	// User Entity
	User struct {
		ID         int64     `json:"id"`
		Email      string    `json:"email" validate:"required"`
		UserRoleID int64     `json:"user_role_id"`
		UpdatedAt  time.Time `json:"updated_at"`
		CreatedAt  time.Time `json:"created_at"`
	}
	// UserRepo is user repository
	// @mock
	UserRepo interface {
		FindOne(ctx context.Context, id int64) (*User, error)
		Find(ctx context.Context) ([]*User, error)
		Insert(ctx context.Context, user User) (lastInsertID int64, err error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, user User) error
		FindUserByEmail(ctx context.Context, email string) (*User, error)
	}
	// UserRepoImpl is implementation user repository
	UserRepoImpl struct {
		dig.In
		*sql.DB
	}
)

// NewUserRepo return new instance of UserRepo
// @ctor
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
			UserTable.ID,
			UserTable.Email,
			UserTable.UserRoleID,
			UserTable.UpdatedAt,
			UserTable.CreatedAt,
		).
		From(UserTableName).
		Where(
			sq.Eq{UserTable.ID: id},
		).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r)).
		QueryRowContext(ctx)

	user := new(User)
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.UserRoleID,
		&user.UpdatedAt,
		&user.CreatedAt,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}

	return user, nil
}

// Find user
func (r *UserRepoImpl) Find(ctx context.Context) (list []*User, err error) {
	builder := sq.StatementBuilder.
		Select(
			UserTable.ID,
			UserTable.Email,
			UserTable.UserRoleID,
			UserTable.UpdatedAt,
			UserTable.CreatedAt,
		).
		From(UserTableName).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	rows, err := builder.QueryContext(ctx)
	if err != nil {
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
			&user.UserRoleID,
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
	query := sq.
		Insert(UserTableName).
		Columns(
			UserTable.UserRoleID,
			UserTable.Email,
		).
		Values(
			user.UserRoleID,
			user.Email,
		).
		Suffix(fmt.Sprintf("RETURNING \"%s\"", UserTable.ID)).
		RunWith(dbtxn.DB(ctx, r)).
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
		Delete(UserTableName).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}

// Update user
func (r *UserRepoImpl) Update(ctx context.Context, user User) (err error) {
	builder := sq.StatementBuilder.
		Update(UserTableName).
		Set(UserTable.UserRoleID, user.UserRoleID).
		Set(UserTable.Email, user.Email).
		Set(UserTable.UpdatedAt, time.Now()).
		Where(
			sq.Eq{UserTable.ID: user.ID},
		).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

// FindUserByEmail address
func (r *UserRepoImpl) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	row := sq.StatementBuilder.
		Select(
			UserTable.ID,
			UserTable.UserRoleID,
		).
		From(UserTableName).
		Where(sq.Eq{UserTable.Email: email}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r)).
		QueryRowContext(ctx)

	user := new(User)
	if err := row.Scan(
		&user.ID,
		&user.UserRoleID,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}
	return user, nil
}
