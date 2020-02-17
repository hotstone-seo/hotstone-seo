import React from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Select, Button,
} from 'antd';

const { Option } = Select;

function RuleForm({ rule, dataSources, handleSubmit }) {
  const [form] = Form.useForm();
  form.setFieldsValue(rule);

  return (
    <Form
      form={form}
      onFinish={handleSubmit}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle />

      <Form.Item
        name="name"
        label="Name"
        rules={[{ required: true, message: 'Please input the name of your Rule' }]}
      >
        <Input placeholder="My Rule" />
      </Form.Item>

      <Form.Item
        name="url_pattern"
        label="URL Pattern"
        rules={[{ required: true, message: 'Please input the URL Pattern' }]}
      >
        <Input placeholder="/my/rule/pattern" />
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

      <Form.Item wrapperCol={{ offset: 6, span: 14 }}>
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

  handleSubmit: PropTypes.func.isRequired,
};

export default RuleForm;
