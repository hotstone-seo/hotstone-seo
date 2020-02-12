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
	}{
		{url: "/people?_sort=-name,address", wantPaginationParam: PaginationParam{Sorts: []Sort{Sort{Col: "name", Order: "DESC"}, Sort{Col: "address", Order: "ASC"}}, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?_start=0&_end=15", wantPaginationParam: PaginationParam{Start: 0, End: 15, Sorts: emptySorts, Filters: emptyFilters, Nexts: emptyNexts}},
		{url: "/people?name=foo-name&address=%25foo-address%25", wantPaginationParam: PaginationParam{Sorts: emptySorts, Filters: map[string]*Filter{"name": &Filter{Col: "name", Cond: "foo-name"}, "address": &Filter{Col: "address", Cond: "%foo-address%"}}, Nexts: emptyNexts}},
	}
	for _, tt := range tests {
		require.Equal(t, tt.wantPaginationParam, BuildPaginationParam(buildQueryParam(t, tt.url), validColumns))
	}
}

func buildQueryParam(t *testing.T, urlStr string) url.Values {
	t.Helper()

	u, err := url.Parse(urlStr)
	require.NoError(t, err)

	t.Logf("#QUERY: %+v", u.Query())
	return u.Query()
}
