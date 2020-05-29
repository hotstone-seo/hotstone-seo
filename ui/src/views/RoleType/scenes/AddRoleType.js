import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { RoleTypeForm } from 'components/RoleType';

function AddRoleType() {
  const history = useHistory();

  const handleCreate = (newRole) => {
    if (newRole.name === undefined) {
      message.error('Role already exists');
    } else {
      history.push('/role-type', {
        message: {
          level: 'success',
          content: `${newRole.name} is successfully created`,
        },
      });
    }
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/role-type')}
        title="Add new Role User"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <RoleTypeForm handleSubmit={handleCreate} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddRoleType;
