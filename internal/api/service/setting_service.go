package service

import (
	"context"
	"errors"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
)

type (
	// SettingService contain logic of setting controller
	// @mock
	SettingService interface {
		Find(ctx context.Context) ([]*repository.Setting, error)
		FindOne(ctx context.Context, key string) (*repository.Setting, error)
		Update(ctx context.Context, key string, setting *repository.Setting) (err error)
	}

	// SettingServiceImpl is implementation of SettingService
	SettingServiceImpl struct {
	}
)

// NewSettingService return new instance of setting
// @ctor
func NewSettingService() SettingService {
	return &SettingServiceImpl{}
}

// Find all setting
func (*SettingServiceImpl) Find(ctx context.Context) ([]*repository.Setting, error) {
	// TODO:
	return nil, errors.New("Not implemented")
}

// FindOne setting
func (*SettingServiceImpl) FindOne(ctx context.Context, key string) (*repository.Setting, error) {
	// TODO:
	return nil, errors.New("Not implemented")
}

// Update specific setting by key
func (*SettingServiceImpl) Update(ctx context.Context, key string, setting *repository.Setting) error {
	// TODO:
	return errors.New("Not implemented")
}
