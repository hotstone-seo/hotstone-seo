import React from 'react';
import PropTypes from 'prop-types';
import {
  Table, Button, Divider, Popconfirm, Tooltip,
} from 'antd';
import moment from 'moment';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';

const formatDate = (dateString) => moment(dateString).fromNow();

function RoleTypeList(props) {
  const {
    roleTypes, loading, onClick, onEdit, onDelete,
  } = props;

  const columns = [
    {
      title: 'Role Name',
      dataIndex: 'name',
      key: 'name',
      render: (text, record) => (
        <Button
          type="link"
          onClick={() => onClick(record)}
          style={{ padding: 0 }}
        >
          {text}
        </Button>
      ),
    },
    {
      title: 'Last Updated',
      dataIndex: 'updated_at',
      key: 'lastUpdated',
      render: (text, record) => <div>{formatDate(record.updated_at)}</div>,
    },
    {
      title: 'Action',
      key: 'action',
      render: (text, record) => (
        <span>
          <Tooltip title="Edit">
            <Button
              onClick={() => onEdit(record)}
              icon={<EditOutlined />}
            >
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
              <Button type="primary" danger icon={<DeleteOutlined />}>Delete</Button>
            </Tooltip>
          </Popconfirm>
        </span>
      ),
    },
  ];

  return (
    <Table
      columns={columns}
      roleType={roleTypes}
      rowKey="id"
      loading={loading}
      scroll={{ x: true }}
    />
  );
}

RoleTypeList.defaultProps = {
  roleTypes: [],
  loading: false,
};

RoleTypeList.propTypes = {
  roleTypes: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number,
      name: PropTypes.string,
    }),
  ),
  loading: PropTypes.bool,
  onClick: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default RoleTypeList;
