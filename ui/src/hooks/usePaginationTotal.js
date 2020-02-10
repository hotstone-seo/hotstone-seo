import React, { useMemo, useRef, useState } from "react";

export function usePaginationTotal(pagination, listData) {
  const [totalData, setTotalData] = useState(
    pagination.current * pagination.pageSize + pagination.pageSize
  );

  return useMemo(() => {
    // const pagination = { ...paginationInfo };
    let total = totalData;
    if (listData.length >= pagination.pageSize) {
      total = pagination.current * pagination.pageSize + pagination.pageSize;
    } else {
      total = pagination.current * pagination.pageSize;
    }
    setTotalData(total);
  }, [listData]);
}
