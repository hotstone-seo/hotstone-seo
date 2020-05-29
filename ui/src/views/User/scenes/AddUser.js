import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { UserForm } from 'components/User';
import { createUser } from 'api/user';
import useRoleTypes from 'hooks/useRoleTypes';

function AddUser() {
  const history = useHistory();
  const [roleTypes] = useRoleTypes();

  const handleCreate = (user) => {
    createUser(user)
      .then((newUser) => {
        if (newUser.email === undefined) {
          message.error(`${user.email} already register`);
        } else {
          history.push('/users', {
            message: {
              level: 'success',
              content: `${newUser.email} is successfully created`,
            },
          });
        }
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/users')}
        title="Add new User"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <UserForm handleSubmit={handleCreate} roleTypes={roleTypes} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddUser;
