import React from 'react';
import { Form, Input } from 'antd';

function CanonicalForm() {
  return (
    <>
      <Form.Item label="Rel" name={['attributes', 'rel']}>
        <Input defaultValue="canonical" />
      </Form.Item>

      <Form.Item label="URL" name={['attributes', 'href']}>
        <Input />
      </Form.Item>
    </>
  );
}

export default CanonicalForm;
