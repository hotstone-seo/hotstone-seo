import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { RuleForm } from 'components/Rule';
import { createRule } from 'api/rule';
import useDataSources from 'hooks/useDataSources';
import styles from './AddRule.module.css';

function AddRule() {
  const history = useHistory();
  const [dataSources] = useDataSources();

  const handleCreate = (rule) => {
    createRule(rule)
      .then((newRule) => {
        history.push(`/rules/${newRule.id}`);
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
          <Col className={styles.container} span={12} style={{ paddingTop: 24 }}>
            <RuleForm handleSubmit={handleCreate} dataSources={dataSources} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddRule;
