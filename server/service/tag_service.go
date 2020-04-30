package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// TagService contain logic for TagController [mock]
type TagService interface {
	repository.TagRepo
}

// TagServiceImpl is implementation of TagService
type TagServiceImpl struct {
	dig.In
	repository.TagRepo
	dbtxn.Transactional
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewTagService return new instance of TagService [constructor]
func NewTagService(impl TagServiceImpl) TagService {
	return &impl
}

// Insert tag
func (s *TagServiceImpl) Insert(ctx context.Context, data repository.Tag) (newID int64, err error) {
	if data.ID, err = s.TagRepo.Insert(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"tags",
			data.ID,
			repository.Insert,
			nil,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return data.ID, nil
}

// Update tag
func (s *TagServiceImpl) Update(ctx context.Context, data repository.Tag) (err error) {
	var oldData *repository.Tag
	if oldData, err = s.TagRepo.FindOne(ctx, data.ID); err != nil {
		return
	}
	if err = s.TagRepo.Update(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"tags",
			data.ID,
			repository.Update,
			oldData,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

// Delete tag
func (s *TagServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var oldData *repository.Tag
	if oldData, err = s.TagRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = s.TagRepo.Delete(ctx, id); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	go func() {
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			oldData.Type+"-tag",
			id,
			oldData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"tags",
			id,
			repository.Delete,
			oldData,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}
