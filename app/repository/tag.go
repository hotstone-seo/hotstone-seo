package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"gopkg.in/go-playground/validator.v9"
)

// Tag represented  tag entity
type Tag struct {
	ID         int64      `json:"id"`
	RuleID     int64      `json:"rule_id" validate:"required"`
	Locale     string     `json:"locale" validate:"required"`
	Type       string     `json:"type" validate:"required"`
	Attributes dbkit.JSON `json:"attributes"`
	Value      string     `json:"value"`
	UpdatedAt  time.Time  `json:"updated_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

type TagFilter struct {
	RuleID int64  `json:"rule_id"`
	Locale string `json:"locale"`
}

// TagRepo to handle tags entity [mock]
type TagRepo interface {
	FindOne(context.Context, int64) (*Tag, error)
	Find(context.Context, TagFilter) ([]*Tag, error)
	Insert(context.Context, Tag) (lastInsertID int64, err error)
	Delete(context.Context, int64) error
	Update(context.Context, Tag) error
}

// NewTagRepo return new instance of TagRepo [constructor]
func NewTagRepo(impl TagRepoImpl) TagRepo {
	return &impl
}

func (tag Tag) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(TagStructLevelValidation, Tag{})
	return validate.Struct(tag)
}

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
	var attributes map[string]string
	if err := json.Unmarshal(tag.Attributes, &attributes); err != nil {
		return false
	}
	for _, key := range keys {
		if attributes[key] == "" {
			return false
		}
	}
	return true
}
