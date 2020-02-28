package service

import (
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
}

// NewDataSourceService return new instance of DataSourceService [constructor]
func NewDataSourceService(impl DataSourceServiceImpl) DataSourceService {
	return &impl
}
