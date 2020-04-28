import React from 'react';
import PropTypes from 'prop-types';
import { Card, Tabs } from 'antd';
import { CloseOutlined } from '@ant-design/icons';

import FAQPageForm from './FAQPageForm';
import BreadcrumbListForm from './BreadcrumbListForm';

const { TabPane } = Tabs;

function StructuredDataForm({ structuredData, afterSubmit, onCancel }) {
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
          key="FAQPage"
          disabled={(structType && structType !== 'FAQPage')}
        >
          <FAQPageForm structuredData={structuredData} afterSubmit={afterSubmit} />
        </TabPane>
        <TabPane
          tab="Breadcrumb List"
          key="BreadcrumbList"
          disabled={(structType && structType !== 'BreadcrumbList')}
        >
          <BreadcrumbListForm structuredData={structuredData} afterSubmit={afterSubmit} />
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
  afterSubmit: PropTypes.func.isRequired,
  onCancel: PropTypes.func.isRequired,
};

export default StructuredDataForm;
