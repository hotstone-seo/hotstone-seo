import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { Link, useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { fetchRules, deleteRule } from 'api/rule';
import { RuleListV2 } from 'components/Rule';

function ViewRules({ match }) {
  const [rules, setRules] = useState([]);

  useEffect(() => {
    let isCancelled = false;
    fetchRules().then((newRules) => {
      if (!isCancelled) {
        setRules(newRules);
      }
    });

    return () => {
      isCancelled = true;
    };
  }, []);

  const history = useHistory();

  const showEditForm = (rule) => {
    history.push(`${match.url}/${rule.id}`);
  };

  const handleDelete = (rule) => {
    deleteRule(rule.id)
      .then(() => {
        message.success(`Successfully deleted ${rule.name}`);
        setRules(rules.filter((item) => item.id !== rule.id));
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        title="Rules"
        subTitle="Manage Tags on matching URL Rules"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
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
    </div>
  );
}

ViewRules.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ViewRules;
