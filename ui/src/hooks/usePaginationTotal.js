import React, { useMemo, useRef, useState } from "react";

export function usePaginationTotal(pagination, listData) {
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
