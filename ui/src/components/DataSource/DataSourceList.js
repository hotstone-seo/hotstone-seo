import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Table, Button, Divider, Popconfirm, Tooltip, message,
} from 'antd';
import moment from 'moment';
import { fetchDataSourcesPagination } from 'api/datasource';
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

function DataSourceList(props) {
  const {
    dataSources, setDataSources, onClick, onEdit, onDelete,
  } = props;

  const [loading, setLoading] = useState(false);
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const total = useTablePaginationTotal(paginationInfo, dataSources);
  const normalizedListData = useTablePaginationNormalizedListData(
    paginationInfo,
    dataSources,
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
        const dsList = await fetchDataSourcesPagination({ params: queryParam });
        setDataSources(dsList);
      } catch (error) {
        message.error(error.message);
      }
      setLoading(false);
    }
    fetchData();
  }, [paginationInfo, filteredInfo, sortedInfo, setDataSources]);

  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      ...useTableFilterProps('name'),
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
      title: 'URL',
      dataIndex: 'url',
      key: 'url',
      sortOrder: sortedInfo.columnKey === 'url' && sortedInfo.order,
    },
    {
      title: 'Last Updated',
      dataIndex: 'updated_at',
      key: 'lastUpdated',
      sortOrder: sortedInfo.columnKey === 'lastUpdated' && sortedInfo.order,
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
            title="Are you sure to delete this data source?"
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
  );
}

DataSourceList.defaultProps = {
  dataSources: [],
};

DataSourceList.propTypes = {
  dataSources: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number,
      name: PropTypes.string,
      url: PropTypes.string,
      updated_at: PropTypes.string,
    }),
  ),
  onClick: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
  setDataSources: PropTypes.func.isRequired,
};

export default DataSourceList;
