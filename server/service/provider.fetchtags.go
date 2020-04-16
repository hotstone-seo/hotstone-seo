package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtype"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/imantung/mario"
)

// FetchTagsWithCache is same with FetchTag but with tag
func (p *ProviderServiceImpl) FetchTagsWithCache(
	ctx context.Context,
	ruleID int64,
	locale string,
	pragma *cachekit.Pragma,
) (itags []*ITag, err error) {

	key := fmt.Sprintf("rule%d_%s", ruleID, locale)
	cache := cachekit.New(key,
		func() (interface{}, error) {
			itags, err := p.FetchTags(ctx, ruleID, locale)
			return itags, err
		},
	)

	itags = []*ITag{}
	if err = cache.Execute(p.Redis, &itags, pragma); err != nil {
		return
	}

	return
}

// FetchTags handle logic for fetching tag
func (p *ProviderServiceImpl) FetchTags(
	ctx context.Context,
	ruleID int64,
	locale string,
) (itags []*ITag, err error) {
	var (
		rule  *repository.Rule
		tags  []*repository.Tag
		ds    *IDataSource
		b     []byte
		param map[string]interface{}
		itag  *ITag
	)

	if rule, err = p.RuleRepo.FindOne(ctx, ruleID); err != nil {
		return
	}

	if tags, err = p.TagRepo.FindByRuleAndLocale(ctx, rule.ID, locale); err != nil {
		return nil, fmt.Errorf("Find-Tags: %w", err)
	}

	if len(tags) < 1 {
		return []*ITag{}, nil
	}

	if rule.DataSourceID == nil {
		for _, tag := range tags {
			itag := ITag(*tag)
			itags = append(itags, &itag)
		}
		return
	}

	if ds, err = p.findAndInterpolateDataSource(ctx, *rule.DataSourceID, map[string]interface{}{
		"id": rule.ID,
	}); err != nil {
		return nil, err
	}

	if b, err = call(ds); err != nil {
		return nil, fmt.Errorf("Call: %w", err)
	}

	if err = json.Unmarshal(b, &param); err != nil {
		return nil, fmt.Errorf("JSON: %w", err)
	}

	for _, tag := range tags {
		if itag, err = interpolateTag(tag, param); err != nil {
			return nil, fmt.Errorf("Interpolate-Tag: %w", err)
		}
		itags = append(itags, itag)
	}

	return
}

func (p *ProviderServiceImpl) findAndInterpolateDataSource(ctx context.Context, dataSourceID int64, param interface{}) (interpolated *IDataSource, err error) {
	var (
		ds *repository.DataSource
	)

	if ds, err = p.DataSourceRepo.FindOne(ctx, dataSourceID); err != nil {
		return nil, fmt.Errorf("DataSource: %w", err)
	}

	if interpolated, err = interpolateDataSource(ds, param); err != nil {
		return nil, fmt.Errorf("Interpolate-DataSource: %w", err)
	}

	return
}

func interpolateDataSource(ds *repository.DataSource, data interface{}) (*IDataSource, error) {
	var (
		buf  bytes.Buffer
		tmpl *mario.Template
		err  error
	)

	if tmpl, err = mario.New().Parse(ds.URL); err != nil {
		return nil, err
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return &IDataSource{
		ID:        ds.ID,
		Name:      ds.Name,
		URL:       buf.String(),
		UpdatedAt: ds.UpdatedAt,
		CreatedAt: ds.CreatedAt,
	}, nil
}

func interpolateTag(tag *repository.Tag, data interface{}) (*ITag, error) {
	var (
		attribute dbtype.JSON
		value     string
		err       error
	)
	if attribute, err = interpolateAttribute(tag.Attributes, data); err != nil {
		return nil, err
	}
	if value, err = interpolateValue(tag.Value, data); err != nil {
		return nil, err
	}
	return &ITag{
		ID:         tag.ID,
		RuleID:     tag.RuleID,
		Locale:     tag.Locale,
		Type:       tag.Type,
		Attributes: attribute,
		Value:      value,
		UpdatedAt:  tag.UpdatedAt,
		CreatedAt:  tag.CreatedAt,
	}, nil

}

func interpolateAttribute(ori dbtype.JSON, param interface{}) (interpolated dbtype.JSON, err error) {
	var (
		tmpl *mario.Template
		buf  bytes.Buffer
	)
	if tmpl, err = mario.New().Parse(string(ori)); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, param); err != nil {
		return
	}
	return buf.Bytes(), nil
}

func interpolateValue(ori string, param interface{}) (s string, err error) {
	var (
		tmpl *mario.Template
		buf  bytes.Buffer
	)
	if tmpl, err = mario.New().Parse(ori); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, param); err != nil {
		return
	}
	return buf.String(), nil
}

func call(ds *IDataSource) (data []byte, err error) {
	var (
		resp *http.Response
	)

	if resp, err = http.Get(ds.URL); err != nil {
		return data, err
	}
	defer resp.Body.Close()

	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return data, err
	}

	return data, nil
}
