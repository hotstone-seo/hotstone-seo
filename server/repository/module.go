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

// Module Entity
type Module struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Path      string    `json:"path"`
	Pattern   string    `json:"pattern"`
	Label     string    `json:"label"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ModuleRepo is module repository
// @mock
type ModuleRepo interface {
	FindOne(ctx context.Context, id int64) (*Module, error)
	Find(ctx context.Context, paginationParam PaginationParam) ([]*Module, error)
}

// ModuleRepoImpl is implementation module repository
type ModuleRepoImpl struct {
	dig.In
	*sql.DB
}

// NewModuleRepo return new instance of ModuleRepo
// @constructor
func NewModuleRepo(impl ModuleRepoImpl) ModuleRepo {
	return &impl
}

// Validate module
func (module Module) Validate() error {
	return validator.New().Struct(module)
}

// FindOne module
func (r *ModuleRepoImpl) FindOne(ctx context.Context, id int64) (*Module, error) {
	row := sq.StatementBuilder.
		Select(
			"id",
			"name",
			"path",
			"pattern",
			"label",
			"updated_at",
			"created_at",
		).
		From("modules").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r)).
		QueryRowContext(ctx)

	module := new(Module)
	if err := row.Scan(
		&module.ID,
		&module.Name,
		&module.Path,
		&module.Pattern,
		&module.Label,
		&module.UpdatedAt,
		&module.CreatedAt,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}

	return module, nil
}

// Find module
func (r *ModuleRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*Module, err error) {
	var (
		rows *sql.Rows
	)

	builder := sq.StatementBuilder.
		Select(
			"id",
			"name",
			"path",
			"pattern",
			"label",
			"updated_at",
			"created_at",
		).
		From("modules").
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	builder = ComposePagination(builder, paginationParam)

	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()

	list = make([]*Module, 0)

	for rows.Next() {
		module := new(Module)
		if err = rows.Scan(
			&module.ID,
			&module.Name,
			&module.Path,
			&module.Pattern,
			&module.Label,
			&module.UpdatedAt,
			&module.CreatedAt,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, module)
	}
	return
}
