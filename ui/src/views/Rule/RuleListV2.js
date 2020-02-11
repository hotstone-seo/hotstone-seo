import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { Table, Divider, Button } from "antd";
import { format, formatDistance } from "date-fns";
import { useTableFilterProps } from "../../hooks/useTableFilterProps";
import { buildQueryParam, onTableChange } from "../../utils/pagination";
import HotstoneAPI from "../../api/hotstone";
import { useTablePaginationTotal } from "../../hooks/useTablePaginationTotal";

const defaultPagination = {
  current: 1,
  pageSize: 2
};

const formatDate = since => {
  const sinceDate = new Date(since);

  const full = format(sinceDate, "dd/MM/yyyy - HH:mm");
  const relative = formatDistance(sinceDate, new Date());

  return `${full}`;
};

function RuleListV2() {
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const [listRule, setListRule] = useState([]);

  const total = useTablePaginationTotal(paginationInfo, listRule);

  useEffect(() => {
    async function fetchThenNormalizeListRule() {
      try {
        const queryParam = buildQueryParam(
          paginationInfo,
          filteredInfo,
          sortedInfo
        );

        var rules = await HotstoneAPI.getRules({ params: queryParam });
        const updatedListRule = await Promise.all(
          rules.map(async rule => {
            if (rule.data_source_id == null) {
              rule["data_source"] = "";
            } else {
              const dataSource = await HotstoneAPI.getDataSource(
                rule.data_source_id
              );
              rule["data_source"] = dataSource.name;
            }
            return rule;
          })
        );

        setListRule(updatedListRule);
      } catch (err) {
        console.log("ERR: ", err);
      }
    }

    fetchThenNormalizeListRule();
  }, [paginationInfo, filteredInfo, sortedInfo]);

  const columns = [
    {
      title: "ID",
      dataIndex: "id",
      key: "id",
      width: "5%",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "id" && sortedInfo.order
    },
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      width: "20%",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "name" && sortedInfo.order,
      ...useTableFilterProps("name"),
      render: (text, record) => (
        <Link to={`/rule-detail/?id=${record.id}`}>{text}</Link>
      )
    },
    {
      title: "URL Pattern",
      dataIndex: "url_pattern",
      key: "url_pattern",
      width: "30%",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "url_pattern" && sortedInfo.order,
      ...useTableFilterProps("url_pattern")
    },
    {
      title: "Data Source",
      dataIndex: "data_source",
      key: "data_source",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "data_source" && sortedInfo.order
    },
    {
      title: "Updated Date",
      dataIndex: "updated_at",
      key: "updated_at",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "updated_at" && sortedInfo.order,
      render: (text, record) => <div>{formatDate(record.updated_at)}</div>
    },
    {
      title: "Action",
      key: "action",
      render: (text, record) => (
        <span>
          <Button>Edit</Button>
          <Divider type="vertical" />
          <Button type="danger">Delete</Button>
        </span>
      )
    }
  ];

  return (
    <div>
      <Table
        rowKey="id"
        columns={columns}
        dataSource={listRule}
        pagination={{ ...paginationInfo, total: total }}
        onChange={onTableChange(
          setPaginationInfo,
          setFilteredInfo,
          setSortedInfo
        )}
      />
    </div>
  );
}

export default RuleListV2;
