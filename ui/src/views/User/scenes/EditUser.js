import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
// import { EditOutlined, BarChartOutlined } from '@ant-design/icons';
import { UserForm } from 'components/User';
import { getUser, updateUser } from 'api/user';
import useRoleTypes from 'hooks/useRoleTypes';


function EditUser() {
  const { id } = useParams();
  const userID = parseInt(id, 10);
  const history = useHistory();
  const [roleTypes] = useRoleTypes();

  const [user, setUser] = useState({});

  useEffect(() => {
    getUser(userID)
      .then((newUser) => {
        setUser(newUser);
      })
      .catch((error) => {
        history.push('/users', {
          message: {
            level: 'error',
            content: error.message,
          },
        });
      });
  }, [userID, history]);

  const handleEdit = (newUser) => {
    updateUser(newUser)
      .then(() => {
        history.push('/users', {
          message: {
            level: 'success',
            content: `Role user ${newUser.email} is successfully edit`,
          },
        });
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/users')}
        title={`Edit Role ${user.email}`}
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <UserForm handleSubmit={handleEdit} user={user} roleTypes={roleTypes} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default EditUser;
