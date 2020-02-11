import { useMemo } from "react";

export function useTablePaginationTotal(pagination, listData) {
  return useMemo(() => {
    let total = 0;
    if (listData.length >= pagination.pageSize) {
      total = pagination.current * pagination.pageSize + pagination.pageSize;
    } else {
      total = pagination.current * pagination.pageSize;
    }

    return total;
  }, [pagination, listData]);
}
