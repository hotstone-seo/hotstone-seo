package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// APIKeyService contain logic for APIKeyController
// @mock
type APIKeyService interface {
	repository.APIKeyRepo
}

// APIKeyServiceImpl is implementation of APIKeyService
type APIKeyServiceImpl struct {
	dig.In
	repository.APIKeyRepo
	dbtxn.Transactional
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewAPIKeyService return new instance of APIKeyService
// @constructor
func NewAPIKeyService(impl APIKeyServiceImpl) APIKeyService {
	return &impl
}

// Insert tag
func (s *APIKeyServiceImpl) Insert(ctx context.Context, data repository.APIKey) (newID int64, err error) {
	if data.ID, err = s.APIKeyRepo.Insert(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"api_keys",
			data.ID,
			repository.Insert,
			nil,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return data.ID, nil
}

// Delete tag
func (s *APIKeyServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var oldData *repository.APIKey
	if oldData, err = s.APIKeyRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = s.APIKeyRepo.Delete(ctx, id); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	go func() {
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			"api_keys",
			id,
			oldData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"api_keys",
			id,
			repository.Delete,
			oldData,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}
