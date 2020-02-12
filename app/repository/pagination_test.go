package repository

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildPaginationParam(t *testing.T) {
	validColumns := []string{"id", "name", "address", "updated_at"}
	emptySorts := []Sort{}
	emptyFilters := map[string]*Filter{}
	emptyNexts := []Next{}

	tests := []struct {
		url                 string
		wantPaginationParam PaginationParam
		wantPaginationType  PaginationType
	}{
		{url: "/people?_sort=-name,address", wantPaginationType: NoPagination, wantPaginationParam: PaginationParam{Sorts: []Sort{Sort{Col: "name", Order: "DESC"}, Sort{Col: "address", Order: "ASC"}}, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_start=0&_end=15", wantPaginationType: OffsetPagination, wantPaginationParam: PaginationParam{Start: 0, End: 15, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?name=foo-name&address=%25foo-address%25", wantPaginationType: NoPagination, wantPaginationParam: PaginationParam{Sorts: emptySorts, Filters: map[string]*Filter{"name": &Filter{Col: "name", Cond: "foo-name"}, "address": &Filter{Col: "address", Cond: "%foo-address%"}}, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id", wantPaginationType: KeysetPagination, wantPaginationParam: PaginationParam{NextKey: &NextKey{Col: "id", Order: "DESC"}, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id&_end=15", wantPaginationType: KeysetPagination, wantPaginationParam: PaginationParam{NextKey: &NextKey{Col: "id", Order: "DESC"}, End: 15, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id&_next=27", wantPaginationType: KeysetPagination, wantPaginationParam: PaginationParam{NextKey: &NextKey{Col: "id", Order: "DESC", Value: "27"}, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_next_key=-id&_sort=updated_at&_next=1581495514,27", wantPaginationType: KeysetPagination, wantPaginationParam: PaginationParam{NextKey: &NextKey{Col: "id", Order: "DESC", Value: "27"}, Sorts: []Sort{Sort{Col: "updated_at", Order: "ASC"}}, Filters: emptyFilters, Nexts: []Next{Next{Col: "updated_at", Value: "1581495514"}}}},
	}
	for _, tt := range tests {
		paginatonParam := BuildPaginationParam(buildQueryParam(t, tt.url), validColumns)
		require.Equal(t, tt.wantPaginationParam, paginatonParam)
		require.Equal(t, tt.wantPaginationType, GetPaginationType(paginatonParam))
	}
}

func buildQueryParam(t *testing.T, urlStr string) url.Values {
	t.Helper()

	u, err := url.Parse(urlStr)
	require.NoError(t, err)

	t.Logf("#QUERY: %+v", u.Query())
	return u.Query()
}
