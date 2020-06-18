package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
)

type (
	// AuditTrail Entity
	AuditTrail struct {
		ID         int64     `json:"id,omitempty"`
		Time       time.Time `json:"time,omitempty"`
		EntityName string    `json:"entity_name,omitempty"`
		EntityID   int64     `json:"entity_id,omitempty"`
		Operation  string    `json:"operation,omitempty"`
		Username   string    `json:"username,omitempty"`
		OldData    JSON      `json:"old_data,omitempty"`
		NewData    JSON      `json:"new_data,omitempty"`
	}
	// AuditTrailRepo is rule repository
	// @mock
	AuditTrailRepo interface {
		Find(ctx context.Context, paginationParam PaginationParam) ([]*AuditTrail, error)
		Insert(ctx context.Context, auditTrail AuditTrail) (lastInsertID int64, err error)
	}
	// AuditTrailRepoImpl is implementation auditTrail repository
	AuditTrailRepoImpl struct {
		dig.In
		*sql.DB
	}
)

// NewAuditTrailRepo return new instance of AuditTrailRepo
// @ctor
func NewAuditTrailRepo(impl AuditTrailRepoImpl) AuditTrailRepo {
	return &impl
}

// Find rule
func (r *AuditTrailRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*AuditTrail, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Select(
			"id",
			"time",
			"entity_name",
			"entity_id",
			"operation",
			"username",
			"old_data",
			"new_data",
		).
		From("audit_trail").
		RunWith(r)

	if rows, err = ComposePagination(builder, paginationParam).QueryContext(ctx); err != nil {
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
		Columns(
			"entity_name",
			"entity_id",
			"operation",
			"username",
			"old_data",
			"new_data",
		).
		Values(m.EntityName, m.EntityID, m.Operation, m.Username, m.OldData, m.NewData).
		Suffix("RETURNING \"id\"").
		RunWith(r).
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
	if err := rows.Scan(
		&m.ID,
		&m.Time,
		&m.EntityName,
		&m.EntityID,
		&m.Operation,
		&m.Username,
		&m.OldData,
		&m.NewData,
	); err != nil {
		return nil, err
	}
	return &m, nil
}
