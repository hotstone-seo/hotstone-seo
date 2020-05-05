import React, { useState, useEffect } from 'react';
import {
  Link, Route, Switch, useLocation, Redirect
} from 'react-router-dom';
import { Layout, Menu, message } from 'antd';
import routes from 'routes';
import logo from 'assets/hotstone-logo.png';
import miniLogo from 'assets/hotstone-logo-mini.png';
import HeaderMenu from './HeaderMenu';
import styles from './DashboardLayout.module.css';

const { Header, Content, Sider } = Layout;

function DashboardLayout() {
  const location = useLocation();
  const [collapsed] = useState(false);
  const [broken, setBroken] = useState(false);

  useEffect(() => {
    if (location.state && location.state.message) {
      const { level, content } = location.state.message;
      message[level](content);
    }
  }, [location.state]);

  return (
    <Layout style={{ height: '100vh' }}>
      <Sider
        breakpoint="lg"
        onBreakpoint={(brokenNewVal) => {
          setBroken(brokenNewVal);
        }}
        onCollapse={collapsed}
      >
        <div
          className={styles.logo}
          style={{
            backgroundImage: (broken ? `url(${miniLogo})` : `url(${logo})`),
            backgroundSize: (broken ? '50px 32px' : '100px 32px'),
          }}
        />
        <Menu theme="dark" mode="inline" defaultSelectedKeys={[location.pathname]}>
          {routes.map(({ path, name, icon }) => (
            <Menu.Item key={path}>
              {icon && React.createElement(icon)}
              <span>{name}</span>
              <Link to={path} />
            </Menu.Item>
          ))}
        </Menu>
      </Sider>
      <Layout>
        <Header className={styles.header} style={{ padding: 0 }}>
          <div className={styles.headerRight}>
            <HeaderMenu />
          </div>
        </Header>
        <Content className={styles.content}>
          <Switch>
            <Redirect exact from="/" to="/rules" />
            {routes.map((route) => (route.component ? (
              <Route
                key={route.path}
                path={route.path}
                exact={route.exact}
                name={route.name}
                component={route.component}
              />
            ) : null))}
          </Switch>
        </Content>
      </Layout>
    </Layout>
  );
}

export default DashboardLayout;
