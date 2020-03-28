package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
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

// Update data source
func (s *DataSourceServiceImpl) Update(ctx context.Context, ds repository.DataSource) (err error) {
	defer s.CommitMe(&ctx)()
	var oldDs *repository.DataSource
	oldDs, err = s.DataSourceRepo.FindOne(ctx, ds.ID)
	if err != nil {
		s.CancelMe(ctx, err)
		log.Warnf("ERR findone: %+v", err)
		return
	}
	if err = s.DataSourceRepo.Update(ctx, ds); err != nil {
		s.CancelMe(ctx, err)
		log.Warnf("ERR update: %+v", err)
		return
	}
	if _, err = s.AuditTrailService.RecordChanges(ctx, "data_source", ds.ID, repository.Update, oldDs, ds); err != nil {
		s.CancelMe(ctx, err)
		log.Warnf("ERR recordChange: %+v", err)
		return
	}
	return nil
}
