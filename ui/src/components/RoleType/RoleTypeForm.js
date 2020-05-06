import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button,
} from 'antd';

function RoleTypeForm({ roleType, handleSubmit }) {
  const [form] = Form.useForm();

  useEffect(() => {
    form.setFieldsValue(roleType);
  }, [roleType, form]);

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
        name="name"
        label="Role"
        rules={[{ required: true, message: 'Please input the role' }]}
      >
        <Input data-testid="input-role" placeholder="Role" maxLength="200" />
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
