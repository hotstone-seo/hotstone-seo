package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/imantung/mario"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
	"golang.org/x/sync/errgroup"
)

const (
	ruleParam   = "_rule"
	localeParam = "_locale"
)

// FetchTagsWithCache is same with FetchTag but with tag
func (p *ProviderServiceImpl) FetchTagsWithCache(ctx context.Context, vals url.Values, pragma *cachekit.Pragma) (itags []*ITag, err error) {
	key := vals.Encode()
	cache := cachekit.New(key,
		func() (interface{}, error) {
			itags, err := p.FetchTags(ctx, vals)
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
func (p *ProviderServiceImpl) FetchTags(ctx context.Context, vals url.Values) (itags []*ITag, err error) {
	var (
		rule            *repository.Rule
		tags            []*repository.Tag
		structuredDatas []*repository.StructuredData
		itag            *ITag
	)

	locale := vals.Get(localeParam)
	ruleID, _ := strconv.ParseInt(vals.Get(ruleParam), 10, 64)

	if ruleID < 1 {
		return nil, errvalid.New("Missing url param for `ID`")
	}

	if locale == "" {
		return nil, errvalid.New("Missing query param for `Locale`")
	}

	if rule, err = p.RuleRepo.FindOne(ctx, ruleID); err != nil {
		return
	}

	if tags, err = p.TagRepo.FindByRuleAndLocale(ctx, rule.ID, locale); err != nil {
		return nil, fmt.Errorf("Find-Tags: %w", err)
	}

	if structuredDatas, err = p.StructuredDataRepo.FindByRule(ctx, rule.ID); err != nil {
		return nil, fmt.Errorf("Find-StructuredDatas: %w", err)
	}

	for _, strData := range structuredDatas {
		strDataTag := strData.ToTag()
		tags = append(tags, &strDataTag)
	}

	if len(tags) < 1 {
		return []*ITag{}, nil
	}

	if len(rule.DataSourceIDs) == 0 {
		for _, tag := range tags {
			itag := ITag(*tag)
			itags = append(itags, &itag)
		}
		return
	}

	findDS, ctx := errgroup.WithContext(ctx)
	sources := make([]*IDataSource, len(rule.DataSourceIDs))
	for index, dsID := range rule.DataSourceIDs {
		index, dsID := index, dsID
		findDS.Go(func() error {
			var (
				ds      *IDataSource
				findErr error
			)
			if ds, findErr = p.findAndInterpolateDataSource(ctx, dsID, ConvertToParams(vals)); findErr != nil {
				return findErr
			}
			sources[index] = ds
			return nil
		})
	}
	if err = findDS.Wait(); err != nil {
		return
	}

	callDS, ctx := errgroup.WithContext(ctx)
	apiData := make(map[string]interface{})
	for _, source := range sources {
		source := source // https://golang.org/doc/faq#closures_and_goroutines
		callDS.Go(func() error {
			var (
				b      []byte
				data   interface{}
				apiErr error
			)
			if b, apiErr = call(source); apiErr != nil {
				return fmt.Errorf("Call: %w", apiErr)
			}
			if data, apiErr = UnmarshalData(b); apiErr != nil {
				return apiErr
			}
			apiData[source.Name] = data
			return nil
		})
	}
	if err = callDS.Wait(); err != nil {
		return
	}

	for _, tag := range tags {
		if itag, err = interpolateTag(tag, apiData); err != nil {
			return nil, fmt.Errorf("Interpolate-Tag: %w", err)
		}
		itags = append(itags, itag)
	}

	return
}

// UnmarshalData to unmarshal data
func UnmarshalData(b []byte) (v interface{}, err error) {
	firstChar := string(b)[0]

	if firstChar == '{' {
		var param map[string]interface{}
		if err = json.Unmarshal(b, &param); err != nil {
			return nil, fmt.Errorf("JSON: %w", err)
		}
		return param, nil
	}

	if firstChar == '[' {
		var params []map[string]interface{}
		if err = json.Unmarshal(b, &params); err != nil {
			return nil, fmt.Errorf("JSON: %w", err)
		}
		if len(params) > 0 {
			return params[0], nil
		}
		return map[string]interface{}{}, nil
	}

	return nil, errors.New("Unmarshal-Data: Invalid data format")
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

func interpolateTag(tag *repository.Tag, data map[string]interface{}) (*ITag, error) {
	var (
		attribute map[string]string
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

func interpolateAttribute(ori map[string]string, param map[string]interface{}) (interpolated map[string]string, err error) {
	var (
		tmpl *mario.Template
		buf  bytes.Buffer
		b    []byte
	)
	if b, err = json.Marshal(ori); err != nil {
		return
	}
	if tmpl, err = mario.New().Parse(string(b)); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, param); err != nil {
		return
	}
	interpolated = make(map[string]string)
	if err = json.Unmarshal(buf.Bytes(), &interpolated); err != nil {
		return nil, err
	}
	return interpolated, nil
}

func interpolateValue(ori string, param map[string]interface{}) (s string, err error) {
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

// ConvertToParams to convert url.Values to params
func ConvertToParams(vals url.Values) map[string]string {
	params := make(map[string]string)
	for key, val := range vals {
		if len(val) > 0 {
			params[key] = val[0]
		}
	}
	return params
}
