import React from 'react';
import PropTypes from 'prop-types';
import { useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { fetchClientKeys, deleteClientKey } from 'api/client_key';
import { fetchClientKeyLastUsed } from 'api/metric';
import useAsync from 'hooks/useAsync';
import { ClientKeyList } from 'components/ClientKey';

import { PlusOutlined } from '@ant-design/icons';

const fetchClientKeysAndLastTimeUsed = async () => {
  const clientKeys = await fetchClientKeys();
  return Promise.all(
    clientKeys.map(async (clientKey) => {
      const { time } = await fetchClientKeyLastUsed({ params: { client_key_id: clientKey.id } });
      return { ...clientKey, last_used_at: time };
    }),
  );
};

export default function ViewClientKeys({ match }) {
  const history = useHistory();
  const {
    pending, value: clientKeys, setValue: setClientKeys, error,
  } = useAsync(fetchClientKeysAndLastTimeUsed);

  if (error) {
    message.error(error.message);
  }

  const showEditScene = (clientKey) => {
    history.push(`${match.url}/${clientKey.id}`);
  };

  const removeClientKey = (clientKey) => {
    deleteClientKey(clientKey.id)
      .then(() => {
        message.success(`Successfully deleted ${clientKey.name}`);
        setClientKeys(
          clientKeys.filter((item) => item.id !== clientKey.id),
        );
      })
      .catch((err) => {
        message.error(err.message);
      });
  };

  const addClientKey = () => {
    history.push(`${match.url}/new`);
  };

  return (
    <div>
      <PageHeader
        title="Client Keys"
        subTitle="Manage keys for HotStone Client"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Button
          type="primary"
          style={{ marginBottom: 16 }}
          icon={<PlusOutlined />}
          onClick={() => addClientKey()}
        >
          Add New Client Key
        </Button>
        <ClientKeyList
          clientKeys={clientKeys}
          loading={pending}
          onClick={showEditScene}
          onEdit={showEditScene}
          onDelete={removeClientKey}
        />
      </div>
    </div>
  );
}

ViewClientKeys.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};
