import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { PageHeader } from 'antd';
import { ViewAnalytics } from './scenes';
import styles from './Analytic.module.css';

function Analytic({ match }) {
  return (
    <div className="Analytic">
      <PageHeader
        className={styles.header}
        title="Analytic"
        subTitle="Analytic"
      />
      <div className={styles.content}>
        <Switch>
          <Route
            exact
            path={match.url}
            render={(props) => <ViewAnalytics {...props} />}
          />
        </Switch>
      </div>
    </div>
  );
}

export default Analytic;
