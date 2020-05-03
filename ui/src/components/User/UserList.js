import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Table, Divider, Button, Popconfirm, Tooltip, message, Switch,
} from 'antd';
import moment from 'moment';
import { fetchUsers } from 'api/user';
// TO DO : import { getRoleType } from 'api/roletype';
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

function UserList(props) {
  const {
    listUser, setListUser, onEdit, onDelete,
  } = props;

  const [loading, setLoading] = useState(false);
  const [paginationInfo, setPaginationInfo] = useState(defaultPagination);
  const [filteredInfo, setFilteredInfo] = useState({});
  const [sortedInfo, setSortedInfo] = useState({});

  const total = useTablePaginationTotal(paginationInfo, listUser);
  const normalizedListData = useTablePaginationNormalizedListData(
    paginationInfo,
    listUser,
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
        const users = await fetchUsers({ params: queryParam });
        console.log(users, 'user');
        const updatedListUser = await Promise.all(
          users.map(async (user) => {
            const modifiedUser = user;
            if (!_.isEmpty(user.role_type_id)) {
              // TO DO : get role type from API
              const roleType = 'super admin';
              modifiedUser.role_type = roleType;
            }
            return modifiedUser;
          }),
        );
        console.log(updatedListUser, 'updatedListUser');
        setListUser(updatedListUser);
      } catch (error) {
        message.error(error.message);
      }
      setLoading(false);
    }
    fetchData();
  }, [paginationInfo, filteredInfo, sortedInfo, setListUser]);

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
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
      width: '20%',
      className: 'col-name',
      sorter: true,
      sortOrder: sortedInfo.columnKey === 'email' && sortedInfo.order,
      ...useTableFilterProps('email'),
      render: (text, record) => (
        <div>{record.email}</div>
      ),
    },
    {
      title: 'Role Type',
      dataIndex: 'roleType',
      key: 'role_type',
      sorter: false,
      sortOrder: sortedInfo.columnKey === 'role_type' && sortedInfo.order,
      render: () => (
        <div />
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
            title="Are you sure to delete this user?"
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

const roleType = PropTypes.shape({
  id: PropTypes.number.isRequired,
  email: PropTypes.string.isRequired,
  role_type: PropTypes.string,
  updated_at: PropTypes.string,
});

UserList.propTypes = {
  listUser: PropTypes.arrayOf(roleType).isRequired,
  setListUser: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default UserList;
