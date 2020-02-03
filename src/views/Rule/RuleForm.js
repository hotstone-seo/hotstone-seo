import React, { useState, useEffect } from 'react';
import { Form, Input, Select, Button } from 'antd';
import { fetchDatasources } from '../../api/datasource';

const { Option } = Select;

const formLayout = { labelCol: { span: 8 }, wrapperCol: { span: 10 } };

const tailLayout = { wrapperCol: { offset: 8, span: 16 } };

function RuleForm() {
  const [form] = Form.useForm();

  const [dataSources, setDataSources] = useState([]);

  useEffect(() => {
    let _isCancelled = false;
    fetchDatasources()
      .then((dataSources) => {
        if (!_isCancelled) {
          setDataSources(dataSources);
        }
      });

    return () => {
      _isCancelled = true;
    };
  });

  return (
    <Form {...formLayout} form={form}>
      <Form.Item
        name="name"
        label="Name"
        rules={[{ required: true, message: 'Please input the name of your Rule' }]}
      >
        <Input placeholder="My Rule" />
      </Form.Item>

      <Form.Item
        name="urlPattern"
        label="URL Pattern"
        rules={[{ required: true, message: 'Please input the URL Pattern' }]}
      >
        <Input placeholder="/my/rule/pattern" />
      </Form.Item>

      <Form.Item
        name="dataSource"
        label="Data Source"
      >
        <Select allowClear>
          {dataSources.map(({ id, name }) => (
            <Option value={id}>{name}</Option>
          ))}
        </Select>
      </Form.Item>
    </Form>
  );
}

export default RuleForm;
