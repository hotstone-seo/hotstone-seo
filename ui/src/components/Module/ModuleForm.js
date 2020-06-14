import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button, message,
} from 'antd';
import { createModule , updateModule } from 'api/module';

function ModuleForm({ module, handleSubmit }) {
  const [form] = Form.useForm();
  const { id } = module;
  useEffect(() => {
    form.setFieldsValue(module);
  }, [module, form]);

  const onFinish = (values) => {
    const formStruct = { ...values, id };
    const submit = id ? updateModule : createModule;
    submit(formStruct)
      .then((response) => {
        response.name = values.name;
        handleSubmit(response);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <Form
      form={form}
      initialValues={{ ...module }}
      onFinish={onFinish}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle>
        <Input type="hidden" />
      </Form.Item>

      <Form.Item
        name="name"
        label="Name"
        rules={[{ required: true, message: 'Please input the module name' }]}
      >
        <Input data-testid="input-name" placeholder="Name" maxLength="60" />
      </Form.Item>

      <Form.Item
        name="path"
        label="Front End Path"
        rules={[{ required: true, message: 'Please input the Front End Path' }]}
      >
        <Input data-testid="input-path" placeholder="Front End Path" maxLength="50" />
      </Form.Item>

      <Form.Item
        name="pattern"
        label="URL Regex Pattern"
        rules={[{ required: true, message: 'Please input the URL Regex Pattern' }]}
      >
        <Input data-testid="input-pattern" placeholder="URL Regex Pattern" maxLength="50" />
      </Form.Item>

      <Form.Item
        name="label"
        label="Label Text"
        rules={[{ required: true, message: 'Please input the Label text' }]}
      >
        <Input data-testid="input-label" placeholder="Label Text" maxLength="30" />
      </Form.Item>
      <Form.Item
        wrapperCol={{ offset: 6, span: 14 }}
      >
        <Button data-testid="btn-save" type="primary" htmlType="submit">
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

ModuleForm.defaultProps = {
  module: {},
};

ModuleForm.propTypes = {
  module: PropTypes.shape({
    id: PropTypes.number,
    name: PropTypes.string,
    path: PropTypes.string,
    pattern: PropTypes.string,
    label: PropTypes.string,
  }),
  handleSubmit: PropTypes.func.isRequired,
};

export default ModuleForm;
