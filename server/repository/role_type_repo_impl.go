package repository

import (
	"context"
	"database/sql"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/dig"
)

// RoleTypeRepoImpl is implementation data_source repository
type RoleTypeRepoImpl struct {
	dig.In
	*sql.DB
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
