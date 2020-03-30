import React from 'react';
import PropTypes from 'prop-types';
import { Form, Input, Select } from 'antd';
import locales from 'locales';

const { Option } = Select;

function MetaForm({ form, tag, onSubmit }) {
  const { id, rule_id, locale, attributes: { name, content } } = tag;

  const onFinish = (values) => {
    const formTag = Object.assign(values, { id, rule_id });
    onSubmit(formTag);
  };

  form.setFieldsValue({ locale, name, content });

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
        label="Name"
        name="name"
        rules={[{ required: true, message: 'Must provide meta name' }]}
      >
        <Input data-testid="input-name" />
      </Form.Item>
      <Form.Item
        label="Content"
        name="content"
        rules={[{ required: true, message: 'Must provide meta content' }]}
      >
        <Input data-testid="input-content" />
      </Form.Item>
    </Form>
  );
}

MetaForm.defaultProps = {
  form: Form.useForm(),
  tag: {},
};

MetaForm.propTypes = {
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
      name: PropTypes.string,
      content: PropTypes.string,
    }),
  }),
};

export default MetaForm;
