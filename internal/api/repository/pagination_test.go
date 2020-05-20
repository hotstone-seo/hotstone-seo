package repository

import (
	"net/url"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T) {
	validColumns := []string{"id", "name", "address", "updated_at"}
	baseBuilder := sq.Select("id", "name", "address", "updated_at").
		From("people")
	emptySorts := []Sort{}
	emptyFilters := map[string]*Filter{}
	emptyNexts := []Next{}

	tests := []struct {
		name    string
		url     string
		pgnType PaginationType
		sql     string
		sqlArgs []interface{}
		param   PaginationParam
	}{
		{url: "/people?_sort=-name,address,notvalidcolumn", pgnType: NoPagination, sql: "SELECT id, name, address, updated_at FROM people ORDER BY name DESC, address ASC", param: PaginationParam{Sorts: []Sort{{Col: "name", Order: "DESC"}, {Col: "address", Order: "ASC"}}, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_offset=1&_limit=15", pgnType: OffsetPagination, sql: "SELECT id, name, address, updated_at FROM people LIMIT 15 OFFSET 1", param: PaginationParam{Offset: 1, Limit: 15, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?name=foo-name&address=%25foo-address%25", pgnType: NoPagination, sql: "SELECT id, name, address, updated_at FROM people WHERE address LIKE ? AND name = ?", sqlArgs: []interface{}{"%foo-address%", "foo-name"}, param: PaginationParam{Sorts: emptySorts, Filters: map[string]*Filter{"name": {Col: "name", Cond: "foo-name"}, "address": {Col: "address", Cond: "%foo-address%"}}, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id", pgnType: KeysetPagination, sql: "SELECT id, name, address, updated_at FROM people ORDER BY id DESC", param: PaginationParam{NextKey: &Next{Col: "id", Order: "DESC"}, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id&_limit=15", pgnType: KeysetPagination, sql: "SELECT id, name, address, updated_at FROM people ORDER BY id DESC LIMIT 15", param: PaginationParam{NextKey: &Next{Col: "id", Order: "DESC"}, Limit: 15, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id&_next=27", pgnType: KeysetPagination, sql: "SELECT id, name, address, updated_at FROM people WHERE id < ? ORDER BY id DESC", sqlArgs: []interface{}{"27"}, param: PaginationParam{NextKey: &Next{Col: "id", Order: "DESC", Value: "27"}, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id&_sort=updated_at&_next=1581495514,27", sql: "SELECT id, name, address, updated_at FROM people WHERE (updated_at > ? OR (updated_at = ? AND id < ?)) ORDER BY updated_at ASC, id DESC", sqlArgs: []interface{}{"1581495514", "1581495514", "27"}, pgnType: KeysetPagination, param: PaginationParam{NextKey: &Next{Col: "id", Order: "DESC", Value: "27"}, Sorts: []Sort{{Col: "updated_at", Order: "ASC"}}, Filters: emptyFilters, Nexts: []Next{{Col: "updated_at", Order: "ASC", Value: "1581495514"}}}},
		{url: "/people?_next_key=-id&_sort=updated_at,-name&_next=1581495514,abc,27", sql: "SELECT id, name, address, updated_at FROM people WHERE ((updated_at > ? AND name < ?) OR ((updated_at = ? AND name = ?) AND id < ?)) ORDER BY updated_at ASC, name DESC, id DESC", sqlArgs: []interface{}{"1581495514", "abc", "1581495514", "abc", "27"}, pgnType: KeysetPagination, param: PaginationParam{NextKey: &Next{Col: "id", Order: "DESC", Value: "27"}, Sorts: []Sort{{Col: "updated_at", Order: "ASC"}, {Col: "name", Order: "DESC"}}, Filters: emptyFilters, Nexts: []Next{{Col: "updated_at", Order: "ASC", Value: "1581495514"}, {Col: "name", Order: "DESC", Value: "abc"}}}},
	}
	for _, tt := range tests {
		paginationParam := BuildPaginationParam(buildQueryParam(t, tt.url), validColumns)
		sql, sqlArgs, err := ComposePagination(baseBuilder, paginationParam).ToSql()
		// t.Logf("=== %d ===", i)
		require.NoError(t, err)
		require.Equal(t, tt.param, paginationParam)
		require.Equal(t, tt.pgnType, GetPaginationType(paginationParam))
		require.Equal(t, tt.sql, sql)
		require.Equal(t, tt.sqlArgs, sqlArgs)
	}
}

func buildQueryParam(t *testing.T, urlStr string) url.Values {
	t.Helper()

	u, err := url.Parse(urlStr)
	require.NoError(t, err)

	return u.Query()
}
