import { useMemo } from "react";

const PageSizeMultiplierHack = 2;

// THIS IS A WORKAROUND if 'count' of total data is not available in backend response.
// If 'count' is availabe, this hook (useTablePaginationTotal) is not necessary
export function useTablePaginationTotal(pagination, listData) {
  return useMemo(() => {
    let total = 0;

    if (listData.length / pagination.pageSize > PageSizeMultiplierHack - 1) {
      total = pagination.current * pagination.pageSize + pagination.pageSize;
    } else {
      total = pagination.current * pagination.pageSize;
    }

    return total;
  }, [pagination, listData]);
}
