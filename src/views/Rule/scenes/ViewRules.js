import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'antd';
import { fetchRules } from 'api/rule';
import { RuleList } from 'components/Rule';

function ViewRules({ match }) {
  const [rules, setRules] = useState([]);

  useEffect(() => {
    let _isCancelled = false;
    fetchRules()
      .then(rules => {
        if (!_isCancelled) {
          setRules(rules);
        }
      });

    return () => {
      _isCancelled = true;
    };
  })

  return (
    <div>
      <Button
        type="primary"
        style={{ marginBottom: 16 }}
      >
        <Link to={`${match.url}/new`}>Add new Rule</Link>
      </Button>
      <RuleList rules={rules} />
    </div>
  )
}

export default ViewRules;
