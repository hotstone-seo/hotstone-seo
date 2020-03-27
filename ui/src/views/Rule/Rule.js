import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';
import { AddRule, EditRule, ViewRules } from './scenes';

function Rule({ match }) {
  return (
    <Switch>
      <Route
        exact
        path={match.url}
        render={() => <ViewRules match={match} />}
      />
      <Route
        exact
        path={`${match.url}/new`}
        render={() => <AddRule />}
      />
      <Route
        path={`${match.url}/:id`}
        render={() => <EditRule />}
      />
    </Switch>
  );
}

Rule.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default Rule;
