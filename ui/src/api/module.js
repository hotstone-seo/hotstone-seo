import client from './client';

export function fetchModules() {
  return client.get('/modules')
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getModule(id) {
  return client.get(`/modules/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createModule(module) {
  return client.post('/modules', module)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateModule(module) {
  return client.put('/modules', module)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteModule(id) {
  return client.delete(`/modules/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const ModuleAPI = {
  fetch: fetchModules,
  get: getModule,
  create: createModule,
  update: updateModule,
  delete: deleteModule,
};

export default ModuleAPI;
