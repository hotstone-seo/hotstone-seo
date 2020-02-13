import React from 'react';
import { useHistory } from 'react-router-dom';
import { Row, Col, message } from 'antd';
import { RuleForm } from 'components/Rule';
import { createRule } from 'api/rule';
import styles from './AddRule.module.css';

function AddRule() {
  const history = useHistory();
  const handleCreate = (rule) => {
    createRule(rule)
      .then(rule => {
        history.push('/rules');
      })
      .catch(error => {
        message.error(error.message);
      });
  };
  return (
    <div>
      <Row>
        <Col className={styles.container} span={12} style={{ paddingTop: 24 }}>
          <RuleForm handleSubmit={handleCreate} />
        </Col>
      </Row>
    </div>
  );
}

export default AddRule;
