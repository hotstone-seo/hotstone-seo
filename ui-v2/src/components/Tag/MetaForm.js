import React from 'react';
import { Form, Input, Select } from 'antd';

const { Option } = Select;

function MetaForm() {
  return (
    <>
      <Form.Item label="Name" name={['attributes', 'name']}>
        <Input />
      </Form.Item>

      <Form.Item label="Content" name={['attributes', 'content']}>
        <Input />
      </Form.Item>
    </>
  );
}

export default MetaForm;
