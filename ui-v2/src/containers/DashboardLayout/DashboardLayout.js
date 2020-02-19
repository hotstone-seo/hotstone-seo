import React, { useState } from 'react';
import { Link, Route, Switch } from 'react-router-dom';
import { Layout, Menu } from 'antd';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons';
import HeaderMenu from './HeaderMenu';
import routes from 'routes';
import styles from './DashboardLayout.module.css';

const { Header, Content, Sider } = Layout;

function DashboardLayout(props) {
  const [collapsed, setCollapsed] = useState(false);
  const { location } = props;
  return (
    <Layout className={styles.base}>
      <Sider
        className={styles.sider}
        trigger={null}
        collapsible
        collapsed={collapsed}
        width={256}
      >
        <div className={styles.logo} />
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
          {React.createElement(collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
            className: styles.trigger,
            onClick: () => { setCollapsed(!collapsed) },
          })}
          <div className={styles.headerRight}>
            <HeaderMenu />
          </div>
        </Header>
        <Content className={styles.content}>
          <Switch>
            {routes.map((route, idx) => {
              return route.component ? (
                <Route
                  key={idx}
                  path={route.path}
                  exact={route.exact}
                  name={route.name}
                  render={props => <route.component {...props} />}
                />
              ) : null;
            })}
          </Switch>
        </Content>
      </Layout>
    </Layout>
  );
}

export default DashboardLayout;
