import _ from "lodash";

export const buildQueryParam = (pagination, filters, sorters) => {
  var queryParam = {};

  const order = sorters["order"];
  if (order !== undefined) {
    const orderSign = order === "descend" ? "-" : "";
    queryParam["_sort"] = `${orderSign}${sorters.field}`;
  }

  Object.entries(filters).forEach(([key, value]) => {
    queryParam[key] = `%${value[0]}%`;
  });

  if (!_.isEmpty(pagination)) {
    queryParam["_start"] = (pagination.current - 1) * pagination.pageSize;
    queryParam["_end"] = pagination.current * pagination.pageSize - 1;
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
