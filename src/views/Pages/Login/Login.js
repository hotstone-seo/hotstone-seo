import React, { useState } from 'react';
import { Layout, Row, Col } from 'antd';
import { useAuth } from '../../../components/AuthProvider';
import LoginForm from './LoginForm';

// TODO: Investigate why calling <LoginForm /> cause error saying cannot update state on
// unmounted component
function Login() {
  const auth = useAuth();

  return (
    <Layout style={{ height: '100vh' }}>
      <Row justify="space-around" align="middle" style={{ height: '100%' }}>
        <Col span={8}>
          <LoginForm login={auth.login} />
        </Col>
      </Row>
    </Layout>
  );
}

export default Login;
