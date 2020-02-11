import React from 'react';
import { Menu, Dropdown, Button } from 'antd';
import { UserOutlined, LogoutOutlined } from '@ant-design/icons';
import { useAuth } from '../../components/AuthProvider';

function HeaderMenu() {
  const auth = useAuth();
  const { name } = auth.currentUser;

  const accountMenu = (
    <Menu>
      <Menu.Item>
        <Button
          type="link"
          icon={<LogoutOutlined />}
          onClick={() => auth.logout()}
        >
          Logout
        </Button>
      </Menu.Item>
    </Menu>
  );

  return (
    <>
      <Dropdown overlay={accountMenu}>
        <Button type="link" icon={<UserOutlined />}>
          {name}
        </Button>
      </Dropdown>
    </>
  );
}

export default HeaderMenu;
