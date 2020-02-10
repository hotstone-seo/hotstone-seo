import React, { useState, useEffect } from "react";
import { Table } from "antd";
import { useFilterProps } from "../../hooks/useFilterProps";
import HotstoneAPI from "../../api/hotstone";

function RuleListV2() {
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const [listRule, setListRule] = useState([]);

  const handleChange = (pagination, filters, sorter) => {
    console.log("Various parameters", pagination, filters, sorter);
    setFilteredInfo(filters);
    setSortedInfo(sorter);
  };

  useEffect(() => {
    async function fetchThenNormalizeListRule() {
      try {
        var rules = await HotstoneAPI.getRules();
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

        console.log("### RULES: ", rules);
      } catch (err) {
        console.log("ERR: ", err);
      }
    }

    fetchThenNormalizeListRule();
  }, [filteredInfo, sortedInfo]);

  const columns = [
    {
      title: "ID",
      dataIndex: "id",
      key: "id",
      width: "10%",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "id" && sortedInfo.order
      // ...useFilterProps("id")
    },
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      width: "20%",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "name" && sortedInfo.order,
      ...useFilterProps("name")
    },
    {
      title: "URL Pattern",
      dataIndex: "url_pattern",
      key: "url_pattern",
      width: "30%",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "url_pattern" && sortedInfo.order,
      ...useFilterProps("url_pattern")
    },
    {
      title: "Data Source",
      dataIndex: "data_source",
      key: "data_source",
      sorter: false,
      sortOrder: sortedInfo.columnKey === "data_source" && sortedInfo.order
      // ...useFilterProps("data_source")
    },
    {
      title: "Updated Date",
      dataIndex: "updated_date",
      key: "updated_date",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "updated_date" && sortedInfo.order
      // ...useFilterProps("updated_date")
    }
  ];

  return (
    <div>
      <Table
        rowKey="id"
        columns={columns}
        dataSource={listRule}
        onChange={handleChange}
      />
    </div>
  );
}

export default RuleListV2;
