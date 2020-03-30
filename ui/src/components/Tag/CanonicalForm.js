import React from 'react';
import PropTypes from 'prop-types';
import { Form, Input, Select } from 'antd';
import locales from 'locales';

const { Option } = Select;

function CanonicalForm({ form, tag, onSubmit }) {
  const { id, rule_id, locale, attributes: { href } } = tag;

  const onFinish = (values) => {
    const formTag = Object.assign(values, { id, rule_id });
    onSubmit(formTag);
  };

  form.setFieldsValue({ locale, href });

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
        label="Href"
        name="href"
        rules={[{ required: true, message: 'Must provide canonical URL' }]}
      >
        <Input data-testid="input-url" />
      </Form.Item>
    </Form>
  );
}

CanonicalForm.defaultProps = {
  form: Form.useForm(),
  tag: {},
};

CanonicalForm.propTypes = {
  form: PropTypes.shape({
    setFieldsValue: PropTypes.func.isRequired,
  }),
  onSubmit: PropTypes.func.isRequired,
  tag: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number,
    type: PropTypes.string,
    locale: PropTypes.string,
    attributes: PropTypes.shape({
      href: PropTypes.string,
    }),
  }),
};

export default CanonicalForm;
