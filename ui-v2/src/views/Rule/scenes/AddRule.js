import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { RuleForm } from 'components/Rule';
import { createRule } from 'api/rule';
import useDataSources from 'hooks/useDataSources';

function AddRule() {
  const history = useHistory();
  const [dataSources] = useDataSources();

  const handleCreate = (rule) => {
    createRule(rule)
      .then((newRule) => {
        history.push(`/rules/${newRule.id}`, { message: `${newRule.name} is successfully created` });
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/rules')}
        title="Add new Rule"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <RuleForm onSubmit={handleCreate} dataSources={dataSources} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddRule;
