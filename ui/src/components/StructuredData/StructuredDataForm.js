import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { Select } from 'antd';

import FAQPageForm from './FAQPageForm';
import BreadcrumbListForm from './BreadcrumbListForm';

const { Option } = Select;

const structTypes = [
  { label: 'FAQ Page', value: 'faqpage' },
  { label: 'Breadcrumb List', value: 'breadcrumblist' },
];

function StructuredDataForm({ structuredData }) {
  const { type: structType, id: structID } = structuredData;
  const [currentType, setCurrentType] = useState(structType);

  const renderSelectedForm = (type) => {
    switch (type) {
      case 'faqpage':
        return <FAQPageForm />;
      case 'breadcrumblist':
        return <BreadcrumbListForm />;
      default:
        return null;
    }
  };

  return (
    <>
      <Select
        defaultValue={currentType}
        onChange={(value) => setCurrentType(value)}
        placeholder="Select a type"
        disabled={structID}
      >
        {structTypes.map(({ label, value }) => (
          <Option key={value} value={value}>{label}</Option>
        ))}
      </Select>
      {renderSelectedForm(currentType)}
    </>
  );
}

StructuredDataForm.propTypes = {
  structuredData: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number.isRequired,
    type: PropTypes.string,
  }).isRequired,
};

export default StructuredDataForm;
