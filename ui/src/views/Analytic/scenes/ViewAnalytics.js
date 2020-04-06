import React, { useState, useEffect } from 'react';
import {
  Button, Row, Col, Card, Form, Select,
} from 'antd';
import useRouter from 'hooks/useRouter';
import { fetchRules } from 'api/rule';
import { HitChartCard, HitCounterCard, UniquePageCounterCard } from 'components/Analytic';

import _ from 'lodash';

const { Option } = Select;

function ViewAnalytics() {
  const { query, history } = useRouter();
  const { ruleID } = query;

  const [listRule, setListRule] = useState([]);

  const onSubmit = ({ ruleID }) => {
    history.push(`/analytic${ruleID === -1 ? '' : `?ruleID=${ruleID}`}`);
  };

  const [form] = Form.useForm();

  useEffect(() => {
    fetchRules().then((rules) => {
      if (rules !== undefined) {
        rules.unshift({ id: -1, name: 'All Rule' });
        setListRule(rules);
      }
    });

    form.setFieldsValue({ ruleID: ruleID ? _.toNumber(ruleID) : -1 });
  }, [form, ruleID]);


  return (
    <div>
      <Row>
        <Col span={24}>
          <Card>
            <Form
              form={form}
              name="horizontal_login"
              layout="inline"
              onFinish={onSubmit}
            >
              <Form.Item
                name="ruleID"
                // label="Rule"
                rules={[{ required: true, message: 'Please select rule' }]}
                style={{ width: '60%' }}
              >
                <Select allowClear>
                  {listRule.map(({ id, name, url_pattern }) => (
                    <Option key={id} value={id}>{`${name}${url_pattern ? `- ${url_pattern}` : ''}`}</Option>
                  ))}
                </Select>
              </Form.Item>
              <Form.Item shouldUpdate>
                {() => (
                  <Button
                    type="primary"
                    htmlType="submit"
                    disabled={
                      // !form.isFieldsTouched(true) ||
                     form
                       .getFieldsError()
                       .filter(({ errors }) => errors.length).length
                    }
                  >
                    Submit
                  </Button>
                )}
              </Form.Item>
            </Form>
          </Card>
        </Col>
      </Row>
      <Row>
        <Col span={4} justify="space-around" align="middle">
          <HitCounterCard ruleID={ruleID} />
        </Col>
        <Col span={4} justify="space-around" align="middle">
          <UniquePageCounterCard ruleID={ruleID} />
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <HitChartCard ruleID={ruleID} />
        </Col>
      </Row>
    </div>
  );
}

export default ViewAnalytics;
