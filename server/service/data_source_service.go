package service

import (
	"context"

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
	repository.DataSourceRepo
	repository.Transactional
	AuditTrailService
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
	if _, err = s.AuditTrailService.RecordChanges(ctx, "data_source", newDsID, repository.Insert, nil, ds); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	return newDsID, nil
}
