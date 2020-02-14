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

  const [currentType, setCurrentType] = useState(tag ? tag.type : tagTypes[0])

  // TODO: Might be a good idea that whenever we change type, the form fields
  // should be cleared to ensure no attributes carried over to new type.

  return (
    <Form {...formLayout} form={form}>
      <Select
        defaultValue={currentType}
        onChange={(value) => setCurrentType(value)}
        style={{ marginBottom: 16 }}
      >
        {tagTypes.map(tagType => (
          <Option value={tagType}>{capitalize(tagType)}</Option>
        ))}
      </Select>
      {
        {
          title: <TitleForm tag={tag} form={form} />,
          meta: <MetaForm tag={tag} form={form} />,
          canonical: null,
          script: null,
        }[currentType]
      }
    </Form>
  );
}

export default TagForm;
