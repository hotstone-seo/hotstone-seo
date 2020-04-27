package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// StructuredDataService manages instances of Structured Data
type StructuredDataService interface {
	repository.StructuredDataRepo
}

// StructuredDataServiceImpl is an impolementation of StructuredDataService
type StructuredDataServiceImpl struct {
	dig.In
	repository.StructuredDataRepo
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewStructuredDataService returns nrw instance of StructuredDataService [constructor]
func NewStructuredDataService(impl StructuredDataServiceImpl) StructuredDataService {
	return &impl
}

func (s *StructuredDataServiceImpl) Insert(ctx context.Context, data repository.StructuredData) (newID int64, err error) {
	if data.ID, err = s.StructuredDataRepo.Insert(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"structured data",
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

func (s *StructuredDataServiceImpl) Update(ctx context.Context, data repository.StructuredData) (err error) {
	var currentData *repository.StructuredData
	if currentData, err = s.StructuredDataRepo.FindOne(ctx, data.ID); err != nil {
		return
	}
	if err = s.StructuredDataRepo.Update(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"structured data",
			data.ID,
			repository.Update,
			currentData,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

func (s *StructuredDataServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var currentData *repository.StructuredData
	if currentData, err = s.StructuredDataRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = s.StructuredDataRepo.Delete(ctx, id); err != nil {
		return
	}
	go func() {
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			"structured data",
			id,
			currentData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"structured data",
			id,
			repository.Delete,
			currentData,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}
