package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// StructuredData represents a single ld+json shaped entity of a schema
type StructuredData struct {
	ID        int64     `json:"id"`
	RuleID    int64     `json:"rule_id" validate:"required"`
	Type      string    `json:"type"`
	Data      JSONMap   `json:"data"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// StructuredDataRepo handles database interaction for Structured Data
// @mock
type StructuredDataRepo interface {
	FindOne(context.Context, int64) (*StructuredData, error)
	Find(context.Context, ...dbkit.FindOption) ([]*StructuredData, error)
	Insert(context.Context, StructuredData) (lastInsertID int64, err error)
	Delete(context.Context, int64) error
	Update(context.Context, StructuredData) error

	FindByRule(ctx context.Context, ruleID int64) ([]*StructuredData, error)
}

// StructuredDataRepoImpl is an implementation of StructuredDataRepo
type StructuredDataRepoImpl struct {
	dig.In
	*sql.DB
}

// NewStructuredDataRepo returns new instance of StructuredDataRepo
// @constructor
func NewStructuredDataRepo(impl StructuredDataRepoImpl) StructuredDataRepo {
	return &impl
}

// FindOne select a single Structured Data by its ID
func (r *StructuredDataRepoImpl) FindOne(ctx context.Context, id int64) (e *StructuredData, err error) {
	row := sq.
		Select(
			"id",
			"rule_id",
			"type",
			"data",
			"updated_at",
			"created_at",
		).
		From("structured_datas").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r)).
		QueryRowContext(ctx)

	e = new(StructuredData)
	if err = row.Scan(
		&e.ID,
		&e.RuleID,
		&e.Type,
		&e.Data,
		&e.UpdatedAt,
		&e.CreatedAt,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}
	return
}

// Find select a list of Structured data by filtering options
func (r *StructuredDataRepoImpl) Find(ctx context.Context, opts ...dbkit.FindOption) (list []*StructuredData, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"id",
			"rule_id",
			"type",
			"data",
			"updated_at",
			"created_at",
		).
		From("structured_datas").
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

	list = make([]*StructuredData, 0)
	for rows.Next() {
		var e StructuredData
		if err = rows.Scan(
			&e.ID,
			&e.RuleID,
			&e.Type,
			&e.Data,
			&e.UpdatedAt,
			&e.CreatedAt,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, &e)
	}
	return
}

// Insert creates a new Structured Data in database, returning its ID if success
func (r *StructuredDataRepoImpl) Insert(ctx context.Context, e StructuredData) (lastInsertID int64, err error) {
	if e.Data == nil {
		e.Data = make(map[string]interface{}, 0)
	}
	builder := sq.
		Insert("structured_datas").
		Columns(
			"rule_id",
			"type",
			"data",
		).
		Values(e.RuleID, e.Type, e.Data).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&e.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = e.ID
	return
}

// Delete removes a Structured Data entry from database
func (r *StructuredDataRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete("structured_datas").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

// Update modifies existing Structured Data fields in the database
func (r *StructuredDataRepoImpl) Update(ctx context.Context, e StructuredData) (err error) {
	if e.Data == nil {
		e.Data = make(map[string]interface{}, 0)
	}

	builder := sq.
		Update("structured_datas").
		Set("rule_id", e.RuleID).
		Set("type", e.Type).
		Set("data", e.Data).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": e.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}

// FindByRule returns list of structured data based on rule ID
func (r *StructuredDataRepoImpl) FindByRule(ctx context.Context, ruleID int64) ([]*StructuredData, error) {
	return r.Find(ctx, dbkit.Equal("rule_id", strconv.FormatInt(ruleID, 10)))
}

// TODO: Create implementation for validating structured data
func (s StructuredData) Validate() error {
	return nil
}

// ToTag returns Tag representation of a Structured Data
func (s StructuredData) ToTag() (tag Tag) {
	tag = Tag{
		RuleID: s.RuleID,
		Type:   "script",
		Attributes: Attrs{
			"type": "application/ld+json",
		},
		UpdatedAt: s.UpdatedAt,
		CreatedAt: s.CreatedAt,
	}
	rawData, err := json.Marshal(s.Data)
	if err != nil {
		log.Error(err)
		return
	}
	tag.Value = string(rawData)
	return
}
