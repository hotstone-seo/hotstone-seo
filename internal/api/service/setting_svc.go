package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
	"go.uber.org/dig"
)

type (
	// SettingSvc contain logic of setting controller
	// @mock
	SettingSvc interface {
		Find(ctx context.Context) ([]*repository.Setting, error)
		FindOne(ctx context.Context, key string) (*repository.Setting, error)
		Update(ctx context.Context, key string, setting *repository.Setting) (err error)
	}

	// SettingSvcImpl is implementation of SettingService
	SettingSvcImpl struct {
		dig.In
		repository.SettingRepo
	}
)

// NewSettingSvc return new instance of setting
// @ctor
func NewSettingSvc() SettingSvc {
	return &SettingSvcImpl{}
}

// Find all setting
func (s *SettingSvcImpl) Find(ctx context.Context) ([]*repository.Setting, error) {
	return s.SettingRepo.Find(ctx)
}

// FindOne setting
func (s *SettingSvcImpl) FindOne(ctx context.Context, key string) (*repository.Setting, error) {
	if key == "" {
		return nil, errvalid.New("Key is missing")
	}

	settings, err := s.SettingRepo.Find(
		ctx,
		dbkit.Equal(repository.SettingCols.Key, key),
	)
	if err != nil {
		return nil, err
	}

	if len(settings) < 1 {
		return nil, sql.ErrNoRows
	}

	return settings[0], nil
}

// Update specific setting by key
func (*SettingSvcImpl) Update(ctx context.Context, key string, setting *repository.Setting) error {
	// TODO:
	return errors.New("Not implemented")
}
