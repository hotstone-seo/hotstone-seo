import React, { useRef } from 'react';
import { Menu, Dropdown, Button } from 'antd';
import { UserOutlined, LogoutOutlined } from '@ant-design/icons';
import { useAuth } from '../../components/AuthProvider';

function HeaderMenu() {
  const auth = useAuth();
  const { email, role } = auth.currentUser;

  const logoutForm = useRef();

  const accountMenu = (
    <Menu>
      <Menu.Item>
        <Button
          type="link"
          icon={<LogoutOutlined />}
          onClick={() => logoutForm.current.submit()}
        >
          Logout
        </Button>
      </Menu.Item>
    </Menu>
  );

  return (
    <>
      Login as : &nbsp;
      {role}
      <Dropdown overlay={accountMenu}>
        <Button type="link" icon={<UserOutlined />}>
          {email}
        </Button>
      </Dropdown>
      <form ref={logoutForm} action="/api/logout" method="post" />
    </>
  );
}

export default HeaderMenu;
