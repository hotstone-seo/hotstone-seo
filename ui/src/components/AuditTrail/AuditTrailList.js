import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import {
  Table, Divider, Button, Popconfirm, Tooltip, message, Switch, Descriptions,
} from 'antd';
import moment from 'moment';
import { fetchAuditTrails } from 'api/audit_trail';
import useTableFilterProps from 'hooks/useTableFilterProps';
import { buildQueryParam, onTableChange } from 'utils/pagination';
import useTablePaginationTotal from 'hooks/useTablePaginationTotal';
import useTablePaginationNormalizedListData from 'hooks/useTablePaginationNormalizedListData';
import _ from 'lodash';

const defaultPagination = {
  current: 1,
  pageSize: 5,
};

const formatDate = (dateString) => moment(dateString).format('YYYY-MM-DD HH:mm');

function AuditTrailList(props) {
  const { listAuditTrail, setListAuditTrail } = props;

  const [loading, setLoading] = useState(false);
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const total = useTablePaginationTotal(paginationInfo, listAuditTrail);
  const normalizedListData = useTablePaginationNormalizedListData(
    paginationInfo,
    listAuditTrail,
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
        const auditTrails = await fetchAuditTrails({ params: queryParam });
        const updatedListAuditTrail = auditTrails;
        setListAuditTrail(updatedListAuditTrail);
      } catch (error) {
        message.error(error.message);
      }
      setLoading(false);
    }
    fetchData();
  }, [paginationInfo, filteredInfo, sortedInfo]);

  const columns = [
    {
      title: 'Changes Time',
      dataIndex: 'time',
      key: 'time',
      sorter: true,
      sortOrder: sortedInfo.columnKey === 'time' && sortedInfo.order,
      render: (text, record) => <div>{formatDate(record.time)}</div>,
    },
    {
      title: 'Changed By',
      dataIndex: 'username',
      key: 'username',
      ...useTableFilterProps('username'),
    },
    {
      title: 'Entity Name',
      dataIndex: 'entity_name',
      key: 'entity_name',
      ...useTableFilterProps('entity_name'),
    },
    {
      title: 'Entity ID',
      dataIndex: 'entity_id',
      key: 'entity_id',
    },
    {
      title: 'Operation',
      dataIndex: 'operation',
      key: 'operation',
      filters: [
        {
          text: 'INSERT',
          value: 'INSERT',
        },
        {
          text: 'UPDATE',
          value: 'UPDATE',
        },
        {
          text: 'DELETE',
          value: 'DELETE',
        },
      ],
      filterMultiple: false,
    },
    {
      title: 'Data Changes',
      render: (text, record) => <>{JSON.stringify(difference(record.new_data, record.old_data))}</>,
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
        size="small"
        expandable={{
          expandedRowRender: (record) => {
            console.log(record);
            return (
              <Descriptions size="small" column={1} bordered>
                <Descriptions.Item label="Old Data">{JSON.stringify(record.old_data)}</Descriptions.Item>
                <Descriptions.Item label="New Data">{JSON.stringify(record.new_data)}</Descriptions.Item>
              </Descriptions>
            );
          },
        }}
      />
    </div>
  );
}

function difference(object, base) {
  function changes(object, base) {
    return _.transform(object, (result, value, key) => {
      if (!_.isEqual(value, base[key])) {
        result[key] = (_.isObject(value) && _.isObject(base[key])) ? changes(value, base[key]) : value;
      }
    });
  }
  return changes(object, base);
}

// const auditTrailType = PropTypes.shape({
//   id: PropTypes.number.isRequired,
//   name: PropTypes.string.isRequired,
//   url_pattern: PropTypes.string.isRequired,
//   data_source: PropTypes.string,
//   time: PropTypes.instanceOf(Date).isRequired,
// });

// AuditTrailList.propTypes = {
//   listAuditTrail: PropTypes.arrayOf(auditTrailType).isRequired,
//   setListAuditTrail: PropTypes.func.isRequired,
// };

export default AuditTrailList;
