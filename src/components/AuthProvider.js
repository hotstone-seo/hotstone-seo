import React from 'react';
import AuthAPI, { register } from '../api/auth';

const AuthContext = React.createContext();

function AuthProvider(props) {
  // TODO: Try to get data from localStorage to initialize the state
  const [currentUser, setCurrentUser] = React.useState(
    JSON.parse(localStorage.getItem('currentUser')) || null
  );

  React.useEffect(() => {
    localStorage.setItem('currentUser', JSON.stringify(currentUser));
  }, [currentUser])
  
  const login = (email, password) => {
    if(!currentUser) {
      return AuthAPI.login(email, password)
                    .then((user) => {
                      setCurrentUser(user);
                    });
    } else {
      throw new Error('Error login: another user already logged in');
    } 
  }
  const logout = () => {
    if(currentUser) {
      return AuthAPI.logout()
                    .then(() => {
                      setCurrentUser(null);
                    })
    } else {
      throw new Error("Error logout: there's no user currently logged in");
    }
  }
  return (
    <AuthContext.Provider value={{ currentUser, login, logout, register }} {...props} />
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
