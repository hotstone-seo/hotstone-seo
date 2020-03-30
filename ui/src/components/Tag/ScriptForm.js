import React from 'react';
import PropTypes from 'prop-types';
import { Form, Input, Select } from 'antd';
import locales from 'locales';

const { Option } = Select;

function ScriptForm({ form, tag, onSubmit }) {
  const { id, rule_id, locale, attributes: { src } } = tag;

  const onFinish = (values) => {
    const formTag = Object.assign(values, { id, rule_id });
    onSubmit(formTag);
  };

  form.setFieldsValue({ locale, source: src });

  return (
    <Form
      form={form}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
      onFinish={onFinish}
    >
      <Form.Item
        label="Locale"
        name="locale"
        rules={[{ required: true, message: 'Please set a locale for the tag' }]}
      >
        <Select
          data-testid="select-locale"
          showSearch
          filterOption={(input, option) => (
            option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
          )}
        >
          {locales.map((loc) => (
            <Option key={loc} value={loc}>{loc}</Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item
        label="Source"
        name="source"
        rules={[{ required: true, message: 'Must provide script URL' }]}
      >
        <Input data-testid="input-src" />
      </Form.Item>
    </Form>
  );
}

ScriptForm.defaultProps = {
  form: Form.useForm(),
  tag: {},
};

ScriptForm.propTypes = {
  form: PropTypes.shape({
    setFieldsValue: PropTypes.func,
  }),
  onSubmit: PropTypes.func.isRequired,
  tag: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number,
    type: PropTypes.string,
    locale: PropTypes.string,
    attributes: PropTypes.shape({
      src: PropTypes.string,
    }),
  }),
};


export default ScriptForm;
