import { useMemo } from 'react';

const PageSizeMultiplierHack = 2;

// THIS IS A WORKAROUND if 'count' of total data is not available in backend response.
// If 'count' is availabe, this hook (useTablePaginationTotal) is not necessary
export default function useTablePaginationNormalizedListData(pagination, listData) {
  return useMemo(() => {
    let normalizedListData = listData;
    if (listData.length / pagination.pageSize > PageSizeMultiplierHack - 1) {
      normalizedListData = listData.slice(
        0,
        pagination.pageSize - listData.length,
      );
    }

    return normalizedListData;
  }, [pagination, listData]);
}
