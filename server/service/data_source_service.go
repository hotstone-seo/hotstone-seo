package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// DataSourceService contain logic for DataSourceController [mock]
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

// NewDataSourceService return new instance of DataSourceService [constructor]
func NewDataSourceService(impl DataSourceServiceImpl) DataSourceService {
	return &impl
}

// Insert data source
func (s *DataSourceServiceImpl) Insert(ctx context.Context, ds repository.DataSource) (newDsID int64, err error) {
	defer s.CommitMe(&ctx)()
	if newDsID, err = s.DataSourceRepo.Insert(ctx, ds); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	newDS, err := s.DataSourceRepo.FindOne(ctx, newDsID)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.AuditTrailService.RecordChanges(ctx, "data_sources", newDsID, repository.Insert, nil, newDS); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	return newDsID, nil
}

// Update data source
func (s *DataSourceServiceImpl) Update(ctx context.Context, ds repository.DataSource) (err error) {
	defer s.CommitMe(&ctx)()
	oldDS, err := s.DataSourceRepo.FindOne(ctx, ds.ID)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if err = s.DataSourceRepo.Update(ctx, ds); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	newDS, err := s.DataSourceRepo.FindOne(ctx, ds.ID)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.AuditTrailService.RecordChanges(ctx, "data_sources", ds.ID, repository.Update, oldDS, newDS); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	return nil
}

// Delete data source
func (s *DataSourceServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer s.CommitMe(&ctx)()
	oldDs, err := s.DataSourceRepo.FindOne(ctx, id)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.HistoryService.RecordHistory(ctx, "data_source", id, oldDs); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if err = s.DataSourceRepo.Delete(ctx, id); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.AuditTrailService.RecordChanges(ctx, "data_sources", id, repository.Delete, oldDs, nil); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	return nil
}
