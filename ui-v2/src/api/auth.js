import client from 'api/client';

// NOTE:
// In this context, authentication simply means the act of a user proving
// his/her credentials through a backend service. If the authentication is successful,
// the application will store the user info until the user is logged out.

export function login(email, password) {
  // TODO: Should be calling API, but the service is yet to be created
  return new Promise((resolve, reject) => {
    resolve({ username: 'johndoe', email: 'john@doe.com', name: 'John Doe' });
  });
}

export function logout() {
  return new Promise((resolve, reject) => {
    resolve('Success!');
  });
}

export function register(user) {
  return new Promise((resolve, reject) => {
    resolve({ username: 'johndoe', email: 'john@doe.com', name: 'John Doe' });
  });
}

export function googleOAuth2GetToken(holder) {
  return client
    .post(
      '/auth/google/token',
      { holder, set_cookie: true },
      { withCredentials: true },
    )
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const AuthAPI = {
  login, logout, register, googleOAuth2GetToken,
};

export default AuthAPI;
