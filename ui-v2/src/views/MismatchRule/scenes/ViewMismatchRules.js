import React, { useState, useEffect } from "react";
import { Table, Divider, Button, Popconfirm, Row, Col } from "antd";
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
    resetPagination();
  }, [filteredInfo]);

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
          setListData(listData);
        } else {
          previousPage();
        }
      } catch (err) {
        console.log("ERR: ", err);
      }
    }

    fetchData();
  }, [pageSize, filteredInfo, sortedInfo, pageToken]);

  const columns = [
    {
      title: "URL",
      dataIndex: "url",
      key: "url",
      // width: "30%",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "url" && sortedInfo.order,
      ...useTableFilterProps("url")
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
      <pre>
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
      </pre>
      <Row>
        <Col span={24}>
          <Table
            rowKey="url"
            columns={columns}
            dataSource={listData}
            pagination={false}
            onChange={onTableChange(setFilteredInfo, setSortedInfo)}
          />
        </Col>
      </Row>
      <Row justify="end">
        <Col>
          <Button
            type="primary"
            onClick={() => resetPagination()}
            disabled={!canPreviousPage}
          >
            {"Latest"}
          </Button>
          <Button onClick={() => previousPage()} disabled={!canPreviousPage}>
            {"Prev"}
          </Button>
          {` Page: ${pageIndex + 1} `}
          <Button onClick={() => nextPage()} disabled={!canNextPage}>
            {"Next"}
          </Button>
        </Col>
      </Row>
    </div>
  );
}

export default ViewMismatchRules;
