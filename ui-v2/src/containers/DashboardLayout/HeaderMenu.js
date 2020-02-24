import React, { useRef } from "react";
import { Menu, Dropdown, Button } from "antd";
import { UserOutlined, LogoutOutlined } from "@ant-design/icons";
import { useAuth } from "../../components/AuthProvider";
import urljoin from "url-join";

function HeaderMenu() {
  const auth = useAuth();
  const { email } = auth.currentUser;

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
      <Dropdown overlay={accountMenu}>
        <Button type="link" icon={<UserOutlined />}>
          {email}
        </Button>
      </Dropdown>
      <form
        ref={logoutForm}
        action={urljoin(process.env.REACT_APP_API_URL, "/auth/logout")}
        method="post"
        onSubmit={env => {
          console.log("== Form submitted ==");
          // alert("Form submitted!");
        }}
      ></form>
    </>
  );
}

export default HeaderMenu;
