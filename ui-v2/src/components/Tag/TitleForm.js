import React from 'react';
import { Form, Input } from 'antd';

function TitleForm() {
  return (
    <>
      <Form.Item
        label="Title"
        name="value"
        rules={[{ required: true, message: 'Must provide a title' }]}
      >
        <Input />
      </Form.Item>
    </>
  );
}

export default TitleForm;
