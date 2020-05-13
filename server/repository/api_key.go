package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// APIKey represented  an api_keys entity
type APIKey struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Key       string    `json:"key" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIKeyRepo to handle api_keys entity
// @mock
type APIKeyRepo interface {
	FindOne(context.Context, int64) (*APIKey, error)
	Find(context.Context, ...dbkit.FindOption) ([]*APIKey, error)
	Insert(context.Context, APIKey) (lastInsertID int64, err error)
	Delete(context.Context, int64) error
}

// APIKeyRepoImpl is implementation API key repository
type APIKeyRepoImpl struct {
	dig.In
	*sql.DB
}

// NewAPIKeyRepo return new instance of APIKeyRepo
// @constructor
func NewAPIKeyRepo(impl APIKeyRepoImpl) APIKeyRepo {
	return &impl
}

// FindOne apiKey
func (r *APIKeyRepoImpl) FindOne(ctx context.Context, id int64) (e *APIKey, err error) {
	row := sq.
		Select(
			"id",
			"name",
			"key",
			"created_at",
			"updated_at",
		).
		From("api_keys").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r)).
		QueryRowContext(ctx)

	e = new(APIKey)
	if err = row.Scan(
		&e.ID,
		&e.Name,
		&e.Key,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}

	return
}

// Find api_keys
func (r *APIKeyRepoImpl) Find(ctx context.Context, opts ...dbkit.FindOption) (list []*APIKey, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"id",
			"name",
			"key",
			"created_at",
			"updated_at",
		).
		From("api_keys").
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	for _, opt := range opts {
		if builder, err = opt.CompileQuery(builder); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
	}

	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()

	list = make([]*APIKey, 0)
	for rows.Next() {
		var e APIKey
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Key,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, &e)
	}
	return
}

// Insert apiKey
func (r *APIKeyRepoImpl) Insert(ctx context.Context, e APIKey) (lastInsertID int64, err error) {

	builder := sq.
		Insert("api_keys").
		Columns(
			"name",
			"key",
		).
		Values(e.Name, e.Key).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&e.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = e.ID
	return
}

// Delete apiKey
func (r *APIKeyRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete("api_keys").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

// Validate apiKey
func (apiKey APIKey) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(TagStructLevelValidation, APIKey{})
	return validate.Struct(apiKey)
}
