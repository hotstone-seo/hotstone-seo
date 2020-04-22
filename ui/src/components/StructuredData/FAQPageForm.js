import React from 'react';
import PropTypes from 'prop-types';
import { Form, Input, Button } from 'antd';
import { PlusOutlined, MinusCircleOutlined } from '@ant-design/icons';

function FAQPageForm({ faqPage }) {
  const [form] = Form.useForm();

  return (
    <Form
      form={form}
      initialValues={faqPage}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 12 }}
    >
      <Form.List name="faqs">
        {(fields, { add, remove }) => (
          <div>
            {fields.map((field, index) => (
              <>
                <Form.Item
                  name={[field.name, 'question']}
                  fieldKey={[field.fieldKey, 'question']}
                >
                  <Input addonBefore="Q:" placeholder={`Question #${index + 1}`} />
                </Form.Item>
                <Form.Item
                  name={[field.name, 'answer']}
                  fieldKey={[field.fieldKey, 'answer']}
                >
                  <Input addonBefore="A:" placeholder="Answer" />
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

FAQPageForm.defaultProps = {
  faqPage: {},
};

FAQPageForm.propTypes = {
  faqPage: PropTypes.shape({
    faqs: PropTypes.arrayOf(PropTypes.shape({
      question: PropTypes.string,
      answer: PropTypes.string,
    })),
  }),
};

export default FAQPageForm;
