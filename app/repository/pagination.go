package repository

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/thoas/go-funk"
)

type PaginationParam struct {
	Sorts   []Sort
	Start   uint64
	End     uint64
	Filters map[string]*Filter
	NextKey *NextKey
	Nexts   []Next
}

type Sort struct {
	Col   string
	Order string
}

type Filter struct {
	Col  string
	Cond string
}

type Next struct {
	Col   string
	Value string
}

type NextKey struct {
	Col   string
	Order string
	Value string
}

type PaginationType int

const (
	NoPagination PaginationType = iota
	OffsetPagination
	KeysetPagination
)

func composePagination(base sq.SelectBuilder, paginationParam PaginationParam) sq.SelectBuilder {

	for _, sort := range paginationParam.Sorts {
		base = base.OrderBy(fmt.Sprintf("%s %s", sort.Col, sort.Order))
	}

	if paginationParam.Start != 0 {
		base = base.Offset(paginationParam.Start)
	}

	if paginationParam.End != 0 {
		base = base.Limit(uint64(paginationParam.End - paginationParam.Start + 1))
	}

	for _, filter := range paginationParam.Filters {
		if strings.ContainsAny(filter.Cond, "%") {
			base = base.Where(sq.Like{filter.Col: filter.Cond})
		} else {
			base = base.Where(sq.Eq{filter.Col: filter.Cond})
		}
	}

	return base
}

func extractColAndOrder(colQueryParam string) (col, order string) {
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

func extractToSort(colQueryParam string) *Sort {
	col, order := extractColAndOrder(colQueryParam)
	if col == "" || order == "" {
		return nil
	}
	return &Sort{Col: col, Order: order}
}

func extractToNextKey(colQueryParam string) *NextKey {
	col, order := extractColAndOrder(colQueryParam)
	if col == "" || order == "" {
		return nil
	}
	return &NextKey{Col: col, Order: order}
}

func BuildPaginationParam(queryParams url.Values, validColumns []string) PaginationParam {
	// log.Warnf("# VALID COLUMNS: %q", validColumns)
	paginationParam := PaginationParam{}

	sorts := []Sort{}
	for _, colSort := range strings.Split(queryParams.Get("_sort"), ",") {
		sort := extractToSort(colSort)
		if sort != nil && funk.ContainsString(validColumns, sort.Col) {
			sorts = append(sorts, Sort{Col: sort.Col, Order: sort.Order})
		}
	}
	paginationParam.Sorts = sorts

	start, _ := strconv.Atoi(queryParams.Get("_start"))
	paginationParam.Start = uint64(start)
	end, _ := strconv.Atoi(queryParams.Get("_end"))
	paginationParam.End = uint64(end)

	filters := map[string]*Filter{}
	for _, col := range validColumns {
		whereCond := queryParams.Get(col)
		if whereCond != "" {
			filters[col] = &Filter{Col: col, Cond: whereCond}
		}
	}
	paginationParam.Filters = filters

	nextKey := extractToNextKey(queryParams.Get("_next_key"))
	if nextKey != nil && funk.ContainsString(validColumns, nextKey.Col) {
		paginationParam.NextKey = nextKey
	}

	listColNextValue := strings.Split(queryParams.Get("_next"), ",")
	if len(listColNextValue) > 0 {
		// pop next key val (at the tail), see:https://github.com/golang/go/wiki/SliceTricks#pop
		nextKeyVal, listColNextValWithoutKey := listColNextValue[len(listColNextValue)-1], listColNextValue[:len(listColNextValue)-1]
		if paginationParam.NextKey != nil {
			paginationParam.NextKey.Value = nextKeyVal
		}

		nexts := []Next{}
		if len(paginationParam.Sorts) == len(listColNextValWithoutKey) {
			for i, sort := range paginationParam.Sorts {
				nextVal := listColNextValWithoutKey[i]
				nexts = append(nexts, Next{Col: sort.Col, Value: nextVal})
			}
		}
		paginationParam.Nexts = nexts
	}

	return paginationParam
}

func GetPaginationType(paginationParam PaginationParam) PaginationType {
	if paginationParam.NextKey != nil {
		return KeysetPagination
	}

	if paginationParam.Start != 0 || paginationParam.End != 0 {
		return OffsetPagination
	}
}
