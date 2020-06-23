package service

import (
	"context"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"go.uber.org/dig"
)

// CenterService is center related logic
// @mock
type CenterService interface {
	AddMetaTag(context.Context, MetaTagRequest) (*repository.Tag, error)
	UpdateMetaTag(context.Context, string, MetaTagRequest) error

	AddTitleTag(context.Context, TitleTagRequest) (*repository.Tag, error)
	UpdateTitleTag(context.Context, string, TitleTagRequest) error

	AddCanonicalTag(context.Context, CanonicalTagRequest) (*repository.Tag, error)
	UpdateCanonicalTag(context.Context, string, CanonicalTagRequest) error

	AddScriptTag(context.Context, ScriptTagRequest) (*repository.Tag, error)
	UpdateScriptTag(context.Context, string, ScriptTagRequest) error

	AddFAQPage(ctx context.Context, req FAQPageRequest) (*repository.StructuredData, error)
	UpdateFAQPage(ctx context.Context, req FAQPageRequest) error

	AddBreadcrumbList(ctx context.Context, req BreadcrumbListRequest) (*repository.StructuredData, error)
	UpdateBreadcrumbList(ctx context.Context, req BreadcrumbListRequest) error

	AddLocalBusiness(ctx context.Context, req LocalBusinessRequest) (*repository.StructuredData, error)
	UpdateLocalBusiness(ctx context.Context, req LocalBusinessRequest) error
}

// CenterServiceImpl implementation of CenterService
type CenterServiceImpl struct {
	dig.In
	TagService
	StructuredDataService
}

// NewCenterService return new instance of CenterService
// @ctor
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
	tag.ID, err = i.TagService.Create(ctx, *tag)
	return
}

// UpdateMetaTag updates existing meta tag
func (i *CenterServiceImpl) UpdateMetaTag(ctx context.Context, id string, req MetaTagRequest) error {
	return i.TagService.Update(ctx, id, repository.Tag{
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
	tag.ID, err = i.TagService.Create(ctx, *tag)
	return
}

// UpdateTitleTag updates existing title tag
func (i *CenterServiceImpl) UpdateTitleTag(ctx context.Context, id string, req TitleTagRequest) error {
	return i.TagService.Update(ctx, id, repository.Tag{
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
	tag.ID, err = i.TagService.Create(ctx, *tag)
	return
}

// UpdateCanonicalTag updates existing canonical tag
func (i *CenterServiceImpl) UpdateCanonicalTag(ctx context.Context, id string, req CanonicalTagRequest) error {
	return i.TagService.Update(ctx, id, repository.Tag{
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
	tag.ID, err = i.TagService.Create(ctx, *tag)
	return
}

// UpdateScriptTag updates existing script tag
func (i *CenterServiceImpl) UpdateScriptTag(ctx context.Context, id string, req ScriptTagRequest) error {
	return i.TagService.Update(ctx, id, repository.Tag{
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

func (i *CenterServiceImpl) AddFAQPage(ctx context.Context, req FAQPageRequest) (structData *repository.StructuredData, err error) {
	structData = &repository.StructuredData{
		RuleID: req.RuleID,
		Type:   "FAQPage",
		Data: map[string]interface{}{
			"@context":   "https://schema.org",
			"@type":      "FAQPage",
			"mainEntity": mapFAQs(req.FAQs),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	structData.ID, err = i.StructuredDataService.Insert(ctx, *structData)
	return
}

func (i *CenterServiceImpl) UpdateFAQPage(ctx context.Context, req FAQPageRequest) (err error) {
	return i.StructuredDataService.Update(ctx, repository.StructuredData{
		ID:     req.ID,
		RuleID: req.RuleID,
		Type:   "FAQPage",
		Data: map[string]interface{}{
			"@context":   "https://schema.org",
			"@type":      "FAQPage",
			"mainEntity": mapFAQs(req.FAQs),
		},
		UpdatedAt: time.Now(),
	})
}

func mapFAQs(faqs []FAQ) []map[string]interface{} {
	faqsMap := make([]map[string]interface{}, len(faqs))
	for index, faq := range faqs {
		faqsMap[index] = map[string]interface{}{
			"@type": "Question",
			"name":  faq.Question,
			"acceptedAnswer": map[string]string{
				"@type": "Answer",
				"text":  faq.Answer,
			},
		}
	}
	return faqsMap
}

func (i *CenterServiceImpl) AddBreadcrumbList(ctx context.Context, req BreadcrumbListRequest) (structData *repository.StructuredData, err error) {
	structData = &repository.StructuredData{
		RuleID: req.RuleID,
		Type:   "BreadcrumbList",
		Data: map[string]interface{}{
			"@context":        "https://schema.org",
			"@type":           "BreadcrumbList",
			"itemListElement": mapBreadcrumbItems(req.ListItem),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	structData.ID, err = i.StructuredDataService.Insert(ctx, *structData)
	return
}

func (i *CenterServiceImpl) UpdateBreadcrumbList(ctx context.Context, req BreadcrumbListRequest) (err error) {
	return i.StructuredDataService.Update(ctx, repository.StructuredData{
		ID:     req.ID,
		RuleID: req.RuleID,
		Type:   "BreadcrumbList",
		Data: map[string]interface{}{
			"@context":        "https://schema.org",
			"@type":           "BreadcrumbList",
			"itemListElement": mapBreadcrumbItems(req.ListItem),
		},
		UpdatedAt: time.Now(),
	})
}

func mapBreadcrumbItems(breadcrumbItems []BreadcrumbItem) []map[string]interface{} {
	itemsMap := make([]map[string]interface{}, len(breadcrumbItems))
	for index, breadcrumbItem := range breadcrumbItems {
		itemsMap[index] = map[string]interface{}{
			"@type":    "ListItem",
			"position": index,
			"name":     breadcrumbItem.Name,
			"item":     breadcrumbItem.Item,
		}
	}
	return itemsMap
}

func (i *CenterServiceImpl) AddLocalBusiness(ctx context.Context, req LocalBusinessRequest) (structData *repository.StructuredData, err error) {
	structData = &repository.StructuredData{
		RuleID:    req.RuleID,
		Type:      "LocalBusiness",
		Data:      req.ToSchema(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	structData.ID, err = i.StructuredDataService.Insert(ctx, *structData)
	return
}

func (i *CenterServiceImpl) UpdateLocalBusiness(ctx context.Context, req LocalBusinessRequest) error {
	return i.StructuredDataService.Update(ctx, repository.StructuredData{
		ID:        req.ID,
		RuleID:    req.RuleID,
		Type:      "LocalBusiness",
		Data:      req.ToSchema(),
		UpdatedAt: time.Now(),
	})
}
