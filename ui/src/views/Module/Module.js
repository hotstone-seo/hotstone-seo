import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { ViewModules } from './scenes';

function Module({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={() => <ViewModules match={match} />}
      />
      <Route
        exact
        path={`${match.url}/new`}
      />
      <Route
        path={`${match.url}/:id`}
      />
    </Switch>
  );
}

Module.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default Module;
