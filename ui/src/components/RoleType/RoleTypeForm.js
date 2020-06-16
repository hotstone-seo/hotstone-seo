import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { Form, Input, Button } from "antd";
import { createRoleType, updateRoleType } from "api/roleType";

const { TextArea } = Input;

function RoleTypeForm({ roleType, handleSubmit }) {
  const [form] = Form.useForm();

  useEffect(() => {
    let menus = "";
    let paths = "";
    if (roleType.menus !== undefined) {
      menus = roleType.menus.join("\n");
    }
    if (roleType.paths !== undefined) {
      paths = roleType.paths.join("\n");
    }

    form.setFieldsValue({
      id: roleType.id,
      name: roleType.name,
      menus: menus,
      paths: paths,
    });
  }, [roleType, form]);

  const onSubmit = roleType.id === undefined ? createRoleType : updateRoleType;

  const onFinish = (values) => {
    onSubmit(values)
      .then((response) => {
        handleSubmit(response);
      })
      .catch(() => {});
  };

  return (
    <Form
      form={form}
      onFinish={onFinish}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle>
        <Input type="hidden" />
      </Form.Item>

      <Form.Item
        name="name"
        label="Role Name"
        rules={[{ required: true, message: "Please input the role" }]}
      >
        <Input
          data-testid="input-role-type"
          placeholder="Role Name"
          maxLength="200"
        />
      </Form.Item>

      <Form.Item
        name="menus"
        label="Menus"
        rules={[{ required: true, message: "Please input the menus" }]}
      >
        <TextArea rows={6} />
      </Form.Item>

      <Form.Item
        name="paths"
        label="Paths"
        rules={[{ required: true, message: "Please input the paths" }]}
      >
        <TextArea rows={6} />
      </Form.Item>

      <Form.Item wrapperCol={{ offset: 6, span: 14 }}>
        <Button data-testid="btn-save" type="primary" htmlType="submit">
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

RoleTypeForm.defaultProps = {
  roleType: {},
};

RoleTypeForm.propTypes = {
  roleType: PropTypes.shape({
    id: PropTypes.number,
    name: PropTypes.string,
  }),

  handleSubmit: PropTypes.func.isRequired,
};

export default RoleTypeForm;
