import React from 'react';
import {
  Form, Input, Button, Row, Col, Card, Alert, Select,
} from 'antd';
import { Machine, assign } from 'xstate';
import { useMachine } from '@xstate/react';
import { Link } from 'react-router-dom';
import _ from 'lodash';
import parse from 'url-parse';

import { match, fetchTags } from 'api/provider';
import { RawHtmlPreview } from 'components/Simulation';

import locales from 'locales';

const { Option } = Select;

const pageMachine = Machine({
  id: 'simulation',
  initial: 'idle',
  context: {
    url: null,
    locale: process.env.REACT_APP_LOCALE,
    listLocale: locales,
    matchResp: null,
    matchError: null,
    pageError: null,
  },
  states: {
    idle: {},
    init: {},
    submitting: {
      entry: assign({
        matchResp: null,
        matchError: null,
      }),
      invoke: {
        src: (context) => matchThenGetTags(context.locale, context.url),
        onDone: {
          target: 'success',
          actions: assign({
            matchResp: (context, event) => {
              console.log('RESP: ', event);
              return event.data;
            },
          }),
        },
        onError: {
          target: 'failed',
          actions: assign({
            matchError: (context, event) => {
              console.log('ERR :', event);
              return event.data;
            },
          }),
        },
      },
    },
    success: {},
    failed: {},
    pageFailed: {},
  },
  on: {
    SUBMIT: {
      target: '.submitting',
      actions: assign({
        url: (context, event) => event.url,
        locale: (context, event) => event.locale,
      }),
    },
  },
});

async function matchThenGetTags(locale, url) {
  const rule = await match(url);
  if (_.isEmpty(rule)) {
    throw new Error('Not matched');
  }
  const tags = await fetchTags(rule, locale);
  const data = { rule, tags };
  return data;
}

function SimulationPage() {
  const [current, send] = useMachine(pageMachine);
  const { matchResp, matchError, pageError } = current.context;

  const [form] = Form.useForm();
  form.setFieldsValue({ locale: current.context.locale });
  const onSubmit = ({ locale, url }) => {
    const urlObj = parse(url);

    send('SUBMIT', { locale, url: urlObj.pathname });
  };

  return (
    <>
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
                name="url"
                rules={[{ required: true, message: 'Please input URL' }]}
                style={{ width: '60%' }}
              >
                <Input placeholder="URL" />
              </Form.Item>
              <Form.Item
                name="locale"
                rules={[{ required: true, message: 'Please select locale' }]}
              >
                <Select>
                  {current.context.listLocale.map((locale, index) => (
                    <Option key={index} value={locale}>
                      {locale}
                    </Option>
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
                      isLoading(current)
                      || current.matches('pageFailed')
                      || form
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
        <Col span={24}>
          {isLoading(current) && <Card>Loading ...</Card>}
          {renderIfSuccess(matchResp)}
          {renderIfError(matchError)}
          {renderIfPageError(pageError)}
        </Col>
      </Row>
    </>
  );
}

function renderIfSuccess(matchResp) {
  if (matchResp) {
    const { rule_id, path_param } = matchResp.rule;
    const { tags } = matchResp;
    return (
      <Card>
        <Alert
          type="success"
          message={(
            <>
              Matched (
              <Link to={`/rule-detail/?id=${rule_id}`}>Rule Detail</Link>
              )
            </>
          )}
          description={(
            <>
              {!_.isEmpty(path_param) && (
                <div>
                  <div>Path params:</div>
                  <ul>
                    {Object.entries(path_param).map(([key, value]) => (
                      <li key={key}>
                        {key}
                        :
                        {value}
                      </li>
                    ))}
                  </ul>
                </div>
              )}
            </>
          )}
        />
        <br />
        <strong>Raw HTML Tags Preview</strong>
        <br />
        <RawHtmlPreview ruleID={rule_id} tags={tags} />
      </Card>
    );
  }
}

function renderIfError(matchError) {
  if (matchError) {
    let msgError = matchError.message;
    if (matchError.response) {
      msgError = matchError.response.data.message;
    }
    return (
      <Card>
        <Alert type="error" message={msgError} />
      </Card>
    );
  }
}

function renderIfPageError(pageError) {
  if (pageError) {
    let msgError = pageError.message;
    if (pageError.response) {
      msgError = pageError.response.data.message;
    }
    return (
      <Card>
        <Alert type="error" message={msgError} />
      </Card>
    );
  }
}

function isLoading(current) {
  return current.matches('init') || current.matches('submitting');
}

export default SimulationPage;
