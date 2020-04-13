package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
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
	defer s.CommitMe(&ctx)()
	if newID, err = s.TagRepo.Insert(ctx, data); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	newData, err := s.TagRepo.FindOne(ctx, newID)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.AuditTrailService.RecordChanges(ctx, "tags", newID, repository.Insert, nil, newData); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	return newID, nil
}

// Update tag
func (s *TagServiceImpl) Update(ctx context.Context, data repository.Tag) (err error) {
	defer s.CommitMe(&ctx)()
	oldData, err := s.TagRepo.FindOne(ctx, data.ID)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if err = s.TagRepo.Update(ctx, data); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	newData, err := s.TagRepo.FindOne(ctx, data.ID)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.AuditTrailService.RecordChanges(ctx, "tags", data.ID, repository.Update, oldData, newData); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	return nil
}

// Delete tag
func (s *TagServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer s.CommitMe(&ctx)()
	oldData, err := s.TagRepo.FindOne(ctx, id)
	if err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.HistoryService.RecordHistory(ctx, oldData.Type+"-tag", id, oldData); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if err = s.TagRepo.Delete(ctx, id); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	if _, err = s.AuditTrailService.RecordChanges(ctx, "tags", id, repository.Delete, oldData, nil); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	return nil
}
