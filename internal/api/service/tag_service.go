package service

import (
	"context"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// TagService provides available method to be used for managing Tag entity
// @mock
type TagService interface {
	FindByRuleAndLocale(ctx context.Context, ruleID int64, locale string) ([]*repository.Tag, error)
	repository.TagRepo
}

// TagServiceImpl is the implementation of TagService
type TagServiceImpl struct {
	dig.In
	repository.TagRepo
	dbtxn.Transactional
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewTagService return new instance of TagService
// @constructor
func NewTagService(impl TagServiceImpl) TagService {
	return &impl
}

// FindByRuleAndLocale returns list of Tag entity by searching on matching ruleID and locale
func (s *TagServiceImpl) FindByRuleAndLocale(ctx context.Context, ruleID int64, locale string) (list []*repository.Tag, err error) {
	return s.Find(ctx,
		dbkit.Equal("rule_id", strconv.FormatInt(ruleID, 10)),
		dbkit.Equal("locale", locale),
	)
}

// Insert creates a new Tag entity
func (s *TagServiceImpl) Insert(ctx context.Context, tag repository.Tag) (newID int64, err error) {
	if tag.Attributes == nil {
		tag.Attributes = map[string]string{}
	}
	if newID, err = s.TagRepo.Insert(ctx, tag); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"tags",
			newID,
			repository.Insert,
			nil,
			tag,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return
}

// Update modify existing Tag entity
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
