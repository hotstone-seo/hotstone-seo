import client from './client';

export function fetchStructuredDatas(cfg = {}) {
  return client.get('/structured-data', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getStructuredData(id) {
  return client.get(`/structured-data/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createStructuredData(values) {
  return client.post('/structured-data', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateStructuredData(values) {
  return client.put(`/structured-data/${values.id}`, values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteStructuredData(id) {
  return client.delete(`/structured-data/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}
