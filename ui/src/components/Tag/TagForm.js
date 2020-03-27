import React, { useState } from 'react';
import { renderToStaticMarkup } from 'react-dom/server';
import PropTypes from 'prop-types';
import { Form, Select } from 'antd';
import locales from 'locales';

import TitleForm from './TitleForm';
import MetaForm from './MetaForm';
import CanonicalForm from './CanonicalForm';
import ScriptForm from './ScriptForm';

const { Option } = Select;

const tagTypes = [
  { label: 'Title', value: 'title' },
  { label: 'Meta', value: 'meta' },
  { label: 'Canonical', value: 'link' },
  { label: 'Script', value: 'script' },
];

function TagForm({ form }) {
  const [currentType, setCurrentType] = useState(form.getFieldValue('type'));
  const [tagValues, setTagValues] = useState(
    form.getFieldsValue(['type', 'attributes', 'value']),
  );

  const renderTagMock = (tag) => {
    if (!tag.type) {
      return null;
    }
    const mock = tag;
    if (['meta', 'link'].includes(mock.type)) {
      mock.value = null;
    }
    return (
      <pre data-testid="text-preview-tag">
        {renderToStaticMarkup(
          React.createElement(mock.type, mock.attributes, mock.value),
        )}
      </pre>
    );
  };

  return (
    <>
      <Form
        form={form}
        labelCol={{ span: 6 }}
        wrapperCol={{ span: 14 }}
        onValuesChange={(changedValues, allValues) => setTagValues(allValues)}
      >
        <Form.Item name="id" noStyle />

        <Form.Item name="rule_id" noStyle />

        <Form.Item
          label="Type"
          name="type"
          rules={[{ required: true, message: 'A type must be selected' }]}
        >
          <Select
            data-testid="select-type"
            onChange={(value) => setCurrentType(value)}
            showSearch
            filterOption={(input, option) => option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0}
          >
            {tagTypes.map(({ label, value }) => (
              <Option key={value} value={value}>{label}</Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item
          label="Locale"
          name="locale"
          rules={[{ required: true, message: 'Please set a locale for the tag' }]}
        >
          <Select
            data-testid="select-locale"
            showSearch
            filterOption={(input, option) => option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0}
          >
            {locales.map((locale) => (
              <Option key={locale} value={locale}>{locale}</Option>
            ))}
          </Select>
        </Form.Item>
        {
          {
            title: <TitleForm />,
            meta: <MetaForm />,
            link: <CanonicalForm form={form} />,
            script: <ScriptForm />,
          }[currentType]
        }
      </Form>
      {renderTagMock(tagValues)}
    </>
  );
}

TagForm.propTypes = {
  form: PropTypes.shape({
    getFieldValue: PropTypes.func.isRequired,
    getFieldsValue: PropTypes.func.isRequired,
  }).isRequired,
};

export default TagForm;
