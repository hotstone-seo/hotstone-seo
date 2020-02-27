import client from './client';

function fetchDataSources() {
  return client.get('/data_sources')
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function getDataSource(id) {
  return client.get(`/data_sources/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function createDataSource(dataSource) {
  return client.post('/data_sources', dataSource)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateDataSource(dataSource) {
  return client.put('/data_sources', dataSource)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function deleteDataSource(id) {
  return client.delete(`/data_sources/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export {
  fetchDataSources, getDataSource, createDataSource, updateDataSource, deleteDataSource,
};

const DataSourceAPI = {
  fetch: fetchDataSources,
  get: getDataSource,
  create: createDataSource,
  update: updateDataSource,
  delete: deleteDataSource,
};

export default DataSourceAPI;
