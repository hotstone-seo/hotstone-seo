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
        <Input data-testid="input-src" />
      </Form.Item>
    </>
  );
}

export default ScriptForm;
