import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { AuthProvider, PrivateRoute } from 'components';
import './App.css';
import DashboardLayout from 'containers/DashboardLayout';
import Login from 'views/Login';

const App = () => (
  <div className="App">
    <AuthProvider>
      <BrowserRouter>
        <Switch>
          <Route
            exact
            path="/login"
            name="Login Page"
            render={(props) => <Login {...props} />}
          />
          <PrivateRoute path="/" name="Home" component={DashboardLayout} />
        </Switch>
      </BrowserRouter>
    </AuthProvider>
  </div>
);

export default App;
