import React, { useState } from 'react';
import { Form, Input, Checkbox, Button } from 'antd';
import { useAuth } from '../../../components/AuthProvider';

const formLayout = { labelCol: { span: 8 }, wrapperCol: { span: 16 } };

const tailLayout = { wrapperCol: { offset: 8, span: 16 } };

function LoginForm({ setAuthenticated }) {
  const [loading, setLoading] = useState(false);

  const [form] = Form.useForm();

  const auth = useAuth();

  const onFinish = (values) => {
    setLoading(true);
    const { email, password } = values;
    auth.login(email, password)
        .then(() => {
          setAuthenticated(true);
        })
        .finally(() => {
          setLoading(false);
        });
  };

  return (
    <Form {...formLayout} form={form} onFinish={onFinish}>
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
        <Input.Password />
      </Form.Item>

      <Form.Item {...tailLayout} name="remember" valuePropName="checked">
        <Checkbox>Remember me</Checkbox>
      </Form.Item>

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit" loading={loading}>
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
}

export default LoginForm;
