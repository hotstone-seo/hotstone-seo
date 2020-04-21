import React, { useState, useMemo } from 'react';
import PropTypes from 'prop-types';
import { Card, Tabs } from 'antd';
import { CloseOutlined } from '@ant-design/icons';

import TitleForm from './TitleForm';
import MetaForm from './MetaForm';
import CanonicalForm from './CanonicalForm';
import ScriptForm from './ScriptForm';

const { TabPane } = Tabs;

function TagForm({ tag, afterSubmit, onCancel }) {
  const { type: tagType, id: tagID } = tag;

  return (
    <Card
      title={tagID ? 'Edit Tag' : 'Add new Tag'}
      bordered={false}
      extra={<CloseOutlined onClick={onCancel} />}
    >
      <Tabs tabPosition="left">
        <TabPane
          tab="Title"
          key="title"
          disabled={(tagType && tagType !== 'title')}
        >
          <TitleForm tag={tag} afterSubmit={afterSubmit} />
        </TabPane>
        <TabPane
          tab="Meta"
          key="meta"
          disabled={(tagType && tagType !== 'meta')}
        >
          <MetaForm tag={tag} afterSubmit={afterSubmit} />
        </TabPane>
        <TabPane
          tab="Canonical"
          key="link"
          disabled={(tagType && tagType !== 'link')}
        >
          <CanonicalForm tag={tag} afterSubmit={afterSubmit} />
        </TabPane>
        <TabPane
          tab="Script"
          key="script"
          disabled={(tagType && tagType !== 'script')}
        >
          <ScriptForm tag={tag} afterSubmit={afterSubmit} />
        </TabPane>
      </Tabs>
    </Card>
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
