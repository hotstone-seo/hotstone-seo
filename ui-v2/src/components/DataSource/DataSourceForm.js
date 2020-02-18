import React from 'react';
import PropTypes from 'prop-types';
import { Form, Input, Button } from 'antd';

function DataSourceForm({ dataSource, handleSubmit }) {
  const [form] = Form.useForm();
  form.setFieldsValue(dataSource);

  return (
    <Form
      form={form}
      onFinish={handleSubmit}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle />

      <Form.Item
        name="name"
        label="Name"
        rules={[{ required: true, message: 'Please input the name of your Data Source' }]}
      >
        <Input placeholder="My Data Source" />
      </Form.Item>

      <Form.Item
        name="url"
        label="Resource URL"
        rules={[{ required: true, message: 'Please input the Resource URL' }]}
      >
        <Input placeholder="http://api.service.com/resource" />
      </Form.Item>

      <Form.Item wrapperCol={{ offset: 6, span: 14 }}>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
}

DataSourceForm.defaultProps = {
  dataSource: {},
};

DataSourceForm.propTypes = {
  dataSource: PropTypes.shape({
    id: PropTypes.number,
    name: PropTypes.string,
    url: PropTypes.string,
  }),

  handleSubmit: PropTypes.func.isRequired,
};

export default DataSourceForm;
