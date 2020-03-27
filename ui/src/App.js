import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Spin } from 'antd';
import { AuthProvider, PrivateRoute } from 'components';
import './App.css';

const loading = () => <Spin className="loading-spin" size="large" />;

const DashboardLayout = React.lazy(() => import('containers/DashboardLayout'));

const Login = React.lazy(() => import('views/Login'));

const App = () => (
  <div className="App">
    <AuthProvider>
      <BrowserRouter>
        <React.Suspense fallback={loading()}>
          <Switch>
            <Route
              exact
              path="/login"
              name="Login Page"
              render={(props) => <Login {...props} />}
            />
            <PrivateRoute path="/" name="Home" component={DashboardLayout} />
          </Switch>
        </React.Suspense>
      </BrowserRouter>
    </AuthProvider>
  </div>
);

export default App;
