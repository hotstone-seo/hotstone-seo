import React, { useState } from 'react';
import { Select } from 'antd';
import TitleForm from './TitleForm';
import MetaForm from './MetaForm';

const { Option } = Select;

const tagTypes = ['title', 'meta', 'canonical', 'script'];

const capitalize = (item) => {
  return item.charAt(0).toUpperCase() + item.slice(1);
}

function TagForm({ tag }) {
  const [currentType, setCurrentType] = useState(tag ? tag.type : tagTypes[0])
  return (
    <div>
      <Select defaultValue={currentType} onChange={(value) => setCurrentType(value)}>
        {tagTypes.map(tagType => (
          <Option value={tagType}>{capitalize(tagType)}</Option>
        ))}
      </Select>
      {
        {
          title: <TitleForm tag={tag} />,
          meta: <MetaForm tag={tag} />,
          canonical: null,
          script: null,
        }[currentType]
      }
    </div>
  );
}

export default TagForm;
