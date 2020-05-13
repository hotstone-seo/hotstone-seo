import client from './client';

function fetchUsers(cfg = {}) {
  return client
    .get('/users', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function getUser(id) {
  return client
    .get(`/users/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function createUser(user) {
  return client
    .post('/users', user)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateUser(user) {
  return client
    .put('/users', user)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function deleteUser(id) {
  return client
    .delete(`/users/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function getUserIsExists(user) {
  return client
    .post('/users_is_exists', user)
    .then((response) => response.data)
    .catch((error) => {
      return 0;
    });
}

export {
  fetchUsers, getUser, createUser, updateUser, deleteUser, getUserIsExists,
};

const UserAPI = {
  fetch: fetchUsers,
  get: getUser,
  create: createUser,
  update: updateUser,
  delete: deleteUser,
  getIsExist: getUserIsExists,
};

export default UserAPI;
