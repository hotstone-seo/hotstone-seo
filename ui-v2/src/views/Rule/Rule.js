import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { PageHeader } from 'antd';
import { AddRule, EditRule, ViewRules } from './scenes';
import styles from './Rule.module.css';

function Rule({ match }) {
  return (
    <div className="Rule">
      <PageHeader
        className={styles.header}
        title="Rules"
        subTitle="Manage tags on matching URL"
      />
      <div className={styles.content}>
        <Switch>
          <Route
            exact
            path={match.url}
            render={(props) => <ViewRules {...props} />}
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
      </div>
    </div>
  );
}

export default Rule;