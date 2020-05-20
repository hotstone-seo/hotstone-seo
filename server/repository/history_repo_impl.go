package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
)

// HistoryRepoImpl is implementation history repository
type HistoryRepoImpl struct {
	dig.In
	*sql.DB
}

// Insert hsitory
func (r *HistoryRepoImpl) Insert(ctx context.Context, m History) (lastInsertID int64, err error) {
	query := sq.Insert("history").
		Columns("entity_id", "entity_from", "username", "data").
		Values(m.EntityID, m.EntityFrom, m.Username, m.Data).
		Suffix("RETURNING \"id\"").
		RunWith(dbtxn.DB(ctx, r)).
		PlaceholderFormat(sq.Dollar)
	if err = query.QueryRowContext(ctx).Scan(&m.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = m.ID
	return
}
