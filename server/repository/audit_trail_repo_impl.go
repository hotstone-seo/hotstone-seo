package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// AuditTrailRepoImpl is implementation auditTrail repository
type AuditTrailRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// Find rule
func (r *AuditTrailRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*AuditTrail, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "time", "entity_name", "entity_id", "operation", "username", "old_data", "new_data").
		From("audit_trail")
	if rows, err = ComposePagination(builder, paginationParam).RunWith(dbtxn.BaseRunner(ctx, r)).QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*AuditTrail, 0)
	for rows.Next() {
		var rule *AuditTrail
		if rule, err = scanAuditTrail(rows); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, rule)
	}
	return
}

// Insert auditTrail
func (r *AuditTrailRepoImpl) Insert(ctx context.Context, m AuditTrail) (lastInsertID int64, err error) {
	query := sq.Insert("audit_trail").
		Columns("entity_name", "entity_id", "operation", "username", "old_data", "new_data").
		Values(m.EntityName, m.EntityID, m.Operation, m.Username, m.OldData, m.NewData).
		Suffix("RETURNING \"id\"").
		RunWith(dbtxn.BaseRunner(ctx, r)).
		PlaceholderFormat(sq.Dollar)
	if err = query.QueryRowContext(ctx).Scan(&m.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = m.ID
	return
}

func scanAuditTrail(rows *sql.Rows) (*AuditTrail, error) {
	var m AuditTrail
	var err error
	if err = rows.Scan(&m.ID, &m.Time, &m.EntityName, &m.EntityID, &m.Operation, &m.Username, &m.OldData, &m.NewData); err != nil {
		return nil, err
	}
	return &m, nil
}
