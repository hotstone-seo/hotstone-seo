import React from 'react';
import { Redirect } from 'react-router-dom';
import {
  Layout, Row, Col, message,
} from 'antd';
import { useAuth } from 'components/AuthProvider';
import { LoginForm } from 'components/Login';

function Login() {
  const auth = useAuth();
  if (auth.currentUser) {
    return <Redirect to="/" />;
  }

  const login = (user) => {
    const { email, password } = user;
    auth.login(email, password)
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <Layout style={{ height: '100vh' }}>
      <Row justify="space-around" align="middle" style={{ height: '100%' }}>
        <Col span={8}>
          <LoginForm login={login} />
        </Col>
      </Row>
    </Layout>
  );
}

export default Login;
