import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { deleteRule, updateStatusStartRule } from 'api/rule';
import { RuleListV2 } from 'components/Rule';
import { PlusOutlined } from '@ant-design/icons';

function ViewRules({ match }) {
  const history = useHistory();
  const [listRule, setListRule] = useState([]);

  const showEditForm = (rule) => {
    history.push(`${match.url}/${rule.id}`);
  };

  const handleDelete = (rule) => {
    deleteRule(rule.id)
      .then(() => {
        message.success(`Successfully deleted ${rule.name}`);
        setListRule(listRule.filter((item) => item.id !== rule.id));
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const handleUpdateStatusStart = (checked, rule) => {
    const onOFF = checked === true ? 'START' : 'STOP';
    /* updateStatusStartRule(rule.id)
      .then(() => {

        setListRule(listRule.filter((item) => item.id !== rule.id));
      })
      .catch((error) => {
        message.error(error.message);
      }); */
    message.success(`Successfully switch ${rule.name} to be ${onOFF}`);
    // console.log(`switch to ${checked}`);
  };

  const addDataRule = () => {
    history.push(`${match.url}/new`);
  };

  return (
    <div>
      <PageHeader
        title="Rules"
        subTitle="Manage Tags on matching URL Rules"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Button
          data-testid="btn-new-rule"
          type="primary"
          style={{ marginBottom: 16 }}
          icon={<PlusOutlined />}
          onClick={() => addDataRule()}
        >
          Add New Rule
        </Button>
        <RuleListV2
          onClick={showEditForm}
          onEdit={showEditForm}
          onDelete={handleDelete}
          listRule={listRule}
          setListRule={setListRule}
          onChangeToggleButton={handleUpdateStatusStart}
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
