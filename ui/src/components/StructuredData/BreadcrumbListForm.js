import React from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button, message,
} from 'antd';
import { PlusOutlined, MinusCircleOutlined } from '@ant-design/icons';
import { addBreadcrumbList, updateBreadcrumbList } from 'api/structuredData';

function BreadcrumbListForm({ structuredData, afterSubmit }) {
  const [form] = Form.useForm();

  const { id, rule_id, data } = structuredData;

  const onFinish = (values) => {
    const formStruct = { ...values, id, rule_id };
    const submit = id ? updateBreadcrumbList : addBreadcrumbList;
    submit(formStruct)
      .then((response) => {
        afterSubmit(response);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <Form
      form={form}
      initialValues={{ list_item: data.itemListElement }}
      onFinish={onFinish}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 12 }}
    >
      <Form.List name="list_item">
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
  structuredData: {},
};

BreadcrumbListForm.propTypes = {
  structuredData: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number.isRequired,
    data: PropTypes.shape({
      itemListElement: PropTypes.arrayOf(
        PropTypes.shape({
          position: PropTypes.number,
          name: PropTypes.string,
          item: PropTypes.string,
        }),
      ),
    }),
  }),
  afterSubmit: PropTypes.func.isRequired,
};

export default BreadcrumbListForm;
