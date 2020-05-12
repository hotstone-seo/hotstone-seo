import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button, Checkbox,
} from 'antd';

const CheckboxGroup = Checkbox.Group;
const plainOptions = ['Rules', 'Data Sources', 'Mismatched Rule', 'Analytic', 'Simulation', 'Audit Trail', 'User', 'Role User'];
const defaultCheckedList = ['Mismatched Rule', 'Analytic', 'Simulation'];

function RoleTypeForm({ roleType, handleSubmit }) {
  const [form] = Form.useForm();
  const [checkedList, setCheckedList] = useState({
    checkedList: defaultCheckedList,
    indeterminate: true,
    checkAll: false,
    checkboxArray: [],
  });

  useEffect(() => {
    form.setFieldsValue(roleType);
  }, [roleType, form]);

  const handleonChange = (checkedListNew) => {
    if (checkedListNew) {
      setCheckedList({
        ...checkedList,
        checkedList: checkedListNew,
        indeterminate: !!checkedListNew.length && checkedListNew.length < plainOptions.length,
        checkAll: checkedListNew.length === plainOptions.length,
      });
    }
  };

  const onCheckAllChange = (e) => {
    setCheckedList({
      checkedList: e.target.checked ? plainOptions : [],
      indeterminate: false,
      checkAll: e.target.checked,
    });
  };

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
        <Input data-testid="input-role-type" placeholder="Role Name" maxLength="200" />
      </Form.Item>

      <Form.Item
        name="modules"
        label="Module Access"
        rules={[{ required: true, message: 'Please check module access' }]}
      >
        <div>
          <div className="site-checkbox-all-wrapper">
            <Checkbox
              indeterminate={checkedList.indeterminate}
              onChange={onCheckAllChange}
              checked={checkedList.checkAll}
            >
              Check all
            </Checkbox>
          </div>
          <br />
          <CheckboxGroup
            options={plainOptions}
            value={checkedList.checkedList}
            onChange={handleonChange}
            name="modules"
          />
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
