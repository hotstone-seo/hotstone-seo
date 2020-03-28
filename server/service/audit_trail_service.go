package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// AuditTrailService contain logic for AuditTrail Controller [mock]
type AuditTrailService interface {
	RecordChanges(ctx context.Context, entityName string, entityID int64, opsType repository.OperationType,
		oldData interface{}, newData interface{}) (lastInsertID int64, err error)
}

// AuditTrailServiceImpl is implementation of AuditTrailService
type AuditTrailServiceImpl struct {
	dig.In
	AuditTrailRepo repository.AuditTrailRepo
}

// NewAuditTrailService return new instance of AuditTrailService [constructor]
func NewAuditTrailService(impl AuditTrailServiceImpl) AuditTrailService {
	return &impl
}

// RecordChanges insert changes
func (r *AuditTrailServiceImpl) RecordChanges(ctx context.Context,
	entityName string, entityID int64, opsType repository.OperationType,
	oldData interface{}, newData interface{}) (lastInsertID int64, err error) {

	auditTrail := repository.AuditTrail{
		EntityName: entityName,
		EntityID:   entityID,
		Operation:  string(opsType),
		Username:   repository.GetUsername(ctx),
		// TODO: set OlData & NewData
	}

	return r.AuditTrailRepo.Insert(ctx, auditTrail)
}
