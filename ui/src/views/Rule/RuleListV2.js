import React, { useState, useEffect } from "react";
import { Table } from "antd";
import { getColumnSearchProps } from "../../utils/pagination";
import { useFilterProps } from "../../hooks/useFilterProps";
import HotstoneAPI from "../../api/hotstone";

const data = [
  {
    key: "1",
    name: "John Brown",
    age: 32,
    address: "New York No. 1 Lake Park"
  },
  {
    key: "2",
    name: "Joe Black",
    age: 42,
    address: "London No. 1 Lake Park"
  },
  {
    key: "3",
    name: "Jim Green",
    age: 32,
    address: "Sidney No. 1 Lake Park"
  },
  {
    key: "4",
    name: "Jim Red",
    age: 32,
    address: "London No. 2 Lake Park"
  }
];

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
    // fetchData();
  });

  const columns2 = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      width: "30%",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "name" && sortedInfo.order,
      // ...getColumnSearchProps("name")
      ...useFilterProps("name")
    },
    {
      title: "Age",
      dataIndex: "age",
      key: "age",
      width: "20%",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "age" && sortedInfo.order,
      // ...getColumnSearchProps("age")
      ...getColumnSearchProps("age")
    },
    {
      title: "Address",
      dataIndex: "address",
      key: "address",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "address" && sortedInfo.order,
      // ...getColumnSearchProps("address")
      ...useFilterProps("address")
    }
  ];

  return (
    <div>
      <Table columns={columns2} dataSource={data} onChange={handleChange} />
    </div>
  );
}

export default RuleListV2;
