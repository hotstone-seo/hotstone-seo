package repository_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
)

func TestTagValidation(t *testing.T) {
	t.Run("When rule ID is missing", func(t *testing.T) {
		tag := repository.Tag{Locale: "en_US", Type: "title", Value: "Some Title"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.RuleID' Error:Field validation for 'RuleID' failed on the 'required' tag")
	})
	t.Run("When locale is missing", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Type: "title", Value: "Some Title"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Locale' Error:Field validation for 'Locale' failed on the 'required' tag")
	})
	t.Run("When type is missing", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Locale: "en_US"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Type' Error:Field validation for 'Type' failed on the 'required' tag")
	})
	t.Run("When valid title tag", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Locale: "en_US", Type: "title", Value: "Some Title"}
		require.NoError(t, tag.Validate())
	})
	t.Run("When value is missing from title tag", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Locale: "en_US", Type: "title"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Value' Error:Field validation for 'Value' failed on the 'noempty' tag")
	})
	t.Run("When valid meta tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "meta",
			Attributes: []byte(`{ "name": "description", "content": "Some description" }`),
		}
		require.NoError(t, tag.Validate())
	})
	t.Run("When 'name' key is missing from attribute in meta tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "meta",
			Attributes: []byte(`{ "content": "Some description" }`),
		}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Attributes' Error:Field validation for 'Attributes' failed on the '' tag")
	})
	t.Run("When 'content' key is missing from attribute in meta tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "meta",
			Attributes: []byte(`{ "name": "description" }`),
		}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Attributes' Error:Field validation for 'Attributes' failed on the '' tag")
	})
	t.Run("When valid canonical tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "link",
			Attributes: []byte(`{ "rel": "canonical", "href": "http://original-site.com" }`),
		}
		require.NoError(t, tag.Validate())
	})
	t.Run("When 'rel' key is missing from attribute in canonical tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "link",
			Attributes: []byte(`{ "href": "http://original-site.com" }`),
		}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Attributes' Error:Field validation for 'Attributes' failed on the '' tag")
	})
	t.Run("When 'href' key is missing from attribute in canonical tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "link",
			Attributes: []byte(`{ "rel": "canonical" }`),
		}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Attributes' Error:Field validation for 'Attributes' failed on the '' tag")
	})
	t.Run("When valid script tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "script",
			Attributes: []byte(`{ "src": "http://mysite.com/script.js" }`),
		}
		require.NoError(t, tag.Validate())
	})
	t.Run("When 'src' key is missing from attribute in script tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "script",
			Attributes: []byte(`{}`),
		}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Attributes' Error:Field validation for 'Attributes' failed on the '' tag")
	})
}

func TestTag_FindByRuleAndLocale(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.TagRepoImpl{DB: typpostgres.NewDB(db)}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, rule_id, locale, type, attributes, value, updated_at, created_at FROM tags WHERE rule_id = $1 AND locale = $2")).
		WithArgs("123", "en_US").
		WillReturnRows(sqlmock.NewRows([]string{"id", "rule_id", "locale", "type", "attributes", "value", "updated_at", "created_at"}))
	_, err = repo.FindByRuleAndLocale(context.Background(), 123, "en_US")
	require.NoError(t, err)

}
