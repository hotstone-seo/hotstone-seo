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

function getUserIsExists(email) {
  return client
    .get(`/users_is_exists/${email}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
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
