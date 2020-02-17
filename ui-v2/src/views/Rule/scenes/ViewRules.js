import React, { useState, useEffect } from "react";
import { Link, useHistory } from "react-router-dom";
import { Button, message } from "antd";
import { fetchRules, deleteRule } from "api/rule";
import { RuleListV2 } from "components/Rule";

function ViewRules({ match }) {
  const [rules, setRules] = useState([]);

  useEffect(() => {
    let _isCancelled = false;
    fetchRules().then(rules => {
      if (!_isCancelled) {
        setRules(rules);
      }
    });

    return () => {
      _isCancelled = true;
    };
  }, []);

  let history = useHistory();

  const showEditForm = rule => {
    history.push(`${match.url}/${rule.id}`);
  };

  const handleDelete = rule => {
    deleteRule(rule.id)
      .then(() => {
        message.success(`Successfully deleted ${rule.name}`);
        setRules(rules.filter(item => item.id !== rule.id));
      })
      .catch(error => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <Button type="primary" style={{ marginBottom: 16 }}>
        <Link to={`${match.url}/new`}>Add new Rule</Link>
      </Button>
      <RuleListV2
        rules={rules}
        onClick={showEditForm}
        onEdit={showEditForm}
        onDelete={handleDelete}
      />
    </div>
  );
}

export default ViewRules;
