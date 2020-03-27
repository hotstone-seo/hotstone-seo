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
        <Input data-testid="input-name" />
      </Form.Item>

      <Form.Item
        label="Content"
        name={['attributes', 'content']}
        rules={[{ required: true, message: 'Must provide meta content' }]}
      >
        <Input data-testid="input-content" />
      </Form.Item>
    </>
  );
}

export default MetaForm;
