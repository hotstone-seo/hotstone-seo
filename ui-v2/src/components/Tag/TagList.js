import React from 'react';
import PropTypes from 'prop-types';
import ReactDOMServer from 'react-dom/server';
import {
  Table, Button, Popconfirm, Divider, List,
} from 'antd';

function TagList({ tags, onEdit, onDelete }) {
  const columns = [
    { title: 'Type', dataIndex: 'type', key: 'type' },
    {
      title: 'Attributes',
      dataIndex: 'attributes',
      key: 'attributes',
      render: (text, record) => {
        const { attributes } = record;
        const attrs = []
        for (const key in attributes) {
          attrs.push(`${key}="${attributes[key]}"`)
        }
        if (attrs.length === 0) {
          return null;
        }
        return (
          <List
            size="small"
            dataSource={attrs}
            renderItem={(item) => (
              <List.Item>{item}</List.Item>
            )}
          />
        );
      }
    },
    { title: 'Value', dataIndex: 'value', key: 'value' },
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
            title="Are you sure to delete this tag?"
            placement="topRight"
            onConfirm={() => onDelete(record)}
          >
            <Button type="link" danger style={{ padding: 0 }}>Delete</Button>
          </Popconfirm>
        </span>
      ),
    },
  ];

  return (
    <Table
      columns={columns}
      dataSource={tags}
      rowKey="id"
    />
  );
}

TagList.defaultProps = {
  tags: [],
};

TagList.propTypes = {
  tags: PropTypes.arrayOf(
    PropTypes.shape({
      type: PropTypes.string.isRequired,
      attributes: PropTypes.object,
      value: PropTypes.string,
    }),
  ),

  onEdit: PropTypes.func.isRequired,

  onDelete: PropTypes.func.isRequired,
};

export default TagList;
