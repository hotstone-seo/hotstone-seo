package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

var (
	// SettingCols is columns of setting
	SettingCols = struct {
		Key   string
		Value string
	}{
		Key:   "key",
		Value: "value",
	}
)

type (
	// Setting entity
	Setting struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	// SettingRepo to get Setting entity
	// @mock
	SettingRepo interface {
		Find(context.Context, ...dbkit.SelectOption) ([]*Setting, error)
		Update(context.Context, *Setting, dbkit.UpdateOption) error
	}

	// SettingRepoImpl is implementation SettingRepo
	SettingRepoImpl struct {
		dig.In
		*sql.DB
	}
)

// NewSettingRepo return new instance of SettingRepo
// @ctor
func NewSettingRepo() SettingRepo {
	return &SettingRepoImpl{}
}

// Find setting
func (*SettingRepoImpl) Find(context.Context, ...dbkit.SelectOption) ([]*Setting, error) {
	return nil, errors.New("Not implemented")
}

// Update setting
func (*SettingRepoImpl) Update(context.Context, *Setting, dbkit.UpdateOption) error {
	return errors.New("Not implemented")
}
