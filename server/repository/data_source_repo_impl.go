package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/dig"
)

// DataSourceRepoImpl is implementation data_source repository
type DataSourceRepoImpl struct {
	dig.In
	*sql.DB
}

// FindOne data_source
func (r *DataSourceRepoImpl) FindOne(ctx context.Context, id int64) (e *DataSource, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "name", "url", "updated_at", "created_at").
		From("data_sources").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		e = new(DataSource)
		if err = rows.Scan(&e.ID, &e.Name, &e.URL, &e.UpdatedAt, &e.CreatedAt); err != nil {
			dbtxn.SetError(ctx, err)
			return nil, err
		}
	}
	return
}

// Find data_source
func (r *DataSourceRepoImpl) Find(ctx context.Context) (list []*DataSource, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "name", "url", "updated_at", "created_at").
		From("data_sources").
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*DataSource, 0)
	for rows.Next() {
		var e0 DataSource
		if err = rows.Scan(&e0.ID, &e0.Name, &e0.URL, &e0.UpdatedAt, &e0.CreatedAt); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, &e0)
	}
	return
}

// Insert data_source
func (r *DataSourceRepoImpl) Insert(ctx context.Context, e DataSource) (lastInsertID int64, err error) {
	builder := sq.
		Insert("data_sources").
		Columns("name", "url").
		Values(e.Name, e.URL).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if err = builder.QueryRowContext(ctx).Scan(&e.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = e.ID
	return
}

// Delete data_source
func (r *DataSourceRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete("data_sources").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

// Update data_source
func (r *DataSourceRepoImpl) Update(ctx context.Context, e DataSource) (err error) {
	builder := sq.
		Update("data_sources").
		Set("name", e.Name).
		Set("url", e.URL).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": e.ID}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}
