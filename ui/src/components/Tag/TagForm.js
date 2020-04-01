import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { Select } from 'antd';

import TitleForm from './TitleForm';
import MetaForm from './MetaForm';
import CanonicalForm from './CanonicalForm';
import ScriptForm from './ScriptForm';

const { Option } = Select;

const tagTypes = [
  { label: 'Title', value: 'title' },
  { label: 'Meta', value: 'meta' },
  { label: 'Canonical', value: 'link' },
  { label: 'Script', value: 'script' },
];

function TagForm({ tag, onSubmit }) {
  const [currentType, setCurrentType] = useState(tag.type);

  const renderSelectedForm = (type) => {
    switch (type) {
      case 'title':
        return <TitleForm tag={tag} />;
      case 'meta':
        return <MetaForm tag={tag} />;
      case 'link':
        return <CanonicalForm tag={tag} />;
      case 'script':
        return <ScriptForm tag={tag} />;
      default:
        return null;
    }
  };

  return (
    <>
      <Select
        data-testid="select-type"
        onChange={(value) => setCurrentType(value)}
        showSearch
        filterOption={(input, option) => (
          option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
        )}
      >
        {tagTypes.map(({ label, value }) => (
          <Option key={value} value={value}>{label}</Option>
        ))}
      </Select>

      {renderSelectedForm(currentType)}
    </>
  );
}

TagForm.propTypes = {
  tag: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number.isRequired,
    type: PropTypes.string,
  }).isRequired,
  onSubmit: PropTypes.func.isRequired,
};

export default TagForm;
