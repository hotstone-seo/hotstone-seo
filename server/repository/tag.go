package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// Tag represented  tag entity
type Tag struct {
	ID         int64     `json:"id"`
	RuleID     int64     `json:"rule_id" validate:"required"`
	Locale     string    `json:"locale" validate:"required"`
	Type       string    `json:"type" validate:"required"`
	Attributes Attrs     `json:"attributes"`
	Value      string    `json:"value"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

// TagRepo to handle tags entity
// @mock
type TagRepo interface {
	FindOne(context.Context, int64) (*Tag, error)
	Find(context.Context, ...dbkit.FindOption) ([]*Tag, error)
	Insert(context.Context, Tag) (lastInsertID int64, err error)
	Delete(context.Context, int64) error
	Update(context.Context, Tag) error

	FindByRuleAndLocale(ctx context.Context, ruleID int64, locale string) ([]*Tag, error)
}

// TagRepoImpl is implementation tag repository
type TagRepoImpl struct {
	dig.In
	*sql.DB
}

// NewTagRepo return new instance of TagRepo
// @constructor
func NewTagRepo(impl TagRepoImpl) TagRepo {
	return &impl
}

// FindOne tag
func (r *TagRepoImpl) FindOne(ctx context.Context, id int64) (e *Tag, err error) {
	row := sq.
		Select(
			"id",
			"rule_id",
			"locale",
			"type",
			"attributes",
			"value",
			"updated_at",
			"created_at",
		).
		From("tags").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r)).
		QueryRowContext(ctx)

	e = new(Tag)
	if err = row.Scan(
		&e.ID,
		&e.RuleID,
		&e.Locale,
		&e.Type,
		&e.Attributes,
		&e.Value,
		&e.UpdatedAt,
		&e.CreatedAt,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}

	return
}

// Find tags
func (r *TagRepoImpl) Find(ctx context.Context, opts ...dbkit.FindOption) (list []*Tag, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"id",
			"rule_id",
			"locale",
			"type",
			"attributes",
			"value",
			"updated_at",
			"created_at",
		).
		From("tags").
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

	list = make([]*Tag, 0)
	for rows.Next() {
		var e Tag
		if err = rows.Scan(
			&e.ID,
			&e.RuleID,
			&e.Locale,
			&e.Type,
			&e.Attributes,
			&e.Value,
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

// FindByRuleAndLocale return list of tag baed on ruleID and locale
func (r *TagRepoImpl) FindByRuleAndLocale(ctx context.Context, ruleID int64, locale string) (list []*Tag, err error) {
	return r.Find(ctx,
		dbkit.Equal("rule_id", strconv.FormatInt(ruleID, 10)),
		dbkit.Equal("locale", locale),
	)
}

// Insert tag
func (r *TagRepoImpl) Insert(ctx context.Context, e Tag) (lastInsertID int64, err error) {
	if e.Attributes == nil {
		e.Attributes = map[string]string{}
	}

	builder := sq.
		Insert("tags").
		Columns(
			"rule_id",
			"locale",
			"type",
			"attributes",
			"value",
		).
		Values(e.RuleID, e.Locale, e.Type, e.Attributes, e.Value).
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

// Delete tag
func (r *TagRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete("tags").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

// Update tag
func (r *TagRepoImpl) Update(ctx context.Context, e Tag) (err error) {
	if e.Attributes == nil {
		e.Attributes = map[string]string{}
	}

	builder := sq.
		Update("tags").
		Set("rule_id", e.RuleID).
		Set("locale", e.Locale).
		Set("type", e.Type).
		Set("attributes", e.Attributes).
		Set("value", e.Value).
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

// Validate tag
func (tag Tag) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(TagStructLevelValidation, Tag{})
	return validate.Struct(tag)
}

// TagStructLevelValidation validate per type
func TagStructLevelValidation(sl validator.StructLevel) {
	tag := sl.Current().Interface().(Tag)
	var validate func(validator.StructLevel, Tag)
	switch tag.Type {
	case "title":
		validate = validateTitleTag
	case "meta":
		validate = validateMetaTag
	case "link":
		validate = validateCanonicalTag
	case "script":
		validate = validateScriptTag
	default:
		return
	}

	validate(sl, tag)
}

func validateTitleTag(sl validator.StructLevel, tag Tag) {
	if tag.Value == "" {
		sl.ReportError(tag.Value, "Value", "Value", "noempty", "")
	}
}

func validateMetaTag(sl validator.StructLevel, tag Tag) {
	if !validAttributesKey(tag, "name", "content") {
		sl.ReportError(tag.Attributes, "Attributes", "Attributes", "", "")
	}
	if tag.Value != "" {
		sl.ReportError(tag.Value, "Value", "Value", "", "")
	}
}

func validateCanonicalTag(sl validator.StructLevel, tag Tag) {
	if !validAttributesKey(tag, "rel", "href") {
		sl.ReportError(tag.Attributes, "Attributes", "Attributes", "", "")
	}
}

func validateScriptTag(sl validator.StructLevel, tag Tag) {
	if !validAttributesKey(tag, "src") {
		sl.ReportError(tag.Attributes, "Attributes", "Attributes", "", "")
	}
}

func validAttributesKey(tag Tag, keys ...string) bool {
	for _, key := range keys {
		if tag.Attributes[key] == "" {
			return false
		}
	}
	return true
}
