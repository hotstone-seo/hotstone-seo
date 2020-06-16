import React, { useState, useEffect } from "react";
import PropTypes from "prop-types";
import { Table, Divider, Button, Popconfirm, Tooltip, message } from "antd";
import { fetchRoleTypes } from "api/roleType";
import useTableFilterProps from "hooks/useTableFilterProps";
import { buildQueryParam, onTableChange } from "utils/pagination";
import useTablePaginationTotal from "hooks/useTablePaginationTotal";
import useTablePaginationNormalizedListData from "hooks/useTablePaginationNormalizedListData";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";

const defaultPagination = {
  current: 1,
  pageSize: 5,
};

function RoleTypeList(props) {
  const { listRoleType, setListRoleType, onEdit, onDelete } = props;

  const [loading, setLoading] = useState(false);
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const total = useTablePaginationTotal(paginationInfo, listRoleType);
  const normalizedListData = useTablePaginationNormalizedListData(
    paginationInfo,
    listRoleType
  );

  useEffect(() => {
    async function fetchData() {
      setLoading(true);
      try {
        const queryParam = buildQueryParam(
          paginationInfo,
          filteredInfo,
          sortedInfo
        );
        const users = await fetchRoleTypes({ params: queryParam });
        setListRoleType(users);
      } catch (error) {
        message.error(error.message);
      }
      setLoading(false);
    }
    fetchData();
  }, [paginationInfo, filteredInfo, sortedInfo, setListRoleType]);

  const columns = [
    {
      title: "Role Name",
      dataIndex: "name",
      key: "name",
      className: "col-name",
      width: "20%",
      sorter: true,
      sortOrder: sortedInfo.columnKey === "name" && sortedInfo.order,
      ...useTableFilterProps("name"),
      render: (text) => <div>{text}</div>,
    },
    {
      title: "Action",
      key: "action",
      render: (text, record) => (
        <span>
          <Tooltip title="Edit">
            <Button onClick={() => onEdit(record)} icon={<EditOutlined />}>
              Edit
            </Button>
          </Tooltip>
          <Divider type="vertical" />
          <Popconfirm
            title="Are you sure to delete this role?"
            placement="topRight"
            onConfirm={() => onDelete(record)}
          >
            <Tooltip title="Delete">
              <Button type="primary" danger icon={<DeleteOutlined />}>
                Delete
              </Button>
            </Tooltip>
          </Popconfirm>
        </span>
      ),
    },
  ];

  return (
    <Table
      rowKey="id"
      columns={columns}
      dataSource={normalizedListData}
      pagination={{ ...paginationInfo, total }}
      onChange={onTableChange(
        setPaginationInfo,
        setFilteredInfo,
        setSortedInfo
      )}
      loading={loading}
      scroll={{ x: true }}
    />
  );
}

const roleType = PropTypes.shape({
  id: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  updated_at: PropTypes.string,
});

RoleTypeList.propTypes = {
  listRoleType: PropTypes.arrayOf(roleType).isRequired,
  setListRoleType: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default RoleTypeList;
