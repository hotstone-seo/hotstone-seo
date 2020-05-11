import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { RoleTypeForm } from 'components/RoleType';
import { createRoleType } from 'api/roleType';

function AddRoleType() {
  const history = useHistory();

  const handleCreate = (rule) => {
    createRoleType(rule)
      .then((newRule) => {
        history.push(`/roletypes/${newRule.id}`, {
          message: {
            level: 'success',
            content: `${newRule.name} is successfully created`,
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
        onBack={() => history.push('/roletypes')}
        title="Add new Role User"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <RoleTypeForm onSubmit={handleCreate} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddRoleType;
