import React from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button, message, Divider,
} from 'antd';
import { PlusOutlined, DeleteTwoTone } from '@ant-design/icons';
import { addFAQPage, updateFAQPage } from 'api/structuredData';

function FAQPageForm({ structuredData, afterSubmit }) {
  const [form] = Form.useForm();

  const { id, rule_id, data } = structuredData;
  const faqs = data && data.mainEntity
    ? data.mainEntity.map((item) => (
      { question: item.name, answer: item.acceptedAnswer.text }
    )) : [];

  const onFinish = (values) => {
    const formStruct = { ...values, id, rule_id };
    const submit = id ? updateFAQPage : addFAQPage;
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
      initialValues={{ faqs }}
      onFinish={onFinish}
      wrapperCol={{ span: 16 }}
    >
      <Form.List name="faqs">
        {(fields, { add, remove }) => (
          <div>
            {fields.map((field, index) => (
              <>
                <Divider orientation="left">{`Question #${index + 1}`}</Divider>
                <Form.Item>
                  <Form.Item
                    name={[field.name, 'question']}
                    fieldKey={[field.fieldKey, 'question']}
                    noStyle
                  >
                    <Input placeholder="Question" style={{ width: '95%' }} />
                  </Form.Item>
                  <DeleteTwoTone
                    twoToneColor="red"
                    onClick={() => { remove(field.name); }}
                    style={{ float: 'right', height: '100%' }}
                  />
                </Form.Item>
                <Form.Item
                  name={[field.name, 'answer']}
                  fieldKey={[field.fieldKey, 'answer']}
                >
                  <Input.TextArea placeholder="Answer" style={{ width: '95%' }} />
                </Form.Item>
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
      <Form.Item wrapperCol={{ span: 16 }}>
        <Button data-testid="btn-save" type="primary" htmlType="submit" style={{ float: 'right' }}>
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

FAQPageForm.defaultProps = {
  structuredData: {},
};

FAQPageForm.propTypes = {
  structuredData: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number.isRequired,
    data: PropTypes.shape({
      mainEntity: PropTypes.arrayOf(
        PropTypes.shape({
          name: PropTypes.string,
          acceptedAnswer: PropTypes.shape({
            text: PropTypes.string,
          }),
        }),
      ),
    }),
  }),
  afterSubmit: PropTypes.func.isRequired,
};

export default FAQPageForm;
