package service_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/mock_repository"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
)

var (
	rule999            = &repository.Rule{ID: 999}
	rule777_noDS       = &repository.Rule{ID: 777}
	tags_rule777_en_US = []*repository.Tag{
		{ID: 1},
		{ID: 2},
	}
)

func TestProviderService_FetchTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dsmock := mock_repository.NewMockDataSourceRepo(ctrl)
	rulemock := mock_repository.NewMockRuleRepo(ctrl)
	tagmock := mock_repository.NewMockTagRepo(ctrl)

	svc := service.ProviderServiceImpl{
		DataSourceRepo: dsmock,
		RuleRepo:       rulemock,
		TagRepo:        tagmock,
	}

	ctx := context.Background()

	t.Run("GIVEN non-existing rule ID", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(999)).Return(nil, sql.ErrNoRows)

		_, err := svc.FetchTags(ctx, int64(999), "en_US")
		require.Equal(t, sql.ErrNoRows, err)
	})

	t.Run("WHEN database error in get tags ", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(999)).Return(rule999, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(999), "en_US").Return(nil, errors.New("some-error"))

		_, err := svc.FetchTags(ctx, int64(999), "en_US")
		require.EqualError(t, err, "Find-Tags: some-error")
	})

	t.Run("WHEN no tags", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(777)).Return(rule777_noDS, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(777), "en_US").Return([]*repository.Tag{}, nil)

		tags, err := svc.FetchTags(ctx, int64(777), "en_US")
		require.NoError(t, err)
		require.Empty(t, tags)
	})

	t.Run("WHEN rule has no datasource", func(t *testing.T) {
		rulemock.EXPECT().FindOne(ctx, int64(777)).Return(rule777_noDS, nil)
		tagmock.EXPECT().FindByRuleAndLocale(ctx, int64(777), "en_US").Return(tags_rule777_en_US, nil)

		tags, err := svc.FetchTags(ctx, int64(777), "en_US")
		require.NoError(t, err)

		// THEN return tags same with the original without interpolation
		for i, tag := range tags {
			og := service.ITag(*tags_rule777_en_US[i])
			require.Equal(t, &og, tag)
		}
	})
}
