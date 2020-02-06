import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { PageHeader } from 'antd';
import { RuleDetail } from 'components/Rule';
import { AddRule, ViewRules } from './scenes';
import styles from './Rule.module.css';

function Rule() {
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
            path="/rules"
            render={() => <ViewRules />}
          />
          <Route
            exact
            path="/rules/new"
            render={() => <AddRule />}
          />
          <Route
            path="/rules/:id"
            render={({ match }) => <RuleDetail />}
          />
        </Switch>
      </div>
    </div>
  );
}

export default Rule;
