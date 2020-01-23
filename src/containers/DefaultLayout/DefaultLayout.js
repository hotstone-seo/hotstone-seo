import React from 'react';
import { Layout } from 'antd';

const { Header, Content, Sider } = Layout;

function DefaultLayout() {
  return (
    <Layout>
      <Header/>
      <Layout>
        <Sider>
        </Sider>
        <Layout>
          <Content>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  );
}

export default DefaultLayout;
