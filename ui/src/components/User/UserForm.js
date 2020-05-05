import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Select, Button,
} from 'antd';

function UserForm(props) {
  const {
    user, roleTypes, onSubmit, formLayout,
  } = props;
  const [form] = Form.useForm();

  useEffect(() => {
    form.setFieldsValue(user);
  }, [user, form]);

  return (
    <Form
      form={form}
      onFinish={onSubmit}
      layout={formLayout}
      labelCol={
        formLayout === 'horizontal' ? { span: 6 } : null
      }
      wrapperCol={
        formLayout === 'horizontal' ? { span: 14 } : null
      }
    >
      <Form.Item name="id" noStyle>
        <Input type="hidden" />
      </Form.Item>

      <Form.Item
        name="email"
        label="Email"
        rules={[{ required: true, message: 'Please input the email' }]}
      >
        <Input data-testid="input-email" placeholder="Email" maxLength="200" />
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
        wrapperCol={
          formLayout === 'horizontal' ? { offset: 6, span: 14 } : null
        }
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
  formLayout: 'horizontal',
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

  formLayout: PropTypes.string,

  onSubmit: PropTypes.func.isRequired,
};

export default UserForm;
