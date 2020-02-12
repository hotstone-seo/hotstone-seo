import React from 'react';
import { Form, Input, Select } from 'antd';
import { locales } from 'constants';

const { Option } = Select;

function TitleForm({ tag }) {
  const [form] = Form.useForm();
  if (tag) {
    form.setFieldsValue(tag);
  }
  form.setFieldsValue({ type: 'title' })

  return (
    <Form form={form}>
      <Form.Item label="Locale" name="locale">
        <Select>
          {locales.map(locale => (
            <Option value={locale}>{locale}</Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item label="Title" name="value">
        <Input />
      </Form.Item>
    </Form>
  )
}

export default TitleForm;
