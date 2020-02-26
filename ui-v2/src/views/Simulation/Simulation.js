import React from "react";
import { Switch, Route } from "react-router-dom";
import { PageHeader } from "antd";
import { SimulationPage } from "./scenes";
import styles from "./Simulation.module.css";

function Simulation({ match }) {
  return (
    <div className="Simulation">
      <PageHeader
        className={styles.header}
        title="Simulation"
        subTitle="Rule Matching"
      />
      <div className={styles.content}>
        <Switch>
          <Route
            exact
            path={match.url}
            render={props => <SimulationPage {...props} />}
          />
        </Switch>
      </div>
    </div>
  );
}

export default Simulation;
