package service

import (
	"context"
	"time"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// CenterService is center related logic [mock]
type CenterService interface {
	AddMetaTag(ctx context.Context, req MetaTagRequest) (*repository.Tag, error)
	UpdateMetaTag(ctx context.Context, req MetaTagRequest) error
	AddTitleTag(ctx context.Context, req TitleTagRequest) (*repository.Tag, error)
	UpdateTitleTag(ctx context.Context, req TitleTagRequest) error
	AddCanonicalTag(ctx context.Context, req CanonicalTagRequest) (*repository.Tag, error)
	UpdateCanonicalTag(ctx context.Context, req CanonicalTagRequest) error
	AddScriptTag(ctx context.Context, req ScriptTagRequest) (*repository.Tag, error)
	UpdateScriptTag(ctx context.Context, req ScriptTagRequest) error
}

// CenterServiceImpl implementation of CenterService
type CenterServiceImpl struct {
	dig.In
	TagService
}

// NewCenterService return new instance of CenterService [constructor]
func NewCenterService(impl CenterServiceImpl) CenterService {
	return &impl
}

// AddMetaTag adds new meta tag
// TODO: Use JSON marshal to set attributes, simple string substitution is prone to be exploited
func (i *CenterServiceImpl) AddMetaTag(ctx context.Context, req MetaTagRequest) (tag *repository.Tag, err error) {
	tag = &repository.Tag{
		RuleID: req.RuleID,
		Locale: req.Locale,
		Type:   "meta",
		Attributes: map[string]string{
			"name":    req.Name,
			"content": req.Content,
		},
		Value:     "",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	tag.ID, err = i.TagService.Insert(ctx, *tag)
	return
}

// UpdateMetaTag updates existing meta tag
func (i *CenterServiceImpl) UpdateMetaTag(ctx context.Context, req MetaTagRequest) error {
	return i.TagService.Update(ctx, repository.Tag{
		ID:     req.ID,
		RuleID: req.RuleID,
		Locale: req.Locale,
		Type:   "meta",
		Attributes: map[string]string{
			"name": req.Name,
			"req":  req.Content,
		},
		Value:     "",
		UpdatedAt: time.Now(),
	})
}

// AddTitleTag adds new title tag
func (i *CenterServiceImpl) AddTitleTag(ctx context.Context, req TitleTagRequest) (tag *repository.Tag, err error) {
	tag = &repository.Tag{
		RuleID:     req.RuleID,
		Locale:     req.Locale,
		Type:       "title",
		Attributes: map[string]string{},
		Value:      req.Title,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	}
	tag.ID, err = i.TagService.Insert(ctx, *tag)
	return
}

// UpdateTitleTag updates existing title tag
func (i *CenterServiceImpl) UpdateTitleTag(ctx context.Context, req TitleTagRequest) error {
	return i.TagService.Update(ctx, repository.Tag{
		ID:         req.ID,
		RuleID:     req.RuleID,
		Locale:     req.Locale,
		Type:       "title",
		Attributes: map[string]string{},
		Value:      req.Title,
		UpdatedAt:  time.Now(),
	})
}

// AddCanonicalTag adds new canonical tag
func (i *CenterServiceImpl) AddCanonicalTag(ctx context.Context, req CanonicalTagRequest) (tag *repository.Tag, err error) {
	tag = &repository.Tag{
		RuleID: req.RuleID,
		Locale: req.Locale,
		Type:   "link",
		Attributes: map[string]string{
			"href": req.Href,
			"rel":  "canonical",
		},
		Value:     "",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	tag.ID, err = i.TagService.Insert(ctx, *tag)
	return
}

// UpdateCanonicalTag updates existing canonical tag
func (i *CenterServiceImpl) UpdateCanonicalTag(ctx context.Context, req CanonicalTagRequest) error {
	return i.TagService.Update(ctx, repository.Tag{
		ID:     req.ID,
		RuleID: req.RuleID,
		Locale: req.Locale,
		Type:   "link",
		Attributes: map[string]string{
			"href": req.Href,
			"rel":  "canonical",
		},
		Value:     "",
		UpdatedAt: time.Now(),
	})
}

// AddScriptTag adds new script tag
func (i *CenterServiceImpl) AddScriptTag(ctx context.Context, req ScriptTagRequest) (tag *repository.Tag, err error) {
	tag = &repository.Tag{
		RuleID: req.RuleID,
		Locale: req.Locale,
		Type:   "script",
		Attributes: map[string]string{
			"source": req.Source,
		},
		Value:     req.Type,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	tag.ID, err = i.TagService.Insert(ctx, *tag)
	return
}

// UpdateScriptTag updates existing script tag
func (i *CenterServiceImpl) UpdateScriptTag(ctx context.Context, req ScriptTagRequest) error {
	return i.TagService.Update(ctx, repository.Tag{
		ID:     req.ID,
		RuleID: req.RuleID,
		Locale: req.Locale,
		Type:   "script",
		Attributes: map[string]string{
			"source": req.Source,
		},
		Value:     req.Type,
		UpdatedAt: time.Now(),
	})
}
