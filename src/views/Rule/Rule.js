import React, { useState, useEffect } from 'react';
import { Switch, Route } from 'react-router-dom';
import { PageHeader } from 'antd';
import { fetchRules } from 'api/rule';
import { RuleDetail, RuleForm, RuleList } from 'components/Rule';
import styles from './Rule.module.css';

function Rule() {
  const [rules, setRules] = useState([]);

  useEffect(() => {
    let _isCancelled = false;
    fetchRules()
      .then((rules) => {
        if (!_isCancelled) {
          setRules(rules);
        }
      });

    return () => {
      _isCancelled = true;
    };
  });

  const findRule = (id) => {
    return rules.find(rule => rule.id.toString() === id) || {};
  };

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
            render={() => <RuleList rules={rules} />}
          />
          <Route
            exact
            path="/rules/new"
            render={() => <RuleForm />}
          />
          <Route
            path="/rules/:id"
            render={({ match }) => <RuleDetail rule={findRule(match.params.id)}/>}
          />
        </Switch>
      </div>
    </div>
  );
}

export default Rule;
