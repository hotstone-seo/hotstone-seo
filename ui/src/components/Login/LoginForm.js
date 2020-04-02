import React, { useRef } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Button, Row, Divider,
} from 'antd';
import { GoogleOutlined } from '@ant-design/icons';

function LoginForm({ login }) {
  const [form] = Form.useForm();
  const loginGoogleForm = useRef();

  return (
    <>
      <Form
        form={form}
        onFinish={login}
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
      >
        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Row>
            <Divider type="vertical" />
            <Button
              type="primary"
              icon={<GoogleOutlined />}
              onClick={() => loginGoogleForm.current.submit()}
            >
              Sign in with Google
            </Button>
          </Row>
        </Form.Item>
      </Form>
      <form ref={loginGoogleForm} action="/auth/google/login" method="post" />
    </>
  );
}

LoginForm.propTypes = {
  login: PropTypes.func.isRequired,
};

export default LoginForm;
