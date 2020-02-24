import React from 'react';
import { Form, Input } from 'antd';

function ScriptForm() {
  return (
    <>
      <Form.Item label="Source" name={['attributes', 'src']}>
        <Input />
      </Form.Item>
    </>
  );
}

export default ScriptForm;
