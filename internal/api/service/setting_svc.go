package service

import (
	"context"
	"database/sql"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

const (
	// SimulationKey is HotStone client key for Simulation
	SimulationKey string = "simulation_key"
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
func NewSettingSvc(impl SettingSvcImpl) SettingSvc {
	return &impl
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
func (s *SettingSvcImpl) Update(ctx context.Context, key string, setting *repository.Setting) error {
	if key == "" {
		return errvalid.New("key is missing")
	}
	if err := validator.New().Struct(setting); err != nil {
		return err
	}
	return s.SettingRepo.Update(ctx,
		setting,
		dbkit.Equal(repository.SettingCols.Key, key),
	)
}
