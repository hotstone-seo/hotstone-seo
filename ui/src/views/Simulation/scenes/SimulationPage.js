import React from "react";
import {
  Form,
  Input,
  Button,
  Row,
  Col,
  Card,
  Alert,
  Select,
  Descriptions,
} from "antd";
import { Machine, assign } from "xstate";
import { useMachine } from "@xstate/react";
import { Link } from "react-router-dom";
import _ from "lodash";
import parse from "url-parse";

import { match, fetchTags } from "api/provider";
import { getRule } from "api/rule";
import { RawHtmlPreview } from "components/Simulation";

import locales from "locales";

const { Option } = Select;

const pageMachine = Machine({
  id: "simulation",
  initial: "idle",
  context: {
    url: null,
    locale: locales[0],
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
          target: "success",
          actions: assign({
            matchResp: (context, event) => {
              console.log("RESP: ", event);
              return event.data;
            },
          }),
        },
        onError: {
          target: "failed",
          actions: assign({
            matchError: (context, event) => {
              console.log("ERR :", event);
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
      target: ".submitting",
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
    throw new Error("Not matched");
  }
  // TODO: double check if rule return ID#0
  const tags = await fetchTags(rule, locale);
  const ruleDetail = await getRule(rule.rule_id);
  const data = { rule, tags, ruleDetail };
  return data;
}

function SimulationPage() {
  const [current, send] = useMachine(pageMachine);
  const { matchResp, matchError, pageError } = current.context;

  const [form] = Form.useForm();
  // form.setFieldsValue({ locale: current.context.locale });
  const onSubmit = ({ locale, url }) => {
    const urlObj = parse(url);

    send("SUBMIT", { locale, url: urlObj.pathname });
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
              initialValues={{
                locale: current.context.locale,
              }}
            >
              <Form.Item
                name="url"
                rules={[{ required: true, message: "Please input URL" }]}
                style={{ width: "60%" }}
              >
                <Input placeholder="URL" />
              </Form.Item>
              <Form.Item
                name="locale"
                rules={[{ required: true, message: "Please select locale" }]}
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
                      isLoading(current) ||
                      current.matches("pageFailed") ||
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
    const { rule, tags, ruleDetail } = matchResp;
    const { rule_id, path_param } = rule;
    return (
      <Card>
        <Alert type="success" message="Matched" />
        <br />
        {!_.isEmpty(ruleDetail) && (
          <Descriptions key={0} title="Rule" column={1} bordered>
            <Descriptions.Item key={1} label="Name">
              {ruleDetail.name}
            </Descriptions.Item>
            <Descriptions.Item key={2} label="URL Pattern">
              {ruleDetail.url_pattern}
            </Descriptions.Item>
            <Descriptions.Item key={3} label="Detail">
              <Link to={`/rules/${rule_id}`}>Rule Detail</Link>
            </Descriptions.Item>

            {!_.isEmpty(path_param) && (
              <Descriptions.Item key={4} label="Path Params">
                <table className="ant-table">
                  <thead>
                    <tr key="id">
                      <th>Path Param</th>
                      <th>Value</th>
                    </tr>
                  </thead>
                  <tbody className="ant-table-tbody">
                    {Object.entries(path_param).map(([key, value]) => (
                      <tr key={key}>
                        <td>{key}</td>
                        <td>{value}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </Descriptions.Item>
            )}
          </Descriptions>
        )}

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
  return current.matches("init") || current.matches("submitting");
}

export default SimulationPage;
