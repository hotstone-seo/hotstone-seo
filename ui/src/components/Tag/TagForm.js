import React, { useState, useMemo } from 'react';
import PropTypes from 'prop-types';
import { Select, PageHeader } from 'antd';

import TagAPI from 'api/tag';
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

const APISets = {
  createAPI: {
    meta: TagAPI.addMeta,
    title: TagAPI.addTitle,
    link: TagAPI.addCanonical,
    script: TagAPI.addScript,
  },
  updateAPI: {
    meta: TagAPI.updateMeta,
    title: TagAPI.updateTitle,
    link: TagAPI.updateCanonical,
    script: TagAPI.updateScript,
  },
};

// TODO: adjust onSubmit to use appropriate API function for each type
function TagForm({ tag, afterSubmit, onCancel }) {
  const [currentType, setCurrentType] = useState(tag.type);
  const APISet = tag.id ? APISets.updateAPI : APISets.createAPI;
  const submitTag = useMemo(() => APISet[currentType], [APISet, currentType]);

  const handleSubmit = (formTag) => {
    submitTag(formTag).then(afterSubmit);
  };

  const renderSelectedForm = (type) => {
    switch (type) {
      case 'title':
        return <TitleForm tag={tag} onSubmit={handleSubmit} />;
      case 'meta':
        return <MetaForm tag={tag} onSubmit={handleSubmit} />;
      case 'link':
        return <CanonicalForm tag={tag} onSubmit={handleSubmit} />;
      case 'script':
        return <ScriptForm tag={tag} onSubmit={handleSubmit} />;
      default:
        return null;
    }
  };

  return (
    <>
      <PageHeader
        title={tag.id ? 'Edit Tag' : 'Add new Tag'}
        onBack={onCancel}
        extra={[
          <Select
            key="selectType"
            defaultValue={currentType}
            placeholder="Select a type"
            onChange={(value) => setCurrentType(value)}
            showSearch
            filterOption={(input, option) => (
              option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
            )}
            dropdownMatchSelectWidth={false}
            data-testid="select-type"
          >
            {tagTypes.map(({ label, value }) => (
              <Option key={value} value={value}>{label}</Option>
            ))}
          </Select>,
        ]}
      />

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
  afterSubmit: PropTypes.func.isRequired,
  onCancel: PropTypes.func.isRequired,
};

export default TagForm;
