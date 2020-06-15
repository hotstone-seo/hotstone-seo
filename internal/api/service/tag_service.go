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
	defer func() {
		if err != nil {
			return
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			Record{
				EntityName: "tags",
				EntityID:   newID,
				Operation:  InsertOp,
				PrevData:   nil,
				NextData:   tag,
			},
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	if tag.Attributes == nil {
		tag.Attributes = map[string]string{}
	}
	return s.TagRepo.Insert(ctx, tag)
}

// Update modify existing Tag entity
func (s *TagServiceImpl) Update(ctx context.Context, tag repository.Tag) (err error) {
	var currentTag *repository.Tag
	if currentTag, err = s.TagRepo.FindOne(ctx, tag.ID); err != nil {
		return
	}
	defer func() {
		if err != nil {
			return
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			Record{
				EntityName: "tags",
				EntityID:   tag.ID,
				Operation:  UpdateOp,
				PrevData:   currentTag,
				NextData:   tag,
			},
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return s.TagRepo.Update(ctx, tag)
}

// Delete tag
func (s *TagServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var currentTag *repository.Tag
	if currentTag, err = s.TagRepo.FindOne(ctx, id); err != nil {
		return
	}
	defer func() {
		if err != nil {
			return
		}
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			currentTag.Type+"-tag",
			id,
			currentTag,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			Record{
				EntityName: "tags",
				EntityID:   id,
				Operation:  DeleteOp,
				PrevData:   currentTag,
				NextData:   nil,
			},
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return s.TagRepo.Delete(ctx, id)
}
