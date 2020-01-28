import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { Layout, Row, Col } from 'antd';
import LoginForm from './LoginForm';

// TODO: Investigate why calling <LoginForm /> cause error saying cannot update state on
// unmounted component
function Login() {
  const [isAuthenticated, setAuthenticated] = useState(false);

  if (isAuthenticated === true) {
    return <Redirect to="/" />
  }

  return (
    <Layout style={{ height: '100vh' }}>
      <Row justify="space-around" align="middle" style={{ height: '100%' }}>
        <Col span={8}>
          <LoginForm setAuthenticated={setAuthenticated} />
        </Col>
      </Row>
    </Layout>
  );
}

export default Login;
