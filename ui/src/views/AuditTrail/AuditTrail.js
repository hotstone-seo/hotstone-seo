import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { ViewAuditTrail } from './scenes';

function AuditTrail({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={() => <ViewAuditTrail match={match} />}
      />
    </Switch>
  );
}

AuditTrail.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default AuditTrail;
