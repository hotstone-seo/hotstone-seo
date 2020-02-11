import React from 'react';
import { Link } from 'react-router-dom';
import { Table, Button, Divider, Popconfirm } from 'antd';

function RuleList(props) {
  const { rules, onClick, onEdit, onDelete } = props;

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
      )
    },
    { title: 'URL Pattern', dataIndex: 'url_pattern', key: 'urlPattern' },
    { title: 'Last Updated', dataIndex: 'updated_at', key: 'lastUpdated' },
    {
      title: 'Action',
      key: 'action',
      render: (text, record) => (
        <span>
          <Button
            type="link"
            onClick={() => onEdit(record)}
            style={{ padding: 0 }}
          >
            Edit
          </Button>
          <Divider type="vertical" />
          <Popconfirm
            title="Are you sure to delete this rule?"
            placement="topRight"
            onConfirm={() => onDelete(record)}
          >
            <Button type="link" danger style={{ padding: 0 }}>Delete</Button>
          </Popconfirm>
        </span>
      )
    }
  ]

  return (
    <Table
      columns={columns}
      dataSource={rules}
      rowKey="id"
    />
  );
}

export default RuleList;
