import React from 'react';
import PropTypes from 'prop-types';
import { Form, Input, Button } from 'antd';
import { PlusOutlined, MinusCircleOutlined } from '@ant-design/icons';

function BreadcrumbListForm(breadcrumbList) {
  const [form] = Form.useForm();

  return (
    <Form
      form={form}
      initialValues={breadcrumbList}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 12 }}
    >
      <Form.List name="listItem">
        {(fields, { add, remove }) => (
          <div>
            {fields.map((field, index) => (
              <>
                <Form.Item
                  name={[field.name, 'name']}
                  fieldKey={[field.fieldKey, 'name']}
                >
                  <Input placeholder={`Page #${index + 1}'s name`} />
                </Form.Item>
                <Form.Item
                  name={[field.name, 'item']}
                  fieldKey={[field.fieldKey, 'item']}
                >
                  <Input placeholder={`URL #${index + 1}`} />
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
              Add List Item
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
