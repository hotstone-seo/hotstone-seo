import React from 'react';
import {
  Form,
  Input,
  Checkbox,
  Button
} from 'antd';

const formLayout = { labelCol: { span: 8 }, wrapperCol: { span: 16 } };

const tailLayout = { wrapperCol: { offset: 8, span: 16 } };

function LoginForm() {
  const [form] = Form.useForm();
  return (
    <Form {...formLayout}>
      <Form.Item
        label="Email"
        name="email"
        rules={[{ required: true, message: 'Please input your email!' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[{ required: true, message: 'Please input your password!' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item {...tailLayout} name="remember" valuePropName="checked">
        <Checkbox>Remember me</Checkbox>
      </Form.Item>

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
}

export default LoginForm;
