import client from './client';

export function fetchDataSources() {
  return client.get('/data_sources')
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getDataSource(id) {
  return client.get(`/data_sources/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createDataSource(dataSource) {
  return client.post('/data_sources', dataSource)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateDataSource(dataSource) {
  return client.put('/data_sources', dataSource)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteDataSource(id) {
  return client.delete(`/data_sources/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const DataSourceAPI = {
  fetch: fetchDataSources,
  get: getDataSource,
  create: createDataSource,
  update: updateDataSource,
  delete: deleteDataSource,
};

export default DataSourceAPI;
