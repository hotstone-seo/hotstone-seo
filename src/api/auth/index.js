// import client '../client';

// NOTE:
// In this context, authentication simply means the act of a user proving
// his/her credentials through a backend service. If the authentication is successful,
// the application will store the user info until the user is logged out.

function login(username, password) {
  // TODO: Should be calling API, but the service is yet to be created
  return new Promise((resolve, reject) => {
    resolve({ username: 'johndoe', email: 'john@doe.com', name: 'John Doe' });
  });
}

function logout() {
  return new Promise((resolve, reject) => {
    resolve('Success!');
  });
}

function register(user) {
  return new Promise((resolve, reject) => {
    resolve({ username: 'johndoe', email: 'john@doe.com', name: 'John Doe' });
  })
}

export { login, logout, register };

const AuthAPI = { login, logout, register }; 

export default AuthAPI;
