import React from 'react';
import { matchPath } from 'react-router';
import { Route, Redirect, useLocation } from 'react-router-dom';
import { useAuth } from './AuthProvider';

function PrivateRoute({ component: Component, ...rest }) {
  const auth = useAuth();
  const location = useLocation();
  // const firstPath = '/'.concat(location.pathname.split('/')[1]);

  let listModules = null;
  let jsonModules = auth.currentUser === null ? null : auth.currentUser.modules;

  let isMatch = false;
  let loopStop = false;

  if (jsonModules) {
    jsonModules = JSON.parse(jsonModules);

    Object.keys(jsonModules).forEach((key) => {
      listModules = jsonModules[key];
    });
    if (location.pathname === '/') isMatch = true; // special case for Home URL (/)
    else {
      listModules.forEach((item, index) => {
        if (loopStop) { return; }

        // check regex pattern
        isMatch = matchPath(location.pathname, {
          path: '/'.concat(item.pattern),
          exact: true,
          strict: false,
        });
        if (isMatch) loopStop = true;
      });
    }
    // if have not access Obj isMatch = null
  }
  const isAllowAccess = isMatch !== null;

  return (
    <Route
      {...rest}
      render={(props) => (
        auth.currentUser && isAllowAccess ? <Component {...props} />
          : (auth.currentUser && (!isAllowAccess) ? <Redirect exact to="/page-403" /> : <Redirect to="/login" />)
      )}
    />
  );
}

export default PrivateRoute;
