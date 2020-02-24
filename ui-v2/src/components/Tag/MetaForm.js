import React from 'react';
import { Form, Input } from 'antd';

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
