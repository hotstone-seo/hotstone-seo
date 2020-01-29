import React, { useState, useEffect } from 'react';
import { Switch, Route } from 'react-router-dom';
import { PageHeader } from 'antd';
import RuleList from './RuleList';
import { fetchRules } from '../../api/rules';

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
    <React.Fragment>
      <PageHeader
        title="Rules"
        subtitle="Manage tags on matching URL"
      />
      <Switch>
        <Route
          exact
          path="/rules"
          render={() => <RuleList rules={rules} />}
        />
      </Switch>
    </React.Fragment>
  );
}

export default Rule;
