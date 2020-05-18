import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Table, Divider, Button, Popconfirm, Tooltip, message,
} from 'antd';
import moment from 'moment';
import { fetchModules } from 'api/module';
import useTableFilterProps from 'hooks/useTableFilterProps';
import { buildQueryParam, onTableChange } from 'utils/pagination';
import useTablePaginationTotal from 'hooks/useTablePaginationTotal';
import useTablePaginationNormalizedListData from 'hooks/useTablePaginationNormalizedListData';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';

const defaultPagination = {
  current: 1,
  pageSize: 5,
};

const formatDate = (dateString) => moment(dateString).fromNow();

function ModuleList(props) {
  const {
    listModule, setListModule, onEdit, onDelete,
  } = props;

  const [loading, setLoading] = useState(false);
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const total = useTablePaginationTotal(paginationInfo, listModule);
  const normalizedListData = useTablePaginationNormalizedListData(
    paginationInfo,
    listModule,
  );

  useEffect(() => {
    async function fetchData() {
      setLoading(true);
      try {
        const queryParam = buildQueryParam(
          paginationInfo,
          filteredInfo,
          sortedInfo,
        );
        const modules = await fetchModules({ params: queryParam });
        setListModule(modules);
      } catch (error) {
        message.error(error.message);
      }
      setLoading(false);
    }
    fetchData();
  }, [paginationInfo, filteredInfo, sortedInfo, setListModule]);

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      width: '5%',
      sorter: false,
      sortOrder: sortedInfo.columnKey === 'id' && sortedInfo.order,
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      width: '20%',
      className: 'col-name',
      sorter: true,
      sortOrder: sortedInfo.columnKey === 'name' && sortedInfo.order,
      ...useTableFilterProps('name'),
      render: (text, record) => (
        <div>{record.name}</div>
      ),
    },
    {
      title: 'Path',
      dataIndex: 'path',
      key: 'path',
      sorter: false,
      sortOrder: sortedInfo.columnKey === 'path' && sortedInfo.order,
      render: (text, record) => (
        <div>{record.path}</div>
      ),
    },
    {
      title: 'Pattern',
      dataIndex: 'pattern',
      key: 'pattern',
      sorter: false,
      sortOrder: sortedInfo.columnKey === 'pattern' && sortedInfo.order,
      render: (text, record) => (
        <div>{record.pattern}</div>
      ),
    },
    {
      title: 'Last Updated',
      dataIndex: 'updated_at',
      key: 'updated_at',
      sorter: true,
      sortOrder: sortedInfo.columnKey === 'updated_at' && sortedInfo.order,
      render: (text, record) => <div>{formatDate(record.updated_at)}</div>,
    },
    {
      title: 'Action',
      key: 'action',
      className: 'col-action',
      render: (text, record) => (
        <span data-testid="colgroup-action">
          <Tooltip title="Edit">
            <Button
              data-testid="btn-edit"
              onClick={() => onEdit(record)}
              icon={<EditOutlined />}
            >
              Edit
            </Button>
          </Tooltip>
          <Divider type="vertical" />
          <Popconfirm
            title="Are you sure to delete this module?"
            placement="topRight"
            onConfirm={() => onDelete(record)}
          >
            <Tooltip title="Delete">
              <Button data-testid="btn-delete" type="primary" danger icon={<DeleteOutlined />}>Delete</Button>
            </Tooltip>
          </Popconfirm>
        </span>
      ),
    },
  ];

  return (
    <div>
      <Table
        rowKey="id"
        columns={columns}
        dataSource={normalizedListData}
        pagination={{ ...paginationInfo, total }}
        onChange={onTableChange(
          setPaginationInfo,
          setFilteredInfo,
          setSortedInfo,
        )}
        loading={loading}
        scroll={{ x: true }}
      />
    </div>
  );
}

const moduleData = PropTypes.shape({
  id: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  path: PropTypes.string,
  pattern: PropTypes.string,
});

ModuleList.propTypes = {
  listModule: PropTypes.arrayOf(moduleData).isRequired,
  setListModule: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default ModuleList;
