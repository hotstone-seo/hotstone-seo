import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { RoleTypeForm } from 'components/RoleType';
import { getRoleType, updateRoleType } from 'api/roleType';

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
        history.push('/roletypes', {
          message: {
            level: 'error',
            content: error.message,
          },
        });
      });
  }, [roleTypeID, history]);

  const handleEdit = (newRoleType) => {
    // TO DO : re-check . Still not get value module access
    updateRoleType(newRoleType)
      .then(() => {
        history.push('/roletypes', {
          message: {
            level: 'success',
            content: `Role ${newRoleType.name} is successfully edit`,
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
