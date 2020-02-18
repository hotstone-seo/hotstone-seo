import React from "react";
import { Switch, Route } from "react-router-dom";
import { PageHeader } from "antd";
import { ViewMismatchRules } from "./scenes";
import styles from "./MismatchRule.module.css";

function MismatchRule({ match }) {
  return (
    <div className="MismatchRule">
      <PageHeader
        className={styles.header}
        title="Mismatch Rules"
        subTitle="URLs not matched with any rules"
      />
      <div className={styles.content}>
        <Switch>
          <Route
            exact
            path={match.url}
            render={props => <ViewMismatchRules {...props} />}
          />
        </Switch>
      </div>
    </div>
  );
}

export default MismatchRule;
