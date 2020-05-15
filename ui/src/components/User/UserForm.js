import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { getRoleType } from 'api/roleType';
import {
  Form, Input, Select, Button,
} from 'antd';

function UserForm({ user, handleSubmit, roleTypes }) {
  const [form] = Form.useForm();
  const [moduleList, setModuleList] = useState([]);

  useEffect(() => {
    setModuleList([]);
    form.setFieldsValue(user);
  }, [user, form]);

  async function handleOnChange(value) {
    try {
      const roleType = await getRoleType(value);
      let mn = null;
      let modulesValue = null;
      let htmlModuleList = '';
      let noModule = 1;
      if (roleType) {
        for (const key in roleType) {
          const value = roleType[key];
          if (key === 'modules') { modulesValue = value; break; }
        }

        Object.keys(modulesValue).forEach((key) => {
          mn = modulesValue[key];
        });

        mn.forEach((item) => {
          htmlModuleList = htmlModuleList.concat(noModule).concat('.').concat(item.name).concat(' ');
          noModule += 1;
        });
      }
      // still show data in common text. TODO : Next time must upper the First Character
      setModuleList(mn);
    } catch (error) {
      console.log(error, 'error');
    }
  }

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
        <Select onChange={handleOnChange}>
          {roleTypes.map(({ id, name }) => (
            <Select.Option value={id} key={id}>{name}</Select.Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item
        name="listMenu"
        label="List of Menu:"
      >
        <Input type="hidden" />
        <div>
          {moduleList.map(({ id, name }) => (
            <span id={id} key={name}>
              {name}
              {' '}
              <br />
            </span>
          ))}

        </div>
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
