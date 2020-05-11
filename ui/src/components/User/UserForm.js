import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Select, Button,
} from 'antd';

function UserForm({ user, handleSubmit, roleTypes }) {
  const [form] = Form.useForm();

  useEffect(() => {
    form.setFieldsValue(user);
  }, [user, form]);

  return (
    <Form
      form={form}
      onFinish={handleSubmit}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle>
        <Input type="hidden" />
      </Form.Item>

      <Form.Item
        name="email"
        label="Email"
        rules={[{ required: true, message: 'Please input the email' }, { type: 'email', message: 'Please input the valid email' }]}
      >
        {user.email === undefined ? (
          <Input data-testid="input-email" placeholder="Email" maxLength="200" />
        ) : (
          <>
            {user.email}
            <Input type="hidden" />
          </>
        )}
      </Form.Item>

      <Form.Item
        name="role_type_id"
        label="Role User"
      >
        <Select>
          {roleTypes.map(({ id, name }) => (
            <Select.Option value={id} key={id}>{name}</Select.Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item
        wrapperCol={{ offset: 6, span: 14 }}
      >
        <Button data-testid="btn-save" type="primary" htmlType="submit">
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

UserForm.defaultProps = {
  user: {},
  roleTypes: [],
};

UserForm.propTypes = {
  user: PropTypes.shape({
    id: PropTypes.number,
    email: PropTypes.string,
    role_type_id: PropTypes.number,
  }),

  roleTypes: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number,
      name: PropTypes.string,
    }),
  ),

  handleSubmit: PropTypes.func.isRequired,
};

export default UserForm;
