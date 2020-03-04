import React from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Select, Button, Tooltip,
} from 'antd';
import { QuestionCircleOutlined } from '@ant-design/icons';

const { Option } = Select;

function RuleForm(props) {
  const {
    rule, dataSources, onSubmit, formLayout,
  } = props;
  const [form] = Form.useForm();
  form.setFieldsValue(rule);

  return (
    <Form
      form={form}
      onFinish={onSubmit}
      layout={formLayout}
      labelCol={
        formLayout === 'horizontal' ? { span: 6 } : null
      }
      wrapperCol={
        formLayout === 'horizontal' ? { span: 14 } : null
      }
    >
      <Form.Item name="id" noStyle />

      <Form.Item
        name="name"
        label="Name"
        rules={[{ required: true, message: 'Please input the name of your Rule' }]}
      >
        <Input placeholder="My Rule" maxLength="200" />
      </Form.Item>

      <Form.Item
        name="url_pattern"
        label="URL Pattern"
        rules={[{ required: true, message: 'Please input the URL Pattern' }]}
        normalize={(value) => {
          if (value[0] !== '/') {
            return `/${value}`;
          }
          return value;
        }}
      >
        <Input
          placeholder="/my/rule/pattern"
          suffix={(
            <Tooltip title="If a request matched this pattern, the tags in the Rule will be returned">
              <QuestionCircleOutlined />
            </Tooltip>
          )}
          maxLength="5000"
        />
      </Form.Item>

      <Form.Item
        name="data_source_id"
        label="Data Source"
      >
        <Select allowClear>
          {dataSources.map(({ id, name }) => (
            <Option value={id} key={id}>{name}</Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item
        wrapperCol={
          formLayout === 'horizontal' ? { offset: 6, span: 14 } : null
        }
      >
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
}

RuleForm.defaultProps = {
  rule: {},
  dataSources: [],
  formLayout: 'horizontal',
};

RuleForm.propTypes = {
  rule: PropTypes.shape({
    id: PropTypes.number,
    name: PropTypes.string,
    url_pattern: PropTypes.string,
    data_source_id: PropTypes.number,
  }),

  dataSources: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number,
      name: PropTypes.string,
    }),
  ),

  formLayout: PropTypes.string,

  onSubmit: PropTypes.func.isRequired,
};

export default RuleForm;
