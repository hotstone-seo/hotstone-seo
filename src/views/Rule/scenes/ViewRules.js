import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'antd';
import { fetchRules } from 'api/rule';
import { RuleList } from 'components/Rule';

function ViewRules() {
  const [rules, setRules] = useState([]);

  useEffect(() => {
    fetchRules()
      .then(rules => {
        setRules(rules);
      });
  })

  return (
    <div>
      <Button
        type="primary"
        style={{ marginBottom: 16 }}
      >
        <Link to="rules/new">Add new Rule</Link>
      </Button>
      <RuleList rules={rules} />
    </div>
  )
}

export default ViewRules;
