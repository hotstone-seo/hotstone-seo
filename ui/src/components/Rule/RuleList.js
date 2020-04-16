import React from 'react';
import PropTypes from 'prop-types';
import {
  Table, Button, Divider, Popconfirm, Tooltip,
} from 'antd';

import { EditOutlined, DeleteOutlined } from '@ant-design/icons';

function RuleList(props) {
  const {
    rules, onClick, onEdit, onDelete,
  } = props;

  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      render: (text, record) => (
        <Button
          type="link"
          onClick={() => onClick(record)}
        >
          {text}
        </Button>
      ),
    },
    { title: 'URL Pattern', dataIndex: 'url_pattern', key: 'urlPattern' },
    { title: 'Last Updated', dataIndex: 'updated_at', key: 'lastUpdated' },
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
            title="Are you sure to delete this rule?"
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
      dataSource={rules}
      rowKey="id"
    />
  );
}

RuleList.defaultProps = {
  rules: [],
};

RuleList.propTypes = {
  rules: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number,
      name: PropTypes.string,
      url_pattern: PropTypes.string,
      updated_at: PropTypes.string,
    }),
  ),

  onClick: PropTypes.func.isRequired,

  onEdit: PropTypes.func.isRequired,

  onDelete: PropTypes.func.isRequired,
};

export default RuleList;
