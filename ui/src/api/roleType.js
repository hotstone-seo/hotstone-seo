import client from './client';

export function fetchRoleTypes() {
  return client.get('/role_types')
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getRoleType(id) {
  return client.get(`/role_types/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createRoleType(roleType) {
  return client.post('/role_types', roleType)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateRoleType(roleType) {
  return client.put('/role_types', roleType)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteRoleType(id) {
  return client.delete(`/role_types/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const RoleTypeAPI = {
  fetch: fetchRoleTypes,
  get: getRoleType,
  create: createRoleType,
  update: updateRoleType,
  delete: deleteRoleType,
};

export default RoleTypeAPI;
