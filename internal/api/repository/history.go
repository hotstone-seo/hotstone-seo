package repository

import (
	"context"
	"time"
)

// History Entity
type History struct {
	ID         int64     `json:"id,omitempty"`
	Time       time.Time `json:"time,omitempty"`
	EntityID   int64     `json:"entity_id,omitempty"`
	EntityFrom string    `json:"entity_from,omitempty"`
	Username   string    `json:"username,omitempty"`
	Data       JSON      `json:"data,omitempty"`
}

// HistoryOperationType is type of changes operation
type HistoryOperationType string

const (
	InsertHistory OperationType = "INSERT"
)

// HistoryRepo is rule repository
// @mock
type HistoryRepo interface {
	Insert(ctx context.Context, history History) (lastInsertID int64, err error)
}

// NewHistoryRepo return new instance of HistoryRepo
// @ctor
func NewHistoryRepo(impl HistoryRepoImpl) HistoryRepo {
	return &impl
}
