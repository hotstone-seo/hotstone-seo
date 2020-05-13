import React from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button, message, Divider,
} from 'antd';
import { addLocalBusiness, updateLocalBusiness } from 'api/structuredData';

function LocalBusinessForm({ structuredData, afterSubmit }) {
  const [form] = Form.useForm();
  const { id, rule_id, data } = structuredData;

  const onFinish = (values) => {
    const formStruct = { ...values, id, rule_id };
    const submit = id ? updateLocalBusiness : addLocalBusiness;
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
      initialValues={{ ...data }}
      onFinish={onFinish}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 12 }}
    >
      <Divider orientation="left">Local Business Info</Divider>
      <Form.Item name="name" label="Name">
        <Input />
      </Form.Item>
      <Form.Item name="image" label="Image URL">
        <Input />
      </Form.Item>
      <Form.Item name="url" label="URL">
        <Input />
      </Form.Item>
      <Form.Item name="telephone" label="Phone">
        <Input />
      </Form.Item>
      <Form.Item name="priceRange" label="Price range">
        <Input />
      </Form.Item>
      <Divider orientation="left">Address</Divider>
      <Form.Item name={['address', 'addressCountry']} label="Country">
        <Input />
      </Form.Item>
      <Form.Item name={['address', 'addressRegion']} label="Region">
        <Input />
      </Form.Item>
      <Form.Item name={['address', 'addressLocality']} label="City">
        <Input />
      </Form.Item>
      <Form.Item name={['address', 'streetAddress']} label="Street">
        <Input />
      </Form.Item>
      <Form.Item name={['address', 'postalCode']} label="Zip code">
        <Input />
      </Form.Item>
      <Divider orientation="left">Rating</Divider>
      <Form.Item name={['aggregateRating', 'ratingValue']} label="Value">
        <Input />
      </Form.Item>
      <Form.Item name={['aggregateRating', 'bestRating']} label="Best rating">
        <Input />
      </Form.Item>
      <Form.Item name={['aggregateRating', 'worstRating']} label="Worst rating">
        <Input />
      </Form.Item>
      <Form.Item name={['aggregateRating', 'reviewCount']} label="Review count">
        <Input />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 6, span: 12 }}>
        <Button type="primary" htmlType="submit">
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

LocalBusinessForm.defaultProps = {
  structuredData: {},
};

LocalBusinessForm.propTypes = {
  structuredData: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number.isRequired,
    data: PropTypes.shape({
      name: PropTypes.string,
      image: PropTypes.string,
      url: PropTypes.string,
      telephone: PropTypes.string,
      priceRange: PropTypes.string,
      address: PropTypes.shape({
        addressCountry: PropTypes.string,
        addressRegion: PropTypes.string,
        addreesLocality: PropTypes.string,
        streetAddress: PropTypes.string,
        postalCode: PropTypes.string,
      }),
      aggregateRating: PropTypes.shape({
        ratingValue: PropTypes.string,
        bestRating: PropTypes.string,
        worstRating: PropTypes.string,
        reviewCount: PropTypes.string,
      }),
    }),
  }),
  afterSubmit: PropTypes.func.isRequired,
};

export default LocalBusinessForm;
