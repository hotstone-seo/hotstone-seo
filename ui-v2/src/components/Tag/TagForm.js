import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { Form, Select } from 'antd';
import TitleForm from './TitleForm';
import MetaForm from './MetaForm';

const { Option } = Select;

const tagTypes = ['title', 'meta', 'canonical', 'script'];

const capitalize = (item) => item.charAt(0).toUpperCase() + item.slice(1);

function TagForm({ form }) {
  const [currentType, setCurrentType] = useState(tagTypes[0]);

  useEffect(() => {
    const type = form.getFieldValue('type');
    if (type) {
      setCurrentType(type);
    }
  }, [form]);

  return (
    <Form
      form={form}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle />

      <Form.Item name="rule_id" noStyle />

      <Form.Item label="Type" name="type">
        <Select
          defaultValue={currentType}
          onChange={(value) => setCurrentType(value)}
        >
          {tagTypes.map((tagType) => (
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

TagForm.propTypes = {
  form: PropTypes.shape({
    getFieldValue: PropTypes.func.isRequired,
  }).isRequired,
};

export default TagForm;
