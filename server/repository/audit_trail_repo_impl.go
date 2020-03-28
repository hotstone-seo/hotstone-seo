package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// AuditTrailRepoImpl is implementation auditTrail repository
type AuditTrailRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// Insert auditTrail
func (r *AuditTrailRepoImpl) Insert(ctx context.Context, m AuditTrail) (lastInsertID int64, err error) {
	query := sq.Insert("audit_trail").
		Columns("entity_name", "entity_id", "operation", "username", "old_data", "new_data").
		Values(m.EntityName, m.EntityID, m.Operation, m.Username, m.OldData, m.NewData).
		Suffix("RETURNING \"id\"").
		RunWith(dbkit.TxCtx(ctx, r)).
		PlaceholderFormat(sq.Dollar)
	if err = query.QueryRowContext(ctx).Scan(&m.ID); err != nil {
		return
	}
	lastInsertID = m.ID
	return
}
