import React from 'react';
import { Table, Button, Divider } from 'antd';

const columns = [
  { title: 'Name', dataIndex: 'name', key: 'name' },
  { title: 'URL Pattern', dataIndex: 'url_pattern', key: 'urlPattern' },
  { title: 'Last Updated', dataIndex: 'updated_at', key: 'lastUpdated' },
  {
    title: 'Action',
    key: 'action',
    render: (text, record) => (
      <span>
        <Button type="link" style={{ padding: 0 }}>Edit</Button>
        <Divider type="vertical" />
        <Button type="link" danger style={{ padding: 0 }}>Delete</Button>
      </span>
    )
  }
]

function RuleList({ rules }) {
  return (
    <Table
      columns={columns}
      dataSource={rules}
      rowKey="id"
    />
  );
}

export default RuleList;
