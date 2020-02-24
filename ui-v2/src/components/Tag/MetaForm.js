import React from 'react';
import { Form, Input } from 'antd';

function MetaForm() {
  return (
    <>
      <Form.Item
        label="Name"
        name={['attributes', 'name']}
        rules={[{ required: true, message: 'Must provide meta name' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Content"
        name={['attributes', 'content']}
        rules={[{ required: true, message: 'Must provide meta content' }]}
      >
        <Input />
      </Form.Item>
    </>
  );
}

export default MetaForm;
