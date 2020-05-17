import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { AddClientKey, EditClientKey, ViewClientKeys } from './scenes';

function ClientKey({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={() => <ViewClientKeys match={match} />}
      />
      <Route
        exact
        path={`${match.url}/new`}
        render={() => <AddClientKey />}
      />
      <Route
        path={`${match.url}/:id`}
        render={() => <EditClientKey />}
      />
    </Switch>
  );
}

ClientKey.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ClientKey;
