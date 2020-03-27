package repository

import (
	"context"
	"time"
	"github.com/typical-go/typical-rest-server/pkg/dbtype"
)

// AuditTrail Entity
type AuditTrail struct {
	ID         int64
	Time       time.Time
	EntityName string
	EntityID   int64
	Operation  string
	Username   string
	OldData    dbtype.JSON
	NewData    dbtype.JSON
}

// AuditTrailRepo is rule repository [mock]
type AuditTrailRepo interface {
	Insert(ctx context.Context, rule AuditTrail) (lastInsertID int64, err error)
}

// NewAuditTrailRepo return new instance of AuditTrailRepo [constructor]
func NewAuditTrailRepo(impl AuditTrailRepoImpl) AuditTrailRepo {
	return &impl
}
