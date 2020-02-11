import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import { Row, Col, message } from 'antd';
import { RuleForm } from 'components/Rule';
import { TagList } from 'components/Tag';
import { getRule, updateRule } from 'api/rule';
import { fetchTags } from 'api/tag';
import styles from './AddRule.module.css';

function EditRule() {
  const { id } = useParams();
  const history = useHistory();

  const [rule, setRule] = useState({});

  const [tags, setTags] = useState([]);

  useEffect(() => {
    getRule(id)
      .then(rule => { setRule(rule) })
      .catch(error => {
        message.error(error.message);
      });
  }, [id]);

  useEffect(() => {
    fetchTags({ rule_id: id })
      .then(tags => { setTags(tags) })
      .catch(error => {
        message.error(error.message);
      });
  }, [id])

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
        <Col className={styles.container} span={12} style={{ paddingTop: 24 }}>
          <RuleForm handleSubmit={handleEdit} rule={rule} />
        </Col>
      </Row>
      <Row style={{ marginTop: 24 }}>
        <Col className={styles.container} span={16} style={{ padding: 24 }}>
          <TagList tags={tags} />
        </Col>
      </Row>
    </div>
  );
}

export default EditRule;
