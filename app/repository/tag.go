package repository

import (
	"context"
	"time"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"
)

// Tag represented  tag entity
type Tag struct {
	ID         int64      `json:"id"`
	RuleID     int64      `json:"rule_id"`
	Locale     string     `json:"locale"`
	Type       string     `json:"type"`
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

// NewTagRepo return new instance of TagRepo [autowire]
func NewTagRepo(impl CachedTagRepoImpl) TagRepo {
	return &impl
}
