import React, { useState } from 'react';
import { Form, Select } from 'antd';
import TitleForm from './TitleForm';
import MetaForm from './MetaForm';

const { Option } = Select;

const formLayout = { labelCol: { span: 6 }, wrapperCol: { span: 14 } };

const tagTypes = ['title', 'meta', 'canonical', 'script'];

const capitalize = (item) => {
  return item.charAt(0).toUpperCase() + item.slice(1);
}

function TagForm({ tag, form }) {
  if (!form) { [form] = Form.useForm(); }

  const type = form.getFieldValue('type');
  const [currentType, setCurrentType] = useState(type || tagTypes[0]);

  return (
    <Form {...formLayout} form={form}>
      <Form.Item name="id" noStyle />

      <Form.Item name="rule_id" noStyle />

      <Form.Item label="Type" name="type">
        <Select
          defaultValue={currentType}
          onChange={(value) => setCurrentType(value)}
        >
          {tagTypes.map(tagType => (
            <Option value={tagType}>{capitalize(tagType)}</Option>
          ))}
        </Select>
      </Form.Item>
      {
        {
          title: <TitleForm form={form} />,
          meta: <MetaForm form={form} />,
          canonical: null,
          script: null,
        }[currentType]
      }
    </Form>
  );
}

export default TagForm;
