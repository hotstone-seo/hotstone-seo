import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { ClientKeyForm } from 'components/ClientKey';
import { getClientKey, updateClientKey } from 'api/client_key';

function EditClientKey() {
  const { id } = useParams();
  const history = useHistory();

  const [clientKey, setClientKey] = useState({});

  useEffect(() => {
    getClientKey(id)
      .then((newClientKey) => {
        setClientKey(newClientKey);
      })
      .catch((error) => {
        message.error(error.message);
      });
  }, [id]);

  const handleEdit = (newClientKey) => {
    updateClientKey(newClientKey)
      .then(() => {
        history.push('/client-keys');
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/client-keys')}
        title={`Edit ${clientKey.name}`}
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <ClientKeyForm handleSubmit={handleEdit} clientKey={clientKey} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default EditClientKey;
