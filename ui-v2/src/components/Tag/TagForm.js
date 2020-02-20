import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { Form, Select } from 'antd';
import locales from 'locales';

import TitleForm from './TitleForm';
import MetaForm from './MetaForm';
import CanonicalForm from './CanonicalForm';

const { Option } = Select;

const tagTypes = [
  { label: 'Title', value: 'title' },
  { label: 'Meta', value: 'meta' },
  { label: 'Canonical', value: 'link' },
  { label: 'Script', value: 'script' },
];

function TagForm({ form }) {
  const type = form.getFieldValue('type');
  const [currentType, setCurrentType] = useState(type);

  return (
    <Form
      form={form}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle />

      <Form.Item name="rule_id" noStyle />

      <Form.Item label="Type" name="type">
        <Select onChange={(value) => setCurrentType(value)}>
          {tagTypes.map(({ label, value }) => (
            <Option key={value} value={value}>{label}</Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item label="Locale" name="locale">
        <Select>
          {locales.map((locale) => (
            <Option key={locale} value={locale}>{locale}</Option>
          ))}
        </Select>
      </Form.Item>
      {
        {
          title: <TitleForm form={form} />,
          meta: <MetaForm form={form} />,
          link: <CanonicalForm form={form} />,
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
