import _ from 'lodash';

const PageSizeMultiplierHack = 2;

export const buildQueryParam = (pagination = {}, filters = {}, sorters = {}) => {
  const queryParam = {};

  if (!_.isEmpty(sorters)) {
    const { order } = sorters;
    if (!_.isEmpty(order)) {
      const orderSign = order === 'descend' ? '-' : '';
      queryParam._sort = `${orderSign}${sorters.field}`;
    }
  }

  if (!_.isEmpty(filters)) {
    Object.entries(filters).forEach(([key, value]) => {
      if (!_.isEmpty(value)) {
        queryParam[key] = `%${value}%`;
      }
    });
  }

  if (!_.isEmpty(pagination)) {
    queryParam._offset = (pagination.current - 1) * pagination.pageSize;
    // THIS IS A WORKAROUND if 'count' of total data is not available in backend response.
    // If 'count' is available, limit = pageSize
    queryParam._limit = pagination.pageSize * PageSizeMultiplierHack;
  }

  return queryParam;
};

export const buildPagination = (queryParam = {}) => {
  const pagination = {};
  const filter = {};
  const sort = {};

  const {
    _limit, _offset, _sort, ...filters
  } = queryParam;

  if (_limit && _offset) {
    pagination.pageSize = _limit / PageSizeMultiplierHack;
    pagination.current = (_offset / pagination.pageSize) + 1;
    pagination.total = pagination.current * pagination.pageSize;
  }

  Object.entries(filters).forEach(([key, value]) => {
    const [valWithoutPercent] = value.split('%').filter((x) => x);
    filter[key] = valWithoutPercent;
  });

  if (_sort) {
    if (_sort[0] === '-') {
      sort.field = _sort.slice(1);
      sort.order = 'descend';
    } else {
      sort.field = _sort;
      sort.order = 'ascend';
    }
  }

  return { pagination, filter, sort };
};

export const onTableChange = (
  setPaginationInfo,
  setFilteredInfo,
  setSortedInfo,
) => (pagination, filters, sorter) => {
  setPaginationInfo(pagination);
  setFilteredInfo(filters);
  setSortedInfo(sorter);
};
