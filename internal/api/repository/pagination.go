package repository

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/thoas/go-funk"
)

type PaginationParam struct {
	Sorts   []Sort
	Offset  uint64
	Limit   uint64
	Filters map[string]*Filter
	NextKey *Next
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
	Order string
	Value string
}

type PaginationType int

const (
	NoPagination PaginationType = iota
	OffsetPagination
	KeysetPagination
)

func ComposePagination(base sq.SelectBuilder, paginationParam PaginationParam) sq.SelectBuilder {

	paginationType := GetPaginationType(paginationParam)

	// compose ORDER BY
	for _, sort := range paginationParam.Sorts {
		base = base.OrderBy(fmt.Sprintf("%s %s", sort.Col, sort.Order))
	}

	if paginationType == KeysetPagination {
		base = base.OrderBy(fmt.Sprintf("%s %s", paginationParam.NextKey.Col, paginationParam.NextKey.Order))
	}

	// compose OFFSET & LIMIT
	if paginationType == OffsetPagination && paginationParam.Offset != 0 {
		base = base.Offset(paginationParam.Offset)

	}

	if paginationParam.Limit != 0 {
		base = base.Limit(paginationParam.Limit)
	}

	// compose WHERE
	keysFilter := getSortedFilterKeys(paginationParam.Filters)
	for _, k := range keysFilter {
		filter := paginationParam.Filters[k]
		if strings.ContainsAny(filter.Cond, "%") {
			base = base.Where(sq.Like{filter.Col: filter.Cond})
		} else {
			base = base.Where(sq.Eq{filter.Col: filter.Cond})
		}
	}

	if paginationType == KeysetPagination {
		if len(paginationParam.Nexts) > 0 {
			var firstCond sq.Sqlizer
			var secondCond sq.Sqlizer

			for _, next := range paginationParam.Nexts {
				nextKeyCond := constructWhereCond(next.Order, next.Col, next.Value)
				if firstCond == nil {
					firstCond = nextKeyCond
				} else {
					firstCond = sq.And{firstCond, nextKeyCond}
				}

				nextKeyEqCond := sq.Eq{next.Col: next.Value}
				if secondCond == nil {
					secondCond = nextKeyEqCond
				} else {
					secondCond = sq.And{secondCond, nextKeyEqCond}
				}
			}
			secondCond = sq.And{secondCond, constructWhereCond(paginationParam.NextKey.Order, paginationParam.NextKey.Col, paginationParam.NextKey.Value)}

			base = base.Where(sq.Or{firstCond, secondCond})
		} else if len(paginationParam.Nexts) == 0 && paginationParam.NextKey.Value != "" {
			base = base.Where(constructWhereCond(paginationParam.NextKey.Order, paginationParam.NextKey.Col, paginationParam.NextKey.Value))
		}
	}

	return base
}

func constructWhereCond(order, operand1, operand2 string) sq.Sqlizer {
	switch order {
	case "ASC":
		return sq.Gt{operand1: operand2}
	case "DESC":
		return sq.Lt{operand1: operand2}
	}

	return nil
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

func extractToNextKey(colQueryParam string) *Next {
	col, order := extractColAndOrder(colQueryParam)
	if col == "" || order == "" {
		return nil
	}
	return &Next{Col: col, Order: order}
}

func getSortedFilterKeys(filters map[string]*Filter) []string {
	keys := make([]string, 0)
	for k := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func BuildPaginationParam(queryParams url.Values, validColumns []string) PaginationParam {
	paginationParam := PaginationParam{}

	sorts := []Sort{}
	for _, colSort := range strings.Split(queryParams.Get("_sort"), ",") {
		sort := extractToSort(colSort)
		if sort != nil && funk.ContainsString(validColumns, sort.Col) {
			sorts = append(sorts, Sort{Col: sort.Col, Order: sort.Order})
		}
	}
	paginationParam.Sorts = sorts

	offset, _ := strconv.Atoi(queryParams.Get("_offset"))
	paginationParam.Offset = uint64(offset)
	limit, _ := strconv.Atoi(queryParams.Get("_limit"))
	paginationParam.Limit = uint64(limit)

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
				nexts = append(nexts, Next{Col: sort.Col, Order: sort.Order, Value: nextVal})
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

	if paginationParam.Offset != 0 {
		return OffsetPagination
	}

	return NoPagination
}
