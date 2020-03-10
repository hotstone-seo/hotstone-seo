import React from 'react';
import Cookies from 'js-cookie';
import jwt from 'jsonwebtoken';
import _ from 'lodash';
import AuthAPI from '../api/auth';

const AuthContext = React.createContext();

const loadUserFromCookie = () => {
  const token = Cookies.get('token');

  if (_.isEmpty(token)) {
    return null;
  }

  const tokenDecoded = jwt.decode(token);
  const user = {
    email: tokenDecoded.email,
    picture: tokenDecoded.picture,
  };

  return user;
};

function AuthProvider(props) {
  const [currentUser, setCurrentUser] = React.useState(loadUserFromCookie());

  const login = (email, password) => {
    if (!currentUser) {
      return AuthAPI.login(email, password).then((user) => {
        setCurrentUser(user);
      });
    }
    throw new Error('Error login: another user already logged in');
  };
  const logout = () => {
    if (currentUser) {
      return AuthAPI.logout().then(() => {
        setCurrentUser(null);
      });
    }
    throw new Error("Error logout: there's no user currently logged in");
  };

  const register = () => AuthAPI.register();

  return (
    <AuthContext.Provider
      value={{
        currentUser, login, logout, register,
      }}
      {...props}
    />
  );
}

function useAuth() {
  const context = React.useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

export { AuthProvider, useAuth };
