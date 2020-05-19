import React from 'react';
import PropTypes from 'prop-types';
import {
  Table, Button, Divider, Popconfirm, Tooltip,
} from 'antd';
import moment from 'moment';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';

const formatDate = (dateString) => (dateString ? moment(dateString).fromNow() : '');

function ClientKeyList(props) {
  const {
    clientKeys, loading, onEdit, onDelete,
  } = props;

  const columns = [
    { title: 'Name', dataIndex: 'name', key: 'name' },
    { title: 'Key Prefix', dataIndex: 'prefix', key: 'prefix' },
    {
      title: 'Last Used',
      dataIndex: 'last_used_at',
      key: 'lastUsedAt',
      render: (text, record) => <div>{formatDate(record.last_used_at)}</div>,
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
            title="Are you sure to delete this client key?"
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
      dataSource={clientKeys}
      rowKey="id"
      loading={loading}
      scroll={{ x: true }}
    />
  );
}

ClientKeyList.defaultProps = {
  clientKeys: [],
  loading: false,
};

ClientKeyList.propTypes = {
  clientKeys: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number,
      name: PropTypes.string,
      url: PropTypes.string,
      updated_at: PropTypes.string,
    }),
  ),
  loading: PropTypes.bool,
  onClick: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default ClientKeyList;
