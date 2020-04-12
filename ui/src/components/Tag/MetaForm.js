import React, { useState } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Select, Button,
} from 'antd';
import locales from 'locales';
import TagPreview from './TagPreview';

const { Option } = Select;

function MetaForm({ tag, onSubmit }) {
  const [form] = Form.useForm();
  const [tagPreview, setTagPreview] = useState({
    type: 'meta',
    value: null,
    attributes: { name: '', content: '' },
    ...tag,
  });

  const {
    id, rule_id, locale, attributes = {},
  } = tag;
  const { name, content } = attributes;

  const onFinish = (values) => {
    const formTag = { ...values, id, rule_id };
    onSubmit(formTag);
  };

  const updateTagPreview = ({ name: formName, content: formContent }) => {
    const attrs = { ...tagPreview.attributes };
    if (formName) {
      attrs.name = formName;
    }
    if (formContent) {
      attrs.content = formContent;
    }
    setTagPreview({ ...tagPreview, attributes: attrs });
  };

  return (
    <>
      <Form
        form={form}
        labelCol={{ span: 6 }}
        wrapperCol={{ span: 14 }}
        initialValues={{ locale, name, content }}
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
        <Form.Item wrapperCol={{ offset: 6, span: 14 }}>
          <Button data-testid="btn-save" type="primary" htmlType="submit">
            Save
          </Button>
        </Form.Item>
      </Form>
      <TagPreview tag={tagPreview} />
    </>
  );
}

MetaForm.defaultProps = {
  tag: {},
};

MetaForm.propTypes = {
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
