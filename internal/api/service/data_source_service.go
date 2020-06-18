package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
)

type (
	// DataSourceService contain logic for DataSourceController
	// @mock
	DataSourceService interface {
		repository.DataSourceRepo
	}
	// DataSourceServiceImpl is implementation of DataSourceService
	DataSourceServiceImpl struct {
		dig.In
		dbtxn.Transactional
		repository.DataSourceRepo
		AuditTrail AuditTrailService
	}
)

// NewDataSourceService return new instance of DataSourceService
// @ctor
func NewDataSourceService(impl DataSourceServiceImpl) DataSourceService {
	return &impl
}

// Insert data source
func (s *DataSourceServiceImpl) Insert(ctx context.Context, data repository.DataSource) (newID int64, err error) {
	if data.ID, err = s.DataSourceRepo.Insert(ctx, data); err != nil {
		return
	}

	s.AuditTrail.RecordInsert(ctx, "data_sources", data.ID, data)
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

	s.AuditTrail.RecordUpdate(ctx, "data_sources", data.ID, oldData, data)
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

	s.AuditTrail.RecordDelete(ctx, "data_sources", id, oldData)
	return nil
}
