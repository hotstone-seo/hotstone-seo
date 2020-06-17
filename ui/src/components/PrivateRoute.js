import React from "react";
import { Route, Redirect } from "react-router-dom";
import { useAuth } from "./AuthProvider";

function PrivateRoute({ component: Component, ...rest }) {
  const auth = useAuth();

  // console.log(auth.currentUser);

  // TODO: check by paths
  const isAllowAccess = true;

  return (
    <Route
      {...rest}
      render={(props) =>
        auth.currentUser && isAllowAccess ? (
          <Component {...props} />
        ) : auth.currentUser && !isAllowAccess ? (
          <Redirect exact to="/page-403" />
        ) : (
          <Redirect to="/login" />
        )
      }
    />
  );
}

export default PrivateRoute;
