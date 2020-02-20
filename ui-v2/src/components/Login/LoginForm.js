import React from "react";
import PropTypes from "prop-types";
import { Form, Input, Checkbox, Button, Row, Divider, Typography } from "antd";
import { GoogleOutlined } from "@ant-design/icons";
import urljoin from "url-join";

const { Text } = Typography;

function LoginForm({ login }) {
  const [form] = Form.useForm();

  return (
    <>
      <Form
        form={form}
        onFinish={login}
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
      >
        <Form.Item
          label="Email"
          name="email"
          rules={[{ required: true, message: "Please input your email!" }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: "Please input your password!" }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item
          name="remember"
          valuePropName="checked"
          wrapperCol={{ offset: 8, span: 16 }}
        >
          <Checkbox>Remember me</Checkbox>
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Row>
            <Button type="primary" htmlType="submit">
              Submit
            </Button>
            <Divider type="vertical" />
            <Text strong>OR</Text>
            <Divider type="vertical" />
            <Button
              type="primary"
              icon={<GoogleOutlined />}
              href={urljoin(
                process.env.REACT_APP_API_URL,
                "/auth/google/login"
              )}
            >
              Sign in with Google
            </Button>
          </Row>
        </Form.Item>
      </Form>
    </>
  );
}

LoginForm.propTypes = {
  login: PropTypes.func.isRequired
};

export default LoginForm;
