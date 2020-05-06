import React, { useState, useCallback } from 'react';
import PropTypes from 'prop-types';
import {
  message, Space, Select, Button,
} from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { fetchTags, deleteTag } from 'api/tag';
import useAsync from 'hooks/useAsync';
import { TagForm, TagList } from 'components/Tag';
import locales from 'locales';

const { Option } = Select;

function ManageTags({ ruleID }) {
  const [focusTag, setFocusTag] = useState(null);
  const [locale, setLocale] = useState(locales[0]);
  const partialFetch = useCallback(
    () => fetchTags({ rule_id: ruleID, locale }),
    [ruleID, locale],
  );
  const {
    value: tags, setValue: setTags, pending, error, execute,
  } = useAsync(partialFetch);

  if (error) {
    message.error(error.message);
  }

  const refreshTag = () => {
    execute().then(() => setFocusTag(null));
  };

  const removeTag = ({ id: tagID }) => {
    deleteTag(tagID)
      .then(() => {
        setTags(tags.filter((tag) => tag.id !== tagID));
      })
      .catch((err) => {
        message.error(err.message);
      });
  };

  return (
    focusTag ? (
      <TagForm tag={focusTag} afterSubmit={refreshTag} onCancel={() => setFocusTag(null)} />
    ) : (
      <Space direction="vertical" style={{ width: '100%' }}>
        <div style={{ float: 'right' }}>
          <span>Locale: </span>
          <Select
            defaultValue={locale}
            onChange={(value) => setLocale(value)}
          >
            {locales.map((localeVal) => (
              <Option key={localeVal} value={localeVal}>
                {localeVal}
              </Option>
            ))}
          </Select>
        </div>
        <Button
          type="dashed"
          onClick={() => setFocusTag({ rule_id: ruleID, locale })}
          style={{ width: '100%' }}
        >
          <PlusOutlined />
          Add Tag
        </Button>
        <TagList
          tags={tags}
          loading={pending}
          onEdit={(tag) => setFocusTag(tag)}
          onDelete={removeTag}
        />
      </Space>
    )
  );
}

ManageTags.propTypes = {
  ruleID: PropTypes.number.isRequired,
};

export default ManageTags;
