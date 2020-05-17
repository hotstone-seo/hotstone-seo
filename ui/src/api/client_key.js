import client from './client';

export function fetchClientKeys() {
  return client.get('/client-keys')
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getClientKey(id) {
  return client.get(`/client-keys/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createClientKey(clientKey) {
  return client.post('/client-keys', clientKey)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateClientKey(clientKey) {
  return client.put('/client-keys', clientKey)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteClientKey(id) {
  return client.delete(`/client-keys/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const ClientKeyAPI = {
  fetch: fetchClientKeys,
  get: getClientKey,
  create: createClientKey,
  update: updateClientKey,
  delete: deleteClientKey,
};

export default ClientKeyAPI;
