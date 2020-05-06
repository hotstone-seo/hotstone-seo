import React from 'react';
import PropTypes from 'prop-types';
import {
  Table, Button, Popconfirm, Divider, List, Tooltip,
} from 'antd';

import { EditOutlined, DeleteOutlined } from '@ant-design/icons';

function TagList(props) {
  const {
    tags, loading, onEdit, onDelete,
  } = props;

  const columns = [
    { title: 'Type', dataIndex: 'type', key: 'type' },
    {
      title: 'Attributes',
      dataIndex: 'attributes',
      key: 'attributes',
      render: (text, record) => {
        const { attributes } = record;
        const attrs = Object.keys(attributes).map((key) => `${key}="${attributes[key]}"`);
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
      },
    },
    { title: 'Value', dataIndex: 'value', key: 'value' },
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
            title="Are you sure to delete this tag?"
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
      dataSource={tags}
      loading={loading}
      rowKey="id"
    />
  );
}

TagList.defaultProps = {
  tags: [],
  loading: false,
};

TagList.propTypes = {
  tags: PropTypes.arrayOf(
    PropTypes.shape({
      type: PropTypes.string.isRequired,
      attributes: PropTypes.object,
      value: PropTypes.string,
    }),
  ),
  loading: PropTypes.bool,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default TagList;
