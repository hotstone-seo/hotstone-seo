import _ from "lodash";

const PageSizeMultiplierHack = 2;

export const buildQueryParam = (pagination, filters, sorters) => {
  var queryParam = {};

  const order = sorters["order"];
  if (order !== undefined) {
    const orderSign = order === "descend" ? "-" : "";
    queryParam["_sort"] = `${orderSign}${sorters.field}`;
  }

  Object.entries(filters).forEach(([key, value]) => {
    if (value[0] !== undefined) {
      queryParam[key] = `%${value[0]}%`;
    }
  });

  if (!_.isEmpty(pagination)) {
    queryParam["_offset"] = (pagination.current - 1) * pagination.pageSize;
    queryParam["_limit"] = pagination.pageSize * PageSizeMultiplierHack; // THIS IS A WORKAROUND if 'count' of total data is not available in backend response. If 'count' is available, limit = pageSize
  }

  return queryParam;
};

export const onTableChange = (
  setPaginationInfo,
  setFilteredInfo,
  setSortedInfo
) => {
  return (pagination, filters, sorter) => {
    setPaginationInfo(pagination);
    setFilteredInfo(filters);
    setSortedInfo(sorter);
  };
};
