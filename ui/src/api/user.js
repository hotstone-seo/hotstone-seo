import client from './client';

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

export {
  getUser, createUser, updateUser, deleteUser,
};

const UserAPI = {
  get: getUser,
  create: createUser,
  update: updateUser,
  delete: deleteUser,
};

export default UserAPI;
