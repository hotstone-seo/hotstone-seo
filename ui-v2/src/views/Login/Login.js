import React from "react";
import { Redirect } from "react-router-dom";
import { Layout, Row, Col, message, Alert } from "antd";
import { useAuth } from "components/AuthProvider";
import { LoginForm } from "components/Login";
import useRouter from "hooks/useRouter";

function Login() {
  const { query } = useRouter();

  const auth = useAuth();
  if (auth.currentUser) {
    return <Redirect to="/" />;
  }

  const login = user => {
    const { email, password } = user;
    auth.login(email, password).catch(error => {
      message.error(error.message);
    });
  };

  return (
    <Layout style={{ height: "100vh" }}>
      <Row justify="space-around" align="middle" style={{ height: "100%" }}>
        <Col span={8}>
          {query.oauth_error && (
            <Alert
              type="error"
              message="Failed to login"
              closable
              showIcon
              style={{ marginBottom: 10 }}
            />
          )}
          <LoginForm login={login} />
        </Col>
      </Row>
    </Layout>
  );
}

export default Login;
