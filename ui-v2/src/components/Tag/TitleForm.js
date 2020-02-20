import React from 'react';
import { Form, Input, Select } from 'antd';
import locales from 'locales';

const { Option } = Select;

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
