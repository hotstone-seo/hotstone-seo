import React, { useState, useEffect } from 'react';
import {
  Table, Divider, Button, Popconfirm,
} from 'antd';
import { format } from 'date-fns';
import { fetchRules } from 'api/rule';
import { getDataSource } from 'api/datasource';
import { useTableFilterProps } from '../../hooks/useTableFilterProps';
import { buildQueryParam, onTableChange } from '../../utils/pagination';
import { useTablePaginationTotal } from '../../hooks/useTablePaginationTotal';
import { useTablePaginationNormalizedListData } from '../../hooks/useTablePaginationNormalizedListData';

const defaultPagination = {
  current: 1,
  pageSize: 3,
};

const formatDate = (since) => {
  const sinceDate = new Date(since);

  const full = format(sinceDate, 'dd/MM/yyyy - HH:mm');

  return `${full}`;
};

function RuleListV2(props) {
  const { onClick, onEdit, onDelete } = props;
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const [listRule, setListRule] = useState([]);

  const total = useTablePaginationTotal(paginationInfo, listRule);
  const normalizedListData = useTablePaginationNormalizedListData(
    paginationInfo,
    listRule,
  );

  useEffect(() => {
    async function fetchData() {
      try {
        const queryParam = buildQueryParam(
          paginationInfo,
          filteredInfo,
          sortedInfo,
        );

        const rules = await fetchRules({ params: queryParam });
        const updatedListRule = await Promise.all(
          rules.map(async (rule) => {
            if (rule.data_source_id == null) {
              rule.data_source = '';
            } else {
              const dataSource = await getDataSource(rule.data_source_id);
              rule.data_source = dataSource.name;
            }
            return rule;
          }),
        );

        setListRule(updatedListRule);
      } catch (err) {
        console.log('ERR: ', err);
      }
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
        <Button type="link" onClick={() => onClick(record)}>
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
      dataIndex: 'data_source',
      key: 'data_source',
      sorter: false,
      sortOrder: sortedInfo.columnKey === 'data_source' && sortedInfo.order,
    },
    {
      title: 'Updated Date',
      dataIndex: 'updated_at',
      key: 'updated_at',
      sorter: true,
      sortOrder: sortedInfo.columnKey === 'updated_at' && sortedInfo.order,
      render: (text, record) => <div>{formatDate(record.updated_at)}</div>,
    },
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
            <Button type="link" danger style={{ padding: 0 }}>
              Delete
            </Button>
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
      />
    </div>
  );
}

export default RuleListV2;
