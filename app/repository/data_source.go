package repository

import (
	"context"
	"time"
)

// DataSource represented  data_source entity
type DataSource struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// DataSourceRepo to handle data_sources entity [mock]
type DataSourceRepo interface {
	FindOne(context.Context, int64) (*DataSource, error)
	Find(context.Context) ([]*DataSource, error)
	Insert(context.Context, DataSource) (lastInsertID int64, err error)
	Delete(context.Context, int64) error
	Update(context.Context, DataSource) error
}

// NewDataSourceRepo return new instance of DataSourceRepo [autowire]
func NewDataSourceRepo(impl CachedDataSourceRepoImpl) DataSourceRepo {
	return &impl
}
