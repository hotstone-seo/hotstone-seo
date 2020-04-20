import React from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, InputNumber, Button,
} from 'antd';
import { PlusOutlined, MinusCircleOutlined } from '@ant-design/icons';

function BreadcrumbListForm(breadcrumbList) {
  const [form] = Form.useForm();

  return (
    <Form
      form={form}
      initialValues={breadcrumbList}
    >
      <Form.List name="listItem">
        {(fields, { add, remove }) => (
          <div>
            {fields.map((field) => (
              <>
                <Form.Item
                  name={[field.name, 'position']}
                  fieldKey={[field.fieldKey, 'position']}
                >
                  <InputNumber />
                </Form.Item>
                <Form.Item
                  name={[field.name, 'name']}
                  fieldKey={[field.fieldKey, 'name']}
                >
                  <Input />
                </Form.Item>
                <Form.Item
                  name={[field.name, 'item']}
                  fieldKey={[field.fieldKey, 'item']}
                >
                  <Input />
                </Form.Item>
                <MinusCircleOutlined
                  onClick={() => { remove(field.name); }}
                />
              </>
            ))}
            <Button
              type="dashed"
              onClick={() => { add(); }}
            >
              <PlusOutlined />
              Add Question
            </Button>
          </div>
        )}
      </Form.List>
    </Form>
  );
}

BreadcrumbListForm.defaultProps = {
  breadcrumbList: {},
};

BreadcrumbListForm.propTypes = {
  breadcrumbList: PropTypes.shape({
    listItem: PropTypes.arrayOf(PropTypes.shape({
      position: PropTypes.number,
      name: PropTypes.string,
      item: PropTypes.string,
    })),
  }),
};

export default BreadcrumbListForm;
