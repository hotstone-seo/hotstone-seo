import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { AddRoleType, EditRoleType, ViewRoleTypes } from './scenes';

function RoleType({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={() => <ViewRoleTypes match={match} />}
      />
      <Route
        exact
        path={`${match.url}/new`}
        render={() => <AddRoleType />}
      />
      <Route
        path={`${match.url}/:id`}
        render={() => <EditRoleType />}
      />
    </Switch>
  );
}

RoleType.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default RoleType;
