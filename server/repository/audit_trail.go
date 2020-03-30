package repository

import (
	"context"
	"time"

	"github.com/typical-go/typical-rest-server/pkg/dbtype"
)

// AuditTrail Entity
type AuditTrail struct {
	ID         int64       `json:"id,omitempty"`
	Time       time.Time   `json:"time,omitempty"`
	EntityName string      `json:"entity_name,omitempty"`
	EntityID   int64       `json:"entity_id,omitempty"`
	Operation  string      `json:"operation,omitempty"`
	Username   string      `json:"username,omitempty"`
	OldData    dbtype.JSON `json:"old_data,omitempty"`
	NewData    dbtype.JSON `json:"new_data,omitempty"`
}

// OperationType is type of changes operation
type OperationType string

const (
	Insert OperationType = "INSERT"
	Update               = "UPDATE"
	Delete               = "DELETE"
)

// AuditTrailRepo is rule repository [mock]
type AuditTrailRepo interface {
	Find(ctx context.Context, paginationParam PaginationParam) ([]*AuditTrail, error)
	Insert(ctx context.Context, auditTrail AuditTrail) (lastInsertID int64, err error)
}

// NewAuditTrailRepo return new instance of AuditTrailRepo [constructor]
func NewAuditTrailRepo(impl AuditTrailRepoImpl) AuditTrailRepo {
	return &impl
}