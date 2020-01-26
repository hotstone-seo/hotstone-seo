import React from 'react';
import { Table } from 'antd';

const columns = [
  { title: 'Name', dataIndex: 'name', key: 'name' },
  { title: 'URL Pattern', dataIndex: 'url_pattern', key: 'urlPattern' },
  { title: 'Last Updated', dataIndex: 'updated_at', key: 'lastUpdated' },
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
