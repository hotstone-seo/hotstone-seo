import client from "./client";

export function fetchRoleTypes(cfg = {}) {
  return client
    .get("/user_roles", cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getRoleType(id) {
  return client
    .get(`/user_roles/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createRoleType(roleType) {
  return client
    .post("/user_roles", roleType)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateRoleType(roleType) {
  return client
    .put("/user_roles", roleType)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteRoleType(id) {
  return client
    .delete(`/user_roles/${id}`)
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
