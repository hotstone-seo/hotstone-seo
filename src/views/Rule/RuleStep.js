import React, { useState } from 'react';
import { Steps, Button, Row, Col } from 'antd';
import RuleForm from './RuleForm';
import styles from './RuleStep.module.css';

const { Step } = Steps;

const steps = [
  {
    key: 'rule',
    title: 'Create a Rule',
    render: (props) => (
      <RuleForm {...props} />
    ),
  },
  {
    key: 'tags',
    title: 'Add Tags',
    render: (props) => (
      <div>This will be the Tag form</div>
    )
  }
];

function RuleStep() {
  const [current, setCurrent] = useState(0);

  const next = () => {
    setCurrent(current + 1);
  }

  const prev = () => {
    setCurrent(current - 1);
  }

  return (
    <div className={styles.container}>
      <Row gutter={[0, 32]}>
        <Col span={12} offset={6}>
          <Steps current={current} size="small" className={styles.steps}>
            {steps.map(({ key, title, description }) => (
              <Step key={key} title={title} description={description} />
            ))}
          </Steps>
        </Col>
      </Row>
      <Row>
        <Col span={16} offset={4}>
          {steps[current].render()}
        </Col>
      </Row>
      <div>
        {current < steps.length - 1 && (
          <Button type="primary" onClick={next}>
            Next
          </Button>
        )}
        {current === steps.length - 1 && (
          <Button type="primary" >
            Done
          </Button>
        )}
        {current > 0 && (
          <Button onClick={prev} className={styles.previousButton}>
            Previous
          </Button>
        )}
      </div>
    </div>
  );
}

export default RuleStep;
