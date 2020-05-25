import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col,
} from 'antd';
import { RoleTypeForm } from 'components/RoleType';
import { getRoleType } from 'api/roleType';

function EditRoleType() {
  const { id } = useParams();
  const roleTypeID = parseInt(id, 10);
  const history = useHistory();

  const [roleType, setRoleType] = useState({});

  useEffect(() => {
    getRoleType(roleTypeID)
      .then((newRoleType) => {
        setRoleType(newRoleType);
      })
      .catch((error) => {
        history.push('/role-type', {
          message: {
            level: 'error',
            content: error.message,
          },
        });
      });
  }, [roleTypeID, history]);

  const handleEdit = (newRoleType) => {
    history.push('/role-type', {
      message: {
        level: 'success',
        content: `Role ${roleType.name} is successfully edit`,
      },
    });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/role-type')}
        title="Edit Role"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <RoleTypeForm handleSubmit={handleEdit} roleType={roleType} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default EditRoleType;
