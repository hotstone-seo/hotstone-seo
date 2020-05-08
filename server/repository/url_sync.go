package repository

import (
	"context"
	"time"
)

// URLSync is database entity of sync information between server and url-store
type URLSync struct {
	Version          int64     `json:"version"`
	Operation        string    `json:"operation" validate:"required"`
	RuleID           int64     `json:"rule_id"`
	LatestURLPattern *string   `json:"latest_url_pattern"`
	CreatedAt        time.Time `json:"-"`
}

// URLSyncRepo is repository of URLSync entity
// @mock
type URLSyncRepo interface {
	FindOne(ctx context.Context, id int64) (*URLSync, error)
	Find(ctx context.Context) ([]*URLSync, error)
	Insert(ctx context.Context, URLSync URLSync) (lastInsertID int64, err error)
	GetLatestVersion(ctx context.Context) (latestVersion int64, err error)
	GetListDiff(ctx context.Context, offsetVersion int64) ([]*URLSync, error)

	FindRule(ctx context.Context, ruleID int64) (*URLSync, error)
}

// NewURLSyncRepo return new instance of URLSyncRepo
// @constructor
func NewURLSyncRepo(impl URLSyncRepoImpl) URLSyncRepo {
	return &impl
}
