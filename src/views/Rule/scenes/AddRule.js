import React from 'react';
import { Row, Col } from 'antd';
import { RuleForm } from 'components/Rule';
import styles from './AddRule.module.css';

function AddRule() {
  return (
    <div>
      <Row>
        <Col className={styles.container} span={12}>
          <RuleForm />
        </Col>
      </Row>
    </div>
  );
}

export default AddRule;
