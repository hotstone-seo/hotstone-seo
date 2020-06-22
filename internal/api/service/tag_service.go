package service

import (
	"context"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
	"go.uber.org/dig"
)

type (
	// TagService provides available method to be used for managing Tag entity
	// @mock
	TagService interface {
		Find(context.Context, map[string][]string) ([]*repository.Tag, error)
		FindOne(context.Context, string) (*repository.Tag, error)
		Create(context.Context, repository.Tag) (id int64, err error)
		Update(context.Context, string, repository.Tag) error
		Delete(context.Context, string) error
	}

	// TagServiceImpl is the implementation of TagService
	TagServiceImpl struct {
		dig.In
		repository.TagRepo
		dbtxn.Transactional
		AuditTrail AuditTrailSvc
	}
)

// NewTagService returns new instance of TagService
// @constructor
func NewTagService(impl TagServiceImpl) TagService {
	return &impl
}

// Find returns list of Tag entity based on provided filters
//
// TODO: Since its possible for QueryParams to provide array of strings, we should extend
// dbkit to support "IN" query
func (s *TagServiceImpl) Find(ctx context.Context, filters map[string][]string) ([]*repository.Tag, error) {
	selectOpts := make([]dbkit.SelectOption, 0)
	if ruleFilter, exists := filters["rule_id"]; exists && len(ruleFilter) > 0 {
		ruleID, err := strconv.ParseInt(ruleFilter[0], 10, 64)
		if err != nil {
			return nil, errvalid.New("rule_id is not valid")
		}
		selectOpts = append(selectOpts, dbkit.Equal("rule_id", ruleID))
	}
	if localeFilter, exists := filters["locale"]; exists && len(localeFilter) > 0 {
		selectOpts = append(selectOpts, dbkit.Equal("locale", localeFilter[0]))
	}
	return s.TagRepo.Find(ctx, selectOpts...)
}

// FindOne returns a single Tag entity with provided ID
func (s *TagServiceImpl) FindOne(ctx context.Context, id string) (*repository.Tag, error) {
	tagID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errvalid.New("ID is not valid")
	}
	return s.TagRepo.FindOne(ctx, tagID)
}

// Create creates a new Tag entity
func (s *TagServiceImpl) Create(ctx context.Context, tag repository.Tag) (id int64, err error) {
	if tag.Attributes == nil {
		tag.Attributes = map[string]string{}
	}
	if err = tag.Validate(); err != nil {
		return
	}
	id, err = s.TagRepo.Insert(ctx, tag)
	if err != nil {
		return -1, err
	}
	tag.ID = id
	s.AuditTrail.RecordInsert(ctx, "tags", id, tag)
	return
}

// Update modify existing Tag entity
func (s *TagServiceImpl) Update(ctx context.Context, id string, tag repository.Tag) (err error) {
	tagID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errvalid.New("ID is not valid")
	}
	tag.ID = tagID
	currentTag, err := s.TagRepo.FindOne(ctx, tagID)
	if err != nil {
		return
	}
	if err = s.TagRepo.Update(ctx, tag); err == nil {
		s.AuditTrail.RecordUpdate(ctx, "tags", tag.ID, currentTag, tag)
	}
	return
}

// Delete tag
func (s *TagServiceImpl) Delete(ctx context.Context, id string) (err error) {
	tagID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errvalid.New("ID is not valid")
	}
	currentTag, err := s.TagRepo.FindOne(ctx, tagID)
	if err != nil {
		return
	}
	if err = s.TagRepo.Delete(ctx, tagID); err == nil {
		s.AuditTrail.RecordDelete(ctx, "tags", tagID, currentTag)
	}
	return
}
