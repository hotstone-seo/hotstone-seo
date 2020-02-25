import React from 'react';
import PropTypes from 'prop-types';
import { Form, Input } from 'antd';

function CanonicalForm({ form }) {
  form.setFieldsValue({
    attributes: { rel: 'canonical' },
  });

  return (
    <>
      <Form.Item name={['attributes', 'rel']} noStyle />

      <Form.Item label="URL" name={['attributes', 'href']}>
        <Input />
      </Form.Item>
    </>
  );
}

CanonicalForm.propTypes = {
  form: PropTypes.shape({
    setFieldsValue: PropTypes.func.isRequired,
  }).isRequired,
};

export default CanonicalForm;