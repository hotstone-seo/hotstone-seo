import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import {
  Modal, PageHeader, Row, Col, message,
} from 'antd';
import { ClientKeyForm } from 'components/ClientKey';
import { createClientKey } from 'api/client_key';

function AddClientKey() {
  const history = useHistory();

  const handleCreate = (dataSource) => {
    createClientKey(dataSource)
      .then((newClientKey) => {
        Modal.success({
          title: 'New Client Key',
          width: '60%',
          content: (
            <div>
              <p>
                A key for
                {' '}
                {newClientKey.name}
                {' '}
                is successfully created.
                Please store it somewhere safe.
              </p>
              <p><strong>It will only be displayed now</strong></p>
              <code>{ `${newClientKey.prefix}.${newClientKey.key}` }</code>
            </div>
          ),
          onOk() {
            history.push('/client-keys');
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
        onBack={() => history.push('/client-keys')}
        title="Add new Client Key"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <ClientKeyForm handleSubmit={handleCreate} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddClientKey;
