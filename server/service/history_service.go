package service

import (
	"context"
	"encoding/json"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// HistoryService contain logic for History Controller [mock]
type HistoryService interface {
	RecordHistory(ctx context.Context, entityFrom string, entityID int64,
		data interface{}) (lastInsertID int64, err error)
}

// HistoryServiceImpl is implementation of HistoryService
type HistoryServiceImpl struct {
	dig.In
	HistoryRepo repository.HistoryRepo
}

// NewHistoryService return new instance of HistoryService [constructor]
func NewHistoryService(impl HistoryServiceImpl) HistoryService {
	return &impl
}

// RecordHistory insert history
func (r *HistoryServiceImpl) RecordHistory(ctx context.Context,
	entityFrom string, entityID int64,
	data interface{}) (lastInsertID int64, err error) {

	dataJSON := repository.JSON("{}")
	if data != nil {
		dataJSON, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	History := repository.History{
		EntityID:   entityID,
		EntityFrom: string(entityFrom),
		Username:   repository.GetUsername(ctx),
		Data:       dataJSON,
	}

	return r.HistoryRepo.Insert(ctx, History)
}
