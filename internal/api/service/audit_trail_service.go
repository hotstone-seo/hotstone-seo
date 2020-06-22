package service

import (
	"context"
	"encoding/json"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type (
	// AuditTrailSvc contain logic for AuditTrail Controller
	// @mock
	AuditTrailSvc interface {
		Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error)
		RecordInsert(ctx context.Context, entity string, id int64, obj interface{})
		RecordDelete(ctx context.Context, entity string, id int64, obj interface{})
		RecordUpdate(ctx context.Context, entity string, id int64, prevObj, nextObj interface{})
	}
	// AuditTrailSvcImpl is implementation of AuditTrailSvc
	AuditTrailSvcImpl struct {
		dig.In
		AuditTrailRepo repository.AuditTrailRepo
	}
)

// NewAuditTrailSvc return new instance of AuditTrailSvc
// @ctor
func NewAuditTrailSvc(impl AuditTrailSvcImpl) AuditTrailSvc {
	return &impl
}

// Find audit trail data
func (r *AuditTrailSvcImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error) {
	return r.AuditTrailRepo.Find(ctx, paginationParam)
}

func marshal(v interface{}) []byte {
	if v != nil {
		b, err := json.Marshal(v)
		if err == nil {
			return b
		}
	}

	return []byte("{}")
}

// RecordInsert to insert audit-trail for insert operation
func (r *AuditTrailSvcImpl) RecordInsert(ctx context.Context, entity string, id int64, obj interface{}) {
	go func() {
		_, err := r.AuditTrailRepo.Insert(ctx, repository.AuditTrail{
			EntityName: entity,
			EntityID:   id,
			Operation:  "INSERT",
			Username:   GetUsername(ctx),
			OldData:    []byte("{}"),
			NewData:    marshal(obj),
		})
		if err != nil {
			log.Warnf("record-insert-%s: %s", entity, err.Error())
		}
	}()
}

// RecordDelete to insert audit-trail for delete operation
func (r *AuditTrailSvcImpl) RecordDelete(ctx context.Context, entity string, id int64, obj interface{}) {
	go func() {
		_, err := r.AuditTrailRepo.Insert(ctx, repository.AuditTrail{
			EntityName: entity,
			EntityID:   id,
			Operation:  "DELETE",
			Username:   GetUsername(ctx),
			OldData:    marshal(obj),
			NewData:    []byte("{}"),
		})
		if err != nil {
			log.Warnf("record-delete-%s: %s", entity, err.Error())
		}
	}()
}

// RecordUpdate to insert audit-trail for update operation
func (r *AuditTrailSvcImpl) RecordUpdate(ctx context.Context, entity string, id int64, oldObj, newObj interface{}) {
	go func() {
		_, err := r.AuditTrailRepo.Insert(ctx, repository.AuditTrail{
			EntityName: entity,
			EntityID:   id,
			Operation:  "UPDATE",
			Username:   GetUsername(ctx),
			OldData:    marshal(oldObj),
			NewData:    marshal(newObj),
		})
		if err != nil {
			log.Warnf("record-update-%s: %s", entity, err.Error())
		}
	}()
}
