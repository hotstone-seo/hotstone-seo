import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Spin } from 'antd';
import './App.css';

const loading = () => (
  <Spin className="loading-spin" size="large" />
);

const DefaultLayout = React.lazy(() => import('./containers/DefaultLayout'));

const Login = React.lazy(() => import('./views/Pages/Login'));

const App = () => (
  <div className="App">
    <BrowserRouter>
      <React.Suspense fallback={loading()}>
        <Switch>
          <Route
            exact
            path="/login"
            name="Login Page"
            render={props => <Login {...props} />}
          />
          <Route
            path="/"
            name="Home"
            render={props => <DefaultLayout {...props} />}
          />
        </Switch>
      </React.Suspense>
    </BrowserRouter>
  </div>
 )

export default App;
