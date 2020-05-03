import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { deleteUser } from 'api/user';
import { UserList } from 'components/User';
import { PlusOutlined } from '@ant-design/icons';

function ViewUsers({ match }) {
  const history = useHistory();
  const [listUser, setListUser] = useState([]);

  const showEditForm = (user) => {
    history.push(`${match.url}/${user.id}`);
  };

  const handleDelete = (user) => {
    deleteUser(user.id)
      .then(() => {
        message.success(`Successfully deleted ${user.email}`);
        setListUser(listUser.filter((item) => item.id !== user.id));
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const addDataUser = () => {
    history.push(`${match.url}/new`);
  };

  return (
    <div>
      <PageHeader
        title="Users"
        subTitle="List of Hotstone User with Role"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Button
          data-testid="btn-new-rule"
          type="primary"
          style={{ marginBottom: 16 }}
          icon={<PlusOutlined />}
          onClick={() => addDataUser()}
        >
          Add New User
        </Button>
        <UserList
          onClick={showEditForm}
          onEdit={showEditForm}
          onDelete={handleDelete}
          listUser={listUser}
          setListUser={setListUser}
        />
      </div>
    </div>
  );
}

ViewUsers.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ViewUsers;
