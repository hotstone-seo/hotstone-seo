import React, { useState } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Select, Button, message,
} from 'antd';
import { addTitleTag, updateTitleTag } from 'api/tag';
import locales from 'locales';
import TagPreview from './TagPreview';

const { Option } = Select;

function TitleForm({ tag, afterSubmit }) {
  const [form] = Form.useForm();
  const [tagPreview, setTagPreview] = useState({ ...tag, type: 'title', attributes: {} });

  const {
    id, rule_id, locale, value,
  } = tag;
  const onSubmit = id ? updateTitleTag : addTitleTag;

  const onFinish = (values) => {
    const formTag = { ...values, id, rule_id };
    onSubmit(formTag)
      .then((response) => {
        afterSubmit(response);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const updateTagPreview = ({ title }) => {
    if (title) {
      setTagPreview({ ...tagPreview, value: title });
    }
  };

  return (
    <Form
      form={form}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 12 }}
      initialValues={{ locale, title: value }}
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
        label="Title"
        name="title"
        rules={[{ required: true, message: 'Must provide a title' }]}
      >
        <Input data-testid="input-title" />
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

TitleForm.defaultProps = {
  tag: {},
};

TitleForm.propTypes = {
  afterSubmit: PropTypes.func.isRequired,
  tag: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number,
    type: PropTypes.string,
    locale: PropTypes.string,
    value: PropTypes.string,
  }),
};

export default TitleForm;
