import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { ViewDataSources } from './scenes';

function DataSource({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={() => <ViewDataSources match={match} />}
      />
    </Switch>
  );
}

DataSource.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default DataSource;
