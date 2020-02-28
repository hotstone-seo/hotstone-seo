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

export const onTableChange = (
  setPaginationInfo,
  setFilteredInfo,
  setSortedInfo,
) => (pagination, filters, sorter) => {
  setPaginationInfo(pagination);
  setFilteredInfo(filters);
  setSortedInfo(sorter);
};
