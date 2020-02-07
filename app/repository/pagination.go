package repository

import (
	"net/url"
	"strconv"
	"strings"
)

func isColInValidColumns(col string, validColumns []string) bool {
	for _, c := range validColumns {
		if c == col {
			return true
		}
	}
	return false
}

func extractColSortAndOrder(colQueryParam string) (col, order string) {
	if colQueryParam == "" {
		return
	}

	if colQueryParam[0] == '-' {
		col = colQueryParam[1:]
		order = "DESC"
	} else {
		col = colQueryParam
		order = "ASC"
	}

	return
}

func BuildPaginationParam(queryParams url.Values, validColumns []string) PaginationParam {
	paginationParam := PaginationParam{}

	sorts := map[string]string{}
	for _, colSort := range strings.Split(queryParams.Get("_sort"), ",") {
		col, order := extractColSortAndOrder(colSort)
		if isColInValidColumns(col, validColumns) {
			sorts[col] = order
		}

	}
	paginationParam.Sorts = sorts

	start, _ := strconv.Atoi(queryParams.Get("_start"))
	paginationParam.Start = start
	end, _ := strconv.Atoi(queryParams.Get("_end"))
	paginationParam.End = end

	filters := map[string]string{}
	for _, col := range validColumns {
		whereCond := queryParams.Get(col)
		if whereCond != "" {
			filters[col] = whereCond
		}
	}
	paginationParam.Filters = filters

	return paginationParam
}
