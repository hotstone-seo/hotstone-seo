import axios from 'axios';
import client from './client';

export function fetchDataSources() {
  return client.get('/data_sources')
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function fetchDataSourcesPagination(cfg = {}) {
  return client.get('/data_sources', cfg)
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

export function getDataSourcesByIDs(ids) {
  const reqURLs = ids.map((id) => `/data_sources/${id}`);
  return axios.all(reqURLs.map((reqURL) => client.get(reqURL)))
    .then((responses) => responses.map((response) => response.data))
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
  fetchPagination: fetchDataSourcesPagination,
};

export default DataSourceAPI;
