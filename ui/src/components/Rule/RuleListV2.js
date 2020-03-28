import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import {
  Table, Divider, Button, Popconfirm, Tooltip, message, Switch,
} from 'antd';
import moment from 'moment';
import { fetchRules } from 'api/rule';
import { getDataSource } from 'api/datasource';
import useTableFilterProps from 'hooks/useTableFilterProps';
import { buildQueryParam, onTableChange } from 'utils/pagination';
import useTablePaginationTotal from 'hooks/useTablePaginationTotal';
import useTablePaginationNormalizedListData from 'hooks/useTablePaginationNormalizedListData';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import _ from 'lodash';

const defaultPagination = {
  current: 1,
  pageSize: 5,
};

const formatDate = (dateString) => moment(dateString).fromNow();

function RuleListV2(props) {
  const {
    listRule, setListRule, onClick, onEdit, onDelete, onChangeToggleButton,
  } = props;

  const [loading, setLoading] = useState(false);
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const total = useTablePaginationTotal(paginationInfo, listRule);
  const normalizedListData = useTablePaginationNormalizedListData(
    paginationInfo,
    listRule,
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
        const rules = await fetchRules({ params: queryParam });
        const updatedListRule = await Promise.all(
          rules.map(async (rule) => {
            const modifiedRule = rule;
            if (!_.isEmpty(rule.data_source_id)) {
              const dataSource = await getDataSource(rule.data_source_id);
              modifiedRule.dataSource = dataSource;
            }
            return modifiedRule;
          }),
        );
        setListRule(updatedListRule);
      } catch (error) {
        message.error(error.message);
      }
      setLoading(false);
    }
    fetchData();
  }, [paginationInfo, filteredInfo, sortedInfo]);

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
      sorter: true,
      sortOrder: sortedInfo.columnKey === 'name' && sortedInfo.order,
      ...useTableFilterProps('name'),
      render: (text, record) => (
        <Button data-testid="btn-detail" type="link" onClick={() => onClick(record)}>
          {text}
        </Button>
      ),
    },
    {
      title: 'URL Pattern',
      dataIndex: 'url_pattern',
      key: 'url_pattern',
      width: '30%',
      sorter: true,
      sortOrder: sortedInfo.columnKey === 'url_pattern' && sortedInfo.order,
      ...useTableFilterProps('url_pattern'),
    },
    {
      title: 'Data Source',
      dataIndex: 'dataSource',
      key: 'data_source',
      sorter: false,
      sortOrder: sortedInfo.columnKey === 'data_source' && sortedInfo.order,
      render: (dataSource) => {
        if (dataSource) {
          return <Link to={`/datasources/${dataSource.id}`}>{dataSource.name}</Link>;
        }
        return null;
      },
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
      title: 'Status',
      key: 'status_start',
      render: (e, record) => (
        <Switch onChange={(value) => onChangeToggleButton(value, record)} defaultChecked={e} checkedChildren="Start" unCheckedChildren="Stop" />
      ),
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
            title="Are you sure to delete this rule?"
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
      />
    </div>
  );
}

const ruleType = PropTypes.shape({
  id: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  url_pattern: PropTypes.string.isRequired,
  data_source: PropTypes.string,
  updated_at: PropTypes.instanceOf(Date).isRequired,
});

RuleListV2.propTypes = {
  listRule: PropTypes.arrayOf(ruleType).isRequired,
  setListRule: PropTypes.func.isRequired,
  onClick: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
  onChangeToggleButton: PropTypes.func.isRequired,
};

export default RuleListV2;
