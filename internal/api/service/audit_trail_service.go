package service

import (
	"context"
	"encoding/json"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type (
	// AuditTrailService contain logic for AuditTrail Controller
	// @mock
	AuditTrailService interface {
		Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error)
		RecordChanges(ctx context.Context, record Record) (lastInsertID int64, err error)
		RecordInsert(ctx context.Context, entity string, id int64, obj interface{})
		RecordDelete(ctx context.Context, entity string, id int64, obj interface{})
		RecordUpdate(ctx context.Context, entity string, id int64, prevObj, nextObj interface{})
	}

	// AuditTrailServiceImpl is implementation of AuditTrailService
	AuditTrailServiceImpl struct {
		dig.In
		AuditTrailRepo repository.AuditTrailRepo
	}

	// Record represents operations that will be logged into audit trail
	Record struct {
		EntityName string
		EntityID   int64
		Operation  OperationType
		PrevData   interface{}
		NextData   interface{}
	}

	// OperationType is type of changes operation
	OperationType string
)

const (
	InsertOp OperationType = "INSERT"
	UpdateOp               = "UPDATE"
	DeleteOp               = "DELETE"
)

// NewAuditTrailService return new instance of AuditTrailService
// @ctor
func NewAuditTrailService(impl AuditTrailServiceImpl) AuditTrailService {
	return &impl
}

// Find audit trail data
func (r *AuditTrailServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error) {
	return r.AuditTrailRepo.Find(ctx, paginationParam)
}

func (r *AuditTrailServiceImpl) RecordChanges(ctx context.Context, record Record) (lastInsertID int64, err error) {
	prevDataJSON := repository.JSON("{}")
	if record.PrevData != nil {
		prevDataJSON, err = json.Marshal(record.PrevData)
		if err != nil {
			return
		}
	}

	nextDataJSON := repository.JSON("{}")
	if record.NextData != nil {
		nextDataJSON, err = json.Marshal(record.NextData)
		if err != nil {
			return
		}
	}

	auditTrail := repository.AuditTrail{
		EntityName: record.EntityName,
		EntityID:   record.EntityID,
		Operation:  string(record.Operation),
		Username:   repository.GetUsername(ctx),
		OldData:    prevDataJSON,
		NewData:    nextDataJSON,
	}

	return r.AuditTrailRepo.Insert(ctx, auditTrail)
}

// RecordInsert to insert audit-trail for insert operation
func (r *AuditTrailServiceImpl) RecordInsert(ctx context.Context, entity string, id int64, obj interface{}) {
	go func() {
		_, err := r.RecordChanges(ctx, Record{
			Operation:  InsertOp,
			EntityName: entity,
			EntityID:   id,
			NextData:   obj,
		})
		if err != nil {
			log.Warnf("record-insert-%s: %s", entity, err.Error())
		}
	}()
}

// RecordDelete to insert audit-trail for delete operation
func (r *AuditTrailServiceImpl) RecordDelete(ctx context.Context, entity string, id int64, obj interface{}) {
	go func() {
		_, err := r.RecordChanges(ctx, Record{
			Operation:  DeleteOp,
			EntityName: entity,
			EntityID:   id,
			PrevData:   obj,
		})
		if err != nil {
			log.Warnf("record-delete-%s: %s", entity, err.Error())
		}
	}()
}

// RecordUpdate to insert audit-trail for update operation
func (r *AuditTrailServiceImpl) RecordUpdate(ctx context.Context, entity string, id int64, oldData, newData interface{}) {
	go func() {
		_, err := r.RecordChanges(ctx, Record{
			Operation:  UpdateOp,
			EntityName: entity,
			EntityID:   id,
			PrevData:   oldData,
			NextData:   newData,
		})
		if err != nil {
			log.Warnf("record-update-%s: %s", entity, err.Error())
		}
	}()
}
