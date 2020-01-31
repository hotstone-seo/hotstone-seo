import React, { useState, useEffect } from 'react';
import { Switch, Route } from 'react-router-dom';
import { PageHeader, Row } from 'antd';
import RuleList from './RuleList';
import RuleForm from './RuleForm';
import { fetchRules } from '../../api/rule';
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
        </Switch>
      </div>
    </div>
  );
}

export default Rule;
