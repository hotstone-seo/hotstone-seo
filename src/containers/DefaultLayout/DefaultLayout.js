import React from 'react';
import { Link, Route, Switch } from 'react-router-dom';
import { Layout, Menu } from 'antd';
import routes from '../../routes';

const { Header, Content, Sider } = Layout;

function DefaultLayout(props) {
  return (
    <Layout>
      <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }} />
      <Layout>
        <Sider>
          <Menu>
            {routes.map((route, idx) => (
              <Menu.Item key={idx}>
                <Link to={route.path}>{route.name}</Link>
              </Menu.Item>
            ))}
          </Menu>
        </Sider>
        <Layout>
          <Content>
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
    </Layout>
  );
}

export default DefaultLayout;
