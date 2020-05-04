package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// DataSourceService contain logic for DataSourceController
// @mock
type DataSourceService interface {
	repository.DataSourceRepo
}

// DataSourceServiceImpl is implementation of DataSourceService
type DataSourceServiceImpl struct {
	dig.In
	dbtxn.Transactional
	repository.DataSourceRepo
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewDataSourceService return new instance of DataSourceService
// @constructor
func NewDataSourceService(impl DataSourceServiceImpl) DataSourceService {
	return &impl
}

// Insert data source
func (s *DataSourceServiceImpl) Insert(ctx context.Context, data repository.DataSource) (newID int64, err error) {
	if data.ID, err = s.DataSourceRepo.Insert(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"data_sources",
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

// Update data source
func (s *DataSourceServiceImpl) Update(ctx context.Context, data repository.DataSource) (err error) {
	var oldData *repository.DataSource
	if oldData, err = s.DataSourceRepo.FindOne(ctx, data.ID); err != nil {
		return
	}
	if err = s.DataSourceRepo.Update(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"data_sources",
			data.ID,
			repository.Update,
			oldData,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

// Delete data source
func (s *DataSourceServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var oldData *repository.DataSource
	if oldData, err = s.DataSourceRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = s.DataSourceRepo.Delete(ctx, id); err != nil {
		return
	}
	go func() {
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			"data_source",
			id,
			oldData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"data_sources",
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
