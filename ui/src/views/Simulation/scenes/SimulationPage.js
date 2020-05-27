import React from "react";
import {
  Form,
  Input,
  Button,
  Card,
  Alert,
  Select,
  Descriptions,
  Collapse,
  Col,
  Result,
} from "antd";
import { Machine, assign } from "xstate";
import { useMachine } from "@xstate/react";
import { Link } from "react-router-dom";
import _ from "lodash";
import parse from "url-parse";
import { match, fetchTags } from "api/provider";
import { getRule } from "api/rule";
import locales from "locales";
import SyntaxHighlighter from "react-syntax-highlighter";
import { docco } from "react-syntax-highlighter/dist/cjs/styles/hljs";

const { Option } = Select;
const { Panel } = Collapse;

let execTime = "";

function millisToMinutesAndSeconds(duration) {
  const milliseconds = parseInt(duration % 1000, 10);
  return "".concat(milliseconds, " milliseconds");
}

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
        src: (context) => simulateMatch(context.locale, context.url),
        onDone: {
          target: "success",
          actions: assign({
            matchResp: (context, event) => {
              return event.data;
            },
          }),
        },
        onError: {
          target: "failed",
          actions: assign({
            matchError: (context, event) => {
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

async function simulateMatch(locale, url) {
  const startTime = new Date().getTime();
  const rule = await match(url);
  if (_.isEmpty(rule) || rule.rule_id === 0) {
    execTime = millisToMinutesAndSeconds(new Date().getTime() - startTime);
    throw new Error("Not matched Rule.");
  }
  const tags = await fetchTags(rule, locale);
  const ruleDetail = await getRule(rule.rule_id);
  const data = { rule, tags, ruleDetail };
  execTime = millisToMinutesAndSeconds(new Date().getTime() - startTime);
  return data;
}

function SimulationPage() {
  const [current, send] = useMachine(pageMachine);
  const { matchResp, matchError, pageError } = current.context;
  const onSubmit = ({ locale, url }) => {
    const urlObj = parse(url);
    send("SUBMIT", { locale, url: urlObj.pathname });
  };

  return (
    <>
      {renderForm(current, onSubmit)}
      {isLoading(current) && <Card>Loading ...</Card>}
      {matchResp && renderResp(matchResp)}
      {matchError && renderMatchError(matchError)}
      {pageError && renderPageError(pageError)}
    </>
  );
}

function renderForm(current, onSubmit) {
  const [form] = Form.useForm();

  return (
    <Col span={24}>
      <Form
        form={form}
        onFinish={onSubmit}
        layout="inline"
        initialValues={{
          locale: current.context.locale,
        }}
      >
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

        <Form.Item
          name="url"
          rules={[{ required: true, message: "Please input URL path" }]}
          style={{ width: "60%" }}
        >
          <Input placeholder="URL Path" />
        </Form.Item>

        <Form.Item shouldUpdate>
          {() => (
            <Button
              type="primary"
              htmlType="submit"
              disabled={
                isLoading(current) ||
                current.matches("pageFailed") ||
                form.getFieldsError().filter(({ errors }) => errors.length)
                  .length
              }
            >
              Submit
            </Button>
          )}
        </Form.Item>
      </Form>
    </Col>
  );
}

function renderResp(matchResp) {
  const { rule, tags, ruleDetail } = matchResp;
  const { rule_id, path_param } = rule;

  return (
    <>
      <br />
      <Card>
        <Alert
          type="success"
          message={`Matched. Execution time : ${execTime}`}
          closable
        />
        <br />
        <Collapse defaultActiveKey={["3"]} expandIconPosition="left">
          {!_.isEmpty(ruleDetail) && (
            <Panel
              header="Rule"
              key="1"
              extra={<Link to={`/rules/${rule_id}`}>Go To Page</Link>}
            >
              {renderRuleDetail(ruleDetail, rule_id)}
            </Panel>
          )}
          {!_.isEmpty(path_param) && (
            <Panel header="Parameter" key="2">
              {renderParameterTable(path_param)}
            </Panel>
          )}
          <Panel header="Preview" key="3">
            {renderPreview(rule_id, tags)}
          </Panel>
        </Collapse>
      </Card>
    </>
  );
}

function renderMatchError(matchError) {
  let msgError = matchError.message;
  let ps = 0;
  if (msgError !== "") ps = msgError.search("500");

  if (ps > 0) {
    msgError = matchError.response.data.message;
  }
  msgError = msgError.concat(" Execution Time :", execTime);
  return (
    <>
      <br />
      <Card>
        <Alert type="error" message={msgError} closable />
        <Result
          status="404"
          title="Rule Not Match"
          subTitle="Sorry, URL path not matching with any rule."
        />
      </Card>
    </>
  );
}

function renderPageError(pageError) {
  let msgError = pageError.message;
  if (pageError.response) {
    msgError = pageError.response.data.message;
  }
  const subTitle = "".concat("Sorry, the server is wrong. ", msgError);
  return (
    <>
      <br />
      <Result
        status="500"
        title="Error"
        subTitle={subTitle}
        extra={<Button type="primary">Back Home</Button>}
      />
    </>
  );
}

function isLoading(current) {
  return current.matches("init") || current.matches("submitting");
}

function renderPreview(ruleID, tags) {
  if (_.isEmpty(tags)) {
    return (
      <div>
        No tags data. Register tags at{" "}
        <Link to={`/rules/${ruleID}`}>Rule Detail</Link>
      </div>
    );
  }
  const textAreaVal = tags
    .map(({ type, value, attributes }) => {
      let attributesStr = "";
      if (!_.isEmpty(attributes)) {
        if (_.isPlainObject(attributes)) {
          Object.entries(attributes).forEach(([key, value]) => {
            attributesStr += ` ${key}="${value}"`;
          });
        } else if (_.isArray(attributes)) {
          attributes.forEach((attributes) => {
            Object.entries(attributes).forEach(([key, value]) => {
              attributesStr += ` ${key}="${value}"`;
            });
          });
        }
      }

      return `<${type}${attributesStr}>${value}</${type}>`;
    })
    .join("\n");

  return (
    <SyntaxHighlighter language="html" style={docco}>
      {textAreaVal}
    </SyntaxHighlighter>
  );
}

function renderRuleDetail(ruleDetail, rule_id) {
  return (
    <Descriptions key={0} title="Rule" column={1} bordered>
      <Descriptions.Item key={1} label="Name">
        {ruleDetail.name}
      </Descriptions.Item>
      <Descriptions.Item key={2} label="URL Pattern">
        {ruleDetail.url_pattern}
      </Descriptions.Item>
    </Descriptions>
  );
}

function renderParameterTable(path_param) {
  return (
    <table className="ant-table">
      <thead className="ant-table-thead">
        <tr key="id">
          <th>Key</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody className="ant-table-tbody">
        {Object.entries(path_param).map(([key, value]) => (
          <tr key={key}>
            <td>
              <code>{key}</code>
            </td>
            <td>{value}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default SimulationPage;
