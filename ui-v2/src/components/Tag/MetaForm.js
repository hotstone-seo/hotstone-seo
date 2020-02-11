import React from 'react';
import { Form, Input, Select } from 'antd';

const { Option } = Select;

function MetaForm({ locales, tag }) {
  const [form] = Form.useForm();
  if (tag) {
    form.setFieldsValue(tag);
  }

  return (
    <Form>
      <Form.Item name="locale">
        <Select>
          {locales.map(locale => (
            <Option value={locale}>{locale}</Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item name={['attributes', 'name']}>
        <Input />
      </Form.Item>
      <Form.Item name={['attributes', 'content']}>
        <Input />
      </Form.Item>
    </Form>
  );
}

export default MetaForm;
