package service

import (
	"context"
	"encoding/json"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtype"
	"go.uber.org/dig"
)

// AuditTrailService contain logic for AuditTrail Controller [mock]
type AuditTrailService interface {
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error)
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

// Find audit trail data
func (r *AuditTrailServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error) {
	return r.AuditTrailRepo.Find(ctx, paginationParam)
}

// RecordChanges insert changes
func (r *AuditTrailServiceImpl) RecordChanges(ctx context.Context,
	entityName string, entityID int64, opsType repository.OperationType,
	oldData interface{}, newData interface{}) (lastInsertID int64, err error) {

	oldDataJSON := dbtype.JSON("{}")
	if oldData != nil {
		oldDataJSON, err = json.Marshal(oldData)
		if err != nil {
			return
		}
	}

	newDataJSON := dbtype.JSON("{}")
	if newData != nil {
		newDataJSON, err = json.Marshal(newData)
		if err != nil {
			return
		}
	}

	auditTrail := repository.AuditTrail{
		EntityName: entityName,
		EntityID:   entityID,
		Operation:  string(opsType),
		Username:   repository.GetUsername(ctx),
		OldData:    oldDataJSON,
		NewData:    newDataJSON,
	}

	return r.AuditTrailRepo.Insert(ctx, auditTrail)
}
