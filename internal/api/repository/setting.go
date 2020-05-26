package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

var (
	// SettingTable is table name for setting entity
	SettingTable = "settings"
	// SettingCols is columns of setting
	SettingCols = struct {
		Key       string
		Value     string
		UpdatedAt string
	}{
		Key:       "key",
		Value:     "value",
		UpdatedAt: "updated_at",
	}
)

type (
	// Setting entity
	Setting struct {
		Key       string    `json:"key"`
		Value     string    `json:"value" validate:"required"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// SettingRepo to get Setting entity
	// @mock
	SettingRepo interface {
		Find(context.Context, ...dbkit.SelectOption) ([]*Setting, error)
		Update(context.Context, *Setting, dbkit.UpdateOption) error
	}

	// SettingRepoImpl is implementation SettingRepo
	SettingRepoImpl struct {
		dig.In
		*sql.DB
	}
)

// NewSettingRepo return new instance of SettingRepo
// @ctor
func NewSettingRepo(impl SettingRepoImpl) SettingRepo {
	return &impl
}

// Find setting
func (s *SettingRepoImpl) Find(ctx context.Context, opts ...dbkit.SelectOption) ([]*Setting, error) {
	var err error

	builder := sq.StatementBuilder.
		Select(
			SettingCols.Key,
			SettingCols.Value,
			SettingCols.UpdatedAt,
		).
		From(SettingTable).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, s))

	for _, opt := range opts {
		if builder, err = opt.CompileSelect(builder); err != nil {
			return nil, err
		}
	}

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}
	defer rows.Close()

	settings := make([]*Setting, 0)

	for rows.Next() {
		var setting Setting
		if err = rows.Scan(
			&setting.Key,
			&setting.Value,
			&setting.UpdatedAt,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return nil, err
		}
		settings = append(settings, &setting)
	}

	return settings, nil
}

// Update setting
func (s *SettingRepoImpl) Update(ctx context.Context, setting *Setting, opt dbkit.UpdateOption) (err error) {
	builder := sq.
		Update(SettingTable).
		Set(SettingCols.Value, setting.Value).
		Set(SettingCols.UpdatedAt, time.Now()).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, s))

	if builder, err = opt.CompileUpdate(builder); err != nil {
		return
	}

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}

	return
}
