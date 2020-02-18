import React from 'react';
import { Form, Input, Select } from 'antd';
import locales from 'locales';

const { Option } = Select;

function TitleForm() {
  return (
    <>
      <Form.Item label="Locale" name="locale">
        <Select>
          {locales.map((locale) => (
            <Option value={locale}>{locale}</Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item label="Title" name="value">
        <Input />
      </Form.Item>
    </>
  );
}

export default TitleForm;
