import React, { useState } from 'react';
import { Link, Route, Switch } from 'react-router-dom';
import { Layout, Menu } from 'antd';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons';
import HeaderMenu from './HeaderMenu';
import routes from '../../routes';
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
          {routes.map(({ path, name }) => (
            <Menu.Item key={path}>
              <Link to={path}>{name}</Link>
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
