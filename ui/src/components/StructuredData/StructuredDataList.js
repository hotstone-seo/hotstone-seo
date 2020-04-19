import React from 'react';
import PropTypes from 'prop-types';
import {
  Table, Button, Tooltip, Divider, Popconfirm,
} from 'antd';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import moment from 'moment';

const formatDate = (dateString) => moment(dateString).fromNow();

function StructuredDataList(props) {
  const {
    structuredDatas, onEdit, onDelete,
  } = props;

  const columns = [
    { title: 'Type', dataIndex: 'type', key: 'type' },
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
            title="Are you sure to delete this structured data?"
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
      dataSource={structuredDatas}
      rowKey="id"
    />
  );
}

StructuredDataList.defaultProps = {
  structuredDatas: [],
};

StructuredDataList.propTypes = {
  structuredDatas: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number,
      type: PropTypes.string,
      updated_at: PropTypes.string,
    }),
  ),
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default StructuredDataList;
