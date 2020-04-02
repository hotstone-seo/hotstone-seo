import _ from 'lodash';

export const buildQueryParam = (
  pageSize,
  filters,
  sorters,
  nextKey,
  pageToken,
) => {
  const queryParam = {};

  const { order } = sorters;
  if (!_.isEmpty(order)) {
    const orderSign = order === 'descend' ? '-' : '';
    queryParam._sort = `${orderSign}${sorters.field}`;
  }

  Object.entries(filters).forEach(([key, value]) => {
    if (!_.isEmpty(value)) {
      queryParam[key] = `%${value[0]}%`;
    }
  });

  queryParam._limit = pageSize;

  if (!_.isEmpty(nextKey)) {
    const orderSign = nextKey.desc ? '-' : '';
    queryParam._next_key = `${orderSign}${nextKey.id}`;
  }

  if (!_.isEmpty(pageToken)) {
    pageToken.map(({ id, desc, val }) => {
      let { _next } = queryParam;
      _next = `${_next === undefined ? '' : `${_next},`}${val}`;
      queryParam._next = _next;
    });
  }

  return queryParam;
};

export const createPageToken = (lastRow, sortBy, nextKey) => {
  const pageToken = [];

  const { order } = sortBy;
  if (!_.isEmpty(order)) {
    const desc = order === 'descend';
    pageToken.push({
      id: sortBy.field,
      desc,
      val: lastRow[sortBy.field],
    });
  }

  pageToken.push({ ...nextKey, val: lastRow[nextKey.id] });
  return pageToken;
};

export const onTableChange = (setFilteredInfo, setSortedInfo) => (pagination, filters, sorter) => {
  setFilteredInfo(filters);
  setSortedInfo(sorter);
};
