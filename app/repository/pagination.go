package repository

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

type PaginationParam struct {
	Sorts   map[string]string
	Start   int
	End     int
	Filters map[string]string
}

func composePagination(base sq.SelectBuilder, paginationParam PaginationParam) sq.SelectBuilder {

	for col, order := range paginationParam.Sorts {
		base = base.OrderBy(fmt.Sprintf("%s %s", col, order))
	}

	base = base.Offset(uint64(paginationParam.Start))

	if paginationParam.End != 0 {
		base = base.Limit(uint64(paginationParam.End - paginationParam.Start + 1))
	}

	for col, whereCond := range paginationParam.Filters {
		if strings.ContainsAny(whereCond, "%") {
			base = base.Where(sq.Like{col: whereCond})
		} else {
			base = base.Where(sq.Eq{col: whereCond})
		}
	}

	return base
}

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
