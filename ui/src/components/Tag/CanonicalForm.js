import React, { useState } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Select, Button, message,
} from 'antd';
import { addCanonicalTag, updateCanonicalTag } from 'api/tag';
import locales from 'locales';
import TagPreview from './TagPreview';

const { Option } = Select;

function CanonicalForm({ tag, afterSubmit }) {
  const [form] = Form.useForm();
  const [tagPreview, setTagPreview] = useState({ ...tag, type: 'link', value: null });

  const {
    id, rule_id, locale, attributes = {},
  } = tag;
  const { href } = attributes;
  const onSubmit = id ? updateCanonicalTag : addCanonicalTag;

  const onFinish = (values) => {
    const formTag = Object.assign(values, { id, rule_id });
    onSubmit(formTag)
      .then((response) => {
        afterSubmit(response);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const updateTagPreview = ({ href: formHref }) => {
    const attrs = { ...tagPreview.attributes, rel: 'canonical' };
    if (formHref) {
      attrs.href = formHref;
    }
    setTagPreview({ ...tagPreview, attributes: attrs });
  };

  return (
    <Form
      form={form}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 12 }}
      initialValues={{ locale, href }}
      onFinish={onFinish}
      onValuesChange={updateTagPreview}
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
      <Form.Item label="Preview">
        <TagPreview tag={tagPreview} />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 6, span: 14 }}>
        <Button data-testid="btn-save" type="primary" htmlType="submit">
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

CanonicalForm.defaultProps = {
  tag: {},
};

CanonicalForm.propTypes = {
  afterSubmit: PropTypes.func.isRequired,
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
