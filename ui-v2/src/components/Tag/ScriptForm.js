import React from 'react';
import { Form, Input } from 'antd';

function ScriptForm() {
  return (
    <>
      <Form.Item
        label="Source"
        name={['attributes', 'src']}
        rules={[{ required: true, message: 'Must provide script URL' }]}
      >
        <Input />
      </Form.Item>
    </>
  );
}

export default ScriptForm;
