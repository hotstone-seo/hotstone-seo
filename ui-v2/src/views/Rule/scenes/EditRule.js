import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import { Row, Col, message } from 'antd';
import { RuleForm } from 'components/Rule';
import { getRule, updateRule } from 'api/rule';
import styles from './AddRule.module.css';

function EditRule() {
  const { id } = useParams();
  const history = useHistory();

  const [rule, setRule] = useState({});

  useEffect(() => {
    getRule(id)
      .then(rule => { setRule(rule) })
      .catch(error => {
        message.error(error.message);
      });
  }, [id]);

  const handleEdit = (newRule) => {
    newRule.id = rule.id;
    updateRule(newRule)
      .then(() => {
        history.push('/rules');
      })
      .catch(error => {
        message.error(error.message);
      });
  };
  return (
    <div>
      <Row>
        <Col className={styles.container} span={12}>
          <RuleForm handleSubmit={handleEdit} rule={rule} />
        </Col>
      </Row>
    </div>
  );
}

export default EditRule;
