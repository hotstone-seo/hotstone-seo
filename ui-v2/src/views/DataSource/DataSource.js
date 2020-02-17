import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { ViewDataSources } from './scenes';

function DataSource({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={(props) => <ViewDataSources {...props} />}
      />
    </Switch>
  );
}

export default DataSource;
