import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { ViewModules, AddModule, EditModule } from './scenes';

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
        render={() => <AddModule />}
      />
      <Route
        path={`${match.url}/:id`}
        render={() => <EditModule />}
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
