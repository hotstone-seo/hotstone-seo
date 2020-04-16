import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { message, Space, Select, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { fetchTags, deleteTag } from 'api/tag';
import { TagForm, TagList } from 'components/Tag';
import locales from 'locales';

const { Option } = Select;

function ManageTags({ ruleID }) {
  const [focusTag, setFocusTag] = useState(null);
  const [tags, setTags] = useState([]);
  const [locale, setLocale] = useState(locales[0]);

  useEffect(() => {
    fetchTags({ rule_id: ruleID, locale })
      .then((newTags) => {
        setTags(newTags);
      })
      .catch((error) => {
        message.error(error.message);
      });
  }, [ruleID, locale]);

  const refreshTag = () => {
    fetchTags({ rule_id: ruleID, locale })
      .then((newTags) => {
        setTags(newTags);
        setFocusTag(null);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const removeTag = ({ id: tagID }) => {
    deleteTag(tagID)
      .then(() => {
        setTags(tags.filter((tag) => tag.id !== tagID));
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div style={{ background: '#fff', padding: 24 }}>
      {focusTag ? (
        <TagForm tag={focusTag} afterSubmit={refreshTag} onCancel={() => setFocusTag(null)} />
      ) : (
        <Space direction="vertical" style={{ width: '100%' }}>
          <Select
            defaultValue={locale}
            onChange={(value) => setLocale(value)}
            style={{ float: 'right' }}
          >
            {locales.map((localeVal) => (
              <Option key={localeVal} value={localeVal}>
                {localeVal}
              </Option>
            ))}
          </Select>
          <Button
            type="dashed"
            onClick={() => setFocusTag({ rule_id: ruleID, locale })}
            style={{ width: '100%' }}
          >
            <PlusOutlined />
            Add Tag
          </Button>
          <TagList tags={tags} onEdit={(tag) => setFocusTag(tag)} onDelete={removeTag} />
        </Space>
      )}
    </div>
  );
}

ManageTags.propTypes = {
  ruleID: PropTypes.number.isRequired,
};

export default ManageTags;
