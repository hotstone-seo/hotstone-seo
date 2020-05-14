import React from 'react';
import { Route, Redirect, useLocation } from 'react-router-dom';
import { useAuth } from './AuthProvider';
import { isAllowed } from './AuthRole';

// TODO : get value from cookie
const roleUserLogin = {
  rights: ['/datasources', '/mismatch-rule', '/rules', '/analytic', '/simulation', '/audit-trail', '/users'],
};

function PrivateRoute({ component: Component, ...rest }) {
  const auth = useAuth();
  const location = useLocation();
  const firstPath = '/'.concat(location.pathname.split('/')[1]);

  let mn;
  /*
  // TODO : get value from cookie
  Object.keys(auth.currentUser.modules).forEach((key) => {
    mn = '/'.concat(auth.currentUser.modules[key]);
  });
  const roleUserLogin = {
    rights: mn,
  };
  */
  const isAllowAccess = isAllowed(roleUserLogin, [firstPath]);

  // TODO : before return, check regex using function matchPath

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

