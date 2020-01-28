import React from 'react';
import { Layout, Row, Col } from 'antd';
import LoginForm from './LoginForm';

function Login() {
  return (
    <Layout style={{ height: '100vh' }}>
      <Row justify="space-around" align="middle" style={{ height: '100%' }}>
        <Col span={8}>
          <LoginForm />
        </Col>
      </Row>
    </Layout>
  );
}

export default Login;
