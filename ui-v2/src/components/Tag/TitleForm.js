import React from 'react';
import { Form, Input, Select } from 'antd';

function TitleForm() {
  return (
    <>
      <Form.Item label="Title" name="value">
        <Input />
      </Form.Item>
    </>
  );
}

export default TitleForm;
