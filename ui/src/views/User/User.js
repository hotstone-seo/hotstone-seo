import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { AddUser, EditUser, ViewUsers } from './scenes';

function User({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={() => <ViewUsers match={match} />}
      />
      <Route
        exact
        path={`${match.url}/new`}
        render={() => <AddUser />}
      />
      <Route
        path={`${match.url}/:id`}
        render={() => <EditUser />}
      />
    </Switch>
  );
}

User.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default User;
