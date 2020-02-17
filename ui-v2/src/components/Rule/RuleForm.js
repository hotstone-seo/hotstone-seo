import React from 'react';
import { Form, Input, Select, Button } from 'antd';

const { Option } = Select;

const formLayout = { labelCol: { span: 6 }, wrapperCol: { span: 14 } };

const tailLayout = { wrapperCol: { offset: 6, span: 14 } };

function RuleForm({ handleSubmit, rule, dataSources }) {
  const [form] = Form.useForm();
  if (rule) {
    form.setFieldsValue(rule);
  }

  return (
    <Form {...formLayout} form={form} onFinish={handleSubmit}>
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

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
}

export default RuleForm;
