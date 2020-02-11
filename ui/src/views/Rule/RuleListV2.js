import React, { useState, useEffect } from "react";
import { Table } from "antd";
import { useTableFilterProps } from "../../hooks/useTableFilterProps";
import { buildQueryParam, onTableChange } from "../../utils/pagination";
import HotstoneAPI from "../../api/hotstone";
import { useTablePaginationTotal } from "../../hooks/useTablePaginationTotal";

const defaultPagination = {
  current: 1,
  pageSize: 2
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
      width: "10%",
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
      ...useTableFilterProps("name")
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
      dataIndex: "updated_date",
      key: "updated_date",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "updated_date" && sortedInfo.order
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
