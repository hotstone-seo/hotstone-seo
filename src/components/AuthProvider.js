import React, { useState } from 'react';
import AuthAPI, { register } from '../api/auth';

const AuthContext = React.createContext();

function AuthProvider({ children }) {
  // TODO: Try to get data from localStorage to initialize the state
  const [user, setUser] = useState(null);
  const login = (email, password) => {
    if(!user) {
      AuthAPI.login(email, password)
             .then((user) => {
               setUser(user);
             });
    } else {
      throw new Error('Error login: another user already logged in');
    } 
  }
  const logout = () => {
    if(user) {
      AuthAPI.logout()
             .then(() => {
               setUser(null);
             })
    } else {
      throw new Error("Error logout: there's no user currently logged in");
    }
  }
  return (
    <AuthContext.Provider value={{ user, login, logout, register }} >
      {children}
    </AuthContext.Provider>
  );
}

const useAuth = () => React.useContext(AuthContext);

export { AuthProvider, useAuth };
