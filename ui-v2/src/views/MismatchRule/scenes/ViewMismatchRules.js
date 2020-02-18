import React, { useState, useEffect } from "react";
import { Table, Divider, Button, Popconfirm } from "antd";
import { format, formatDistance } from "date-fns";
import { useTableFilterProps } from "hooks/useTableFilterProps";
import {
  buildQueryParam,
  onTableChange,
  createPageToken
} from "utils/pagination_cursor";
import { fetchMismatched } from "api/metric";
import { useTokenPagination } from "hooks/useTokenPagination";
import _ from "lodash";

const formatDate = since => {
  const sinceDate = new Date(since);

  const full = format(sinceDate, "dd/MM/yyyy - HH:mm");
  const relative = formatDistance(sinceDate, new Date());

  return `${full} (${relative} ago)`;
};

function ViewMismatchRules(props) {
  const { onClick, onEdit, onDelete } = props;

  const [pageSize, setPageSize] = useState(5);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({
    order: "descend",
    field: "since",
    columnKey: "since"
  });

  const [listData, setListData] = useState([]);

  const instTokenPagination = useTokenPagination();
  const {
    setNextPageToken,
    pageToken,
    pageIndex,
    previousPage,
    nextPage,
    canPreviousPage,
    canNextPage,
    resetPagination
  } = instTokenPagination;

  useEffect(() => {
    async function fetchData() {
      try {
        const nextKey = { id: "count", desc: true };
        const queryParam = buildQueryParam(
          pageSize,
          filteredInfo,
          sortedInfo,
          nextKey,
          pageToken
        );

        const listData = await fetchMismatched({ params: queryParam });

        if (!_.isEmpty(listData)) {
          const lastRow = listData[listData.length - 1];

          const nextPageToken = createPageToken(lastRow, sortedInfo, nextKey);
          setNextPageToken(nextPageToken);
        } else {
          previousPage();
        }

        setListData(listData);
      } catch (err) {
        console.log("ERR: ", err);
      }
    }

    fetchData();
  }, [pageSize, filteredInfo, sortedInfo, pageToken]);

  const columns = [
    {
      title: "URL",
      dataIndex: "request_path",
      key: "request_path",
      // width: "30%",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "request_path" && sortedInfo.order,
      ...useTableFilterProps("request_path")
    },

    {
      title: "Since",
      dataIndex: "since",
      key: "since",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "since" && sortedInfo.order,
      render: (text, record) => <div>{formatDate(record.since)}</div>
    },
    {
      title: "Count",
      dataIndex: "count",
      key: "count",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "count" && sortedInfo.order
    }
  ];

  return (
    <div>
      {/* <pre>
        <code>
          {JSON.stringify(
            {
              pageSize,
              sortedInfo,
              filteredInfo,

              ...instTokenPagination
            },
            null,
            2
          )}
        </code>
      </pre> */}
      <Table
        rowKey="request_path"
        columns={columns}
        dataSource={listData}
        pagination={false}
        onChange={onTableChange(setFilteredInfo, setSortedInfo)}
      />
      <div className="ant-pagination">
        <button onClick={() => resetPagination()} disabled={!canPreviousPage}>
          {"Latest"}
        </button>{" "}
        <button onClick={() => previousPage()} disabled={!canPreviousPage}>
          {"Prev"}
        </button>
        {` Page: ${pageIndex + 1} `}
        <button onClick={() => nextPage()} disabled={!canNextPage}>
          {"Next"}
        </button>{" "}
      </div>
    </div>
  );
}

export default ViewMismatchRules;
