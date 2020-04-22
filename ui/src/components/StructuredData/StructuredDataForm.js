import React from 'react';
import PropTypes from 'prop-types';
import { Card, Tabs } from 'antd';
import { CloseOutlined } from '@ant-design/icons';

import FAQPageForm from './FAQPageForm';
import BreadcrumbListForm from './BreadcrumbListForm';

const { TabPane } = Tabs;

function StructuredDataForm({ structuredData, onCancel }) {
  const { type: structType, id: structID } = structuredData;

  return (
    <Card
      title={structID ? 'Edit Structured Data' : 'Add new Structured Data'}
      bordered={false}
      extra={<CloseOutlined onClick={onCancel} />}
    >
      <Tabs tabPosition="left">
        <TabPane
          tab="FAQ Page"
          key="faq"
          disabled={(structType && structType !== 'faq')}
        >
          <FAQPageForm />
        </TabPane>
        <TabPane
          tab="Breadcrumb List"
          key="breadcrumb"
          disabled={(structType && structType !== 'breadcrumb')}
        >
          <BreadcrumbListForm />
        </TabPane>
      </Tabs>
    </Card>
  );
}

StructuredDataForm.propTypes = {
  structuredData: PropTypes.shape({
    id: PropTypes.number,
    rule_id: PropTypes.number.isRequired,
    type: PropTypes.string,
  }).isRequired,
  onCancel: PropTypes.func.isRequired,
};

export default StructuredDataForm;
