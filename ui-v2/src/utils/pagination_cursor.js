import _ from "lodash";

export const buildQueryParam = (
  pageSize,
  filters,
  sorters,
  nextKey,
  pageToken
) => {
  var queryParam = {};

  const order = sorters["order"];
  if (!_.isEmpty(order)) {
    const orderSign = order === "descend" ? "-" : "";
    queryParam["_sort"] = `${orderSign}${sorters.field}`;
  }

  Object.entries(filters).forEach(([key, value]) => {
    if (!_.isEmpty(value)) {
      queryParam[key] = `%${value[0]}%`;
    }
  });

  queryParam["_limit"] = pageSize;

  if (!_.isEmpty(nextKey)) {
    const orderSign = nextKey.desc ? "-" : "";
    queryParam["_next_key"] = `${orderSign}${nextKey.id}`;
  }

  if (!_.isEmpty(pageToken)) {
    pageToken.map(({ id, desc, val }) => {
      var _next = queryParam["_next"];
      _next = (_next == undefined ? "" : `${_next},`) + `${val}`;
      queryParam["_next"] = _next;
    });
  }

  return queryParam;
};

export const createPageToken = (lastRow, sortBy, nextKey) => {
  var pageToken = [];

  const order = sortBy["order"];
  if (!_.isEmpty(order)) {
    const desc = order === "descend";
    pageToken.push({
      id: sortBy.field,
      desc: desc,
      val: lastRow[sortBy.field]
    });
  }

  pageToken.push({ ...nextKey, val: lastRow[nextKey.id] });
  return pageToken;
};

export const onTableChange = (setFilteredInfo, setSortedInfo) => {
  return (pagination, filters, sorter) => {
    setFilteredInfo(filters);
    setSortedInfo(sorter);
  };
};
