package service

import (
	"context"
	"time"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// CenterService is center related logic
type CenterService interface {
	AddMetaTag(req AddMetaTagRequest) (int64, error)
	AddTitleTag(ctx context.Context, req AddTitleTagRequest) (int64, error)
	AddCanonicalTag(ctx context.Context, req AddCanonicalTagRequest) (int64, error)
	AddScriptTag(req AddScriptTagRequest) (int64, error)
}

// CenterServiceImpl implementation of CenterService
type CenterServiceImpl struct {
	dig.In
	repository.TagRepo
}

// NewCenterService return new instance of CenterService
func NewCenterService(impl CenterServiceImpl) CenterService {
	return &impl
}

// AddMetaTag to add metaTag
func (*CenterServiceImpl) AddMetaTag(req AddMetaTagRequest) (lastInsertedID int64, err error) {
	return
}

// AddTitleTag to add titleTag
func (i *CenterServiceImpl) AddTitleTag(ctx context.Context, req AddTitleTagRequest) (lastInsertedID int64, err error) {
	lastInsertedID, err = i.TagRepo.Insert(ctx, repository.Tag{
		RuleID:     req.RuleID,
		LocaleID:   req.LocaleID,
		Type:       "title",
		Attributes: dbkit.JSON(`{}`),
		Value:      req.Title,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	})
	return
}

// AddCanonicalTag to add canonicalTag
func (i *CenterServiceImpl) AddCanonicalTag(ctx context.Context, req AddCanonicalTagRequest) (lastInsertedID int64, err error) {
	lastInsertedID, err = i.TagRepo.Insert(ctx, repository.Tag{
		RuleID:     req.RuleID,
		LocaleID:   req.LocaleID,
		Type:       "canonical",
		Attributes: dbkit.JSON(`{}`),
		Value:      req.Canonical,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	})
	return
}

// AddScriptTag to add scriptTag
func (*CenterServiceImpl) AddScriptTag(req AddScriptTagRequest) (lastInsertedID int64, err error) {
	return
}
