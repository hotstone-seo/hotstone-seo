import React from 'react';
import { Form, Input, Button } from 'antd';
import { PlusOutlined, MinusCircleOutlined } from '@ant-design/icons';

function FAQPageForm() {
  const [form] = Form.useForm();

  const faqPage = {
    faqs: [
      { question: 'What\'s your name?', answer: 'nobody' },
      { question: 'How long?', answer: 'in a minute' },
    ],
  };

  return (
    <Form
      form={form}
      initialValues={faqPage}
    >
      <Form.List name="faqs">
        {(fields, { add, remove }) => (
          <div>
            {fields.map((field) => (
              <>
                <Form.Item
                  name={[field.name, 'question']}
                  fieldKey={[field.fieldKey, 'question']}
                >
                  <Input addonBefore="Q:" />
                </Form.Item>
                <Form.Item
                  name={[field.name, 'answer']}
                  fieldKey={[field.fieldKey, 'answer']}
                >
                  <Input addonBefore="A:" />
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

export default FAQPageForm;
