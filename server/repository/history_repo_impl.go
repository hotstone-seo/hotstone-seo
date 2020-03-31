package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// HistoryRepoImpl is implementation history repository
type HistoryRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// Insert hsitory
func (r *HistoryRepoImpl) Insert(ctx context.Context, m History) (lastInsertID int64, err error) {
	query := sq.Insert("history").
		Columns("entity_id", "entity_from", "username", "data").
		Values(m.EntityID, m.EntityFrom, m.Username, m.Data).
		Suffix("RETURNING \"id\"").
		RunWith(dbkit.TxCtx(ctx, r)).
		PlaceholderFormat(sq.Dollar)
	if err = query.QueryRowContext(ctx).Scan(&m.ID); err != nil {
		return
	}
	lastInsertID = m.ID
	return
}
