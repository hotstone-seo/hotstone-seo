import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { PageHeader } from 'antd';
import { AuditTrailList } from 'components/AuditTrail';

function ViewAuditTrail() {
  const [listAuditTrail, setListAuditTrail] = useState([]);

  return (
    <div>
      <PageHeader
        title="Audit Trail"
        subTitle="View audit logs"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <AuditTrailList
          listAuditTrail={listAuditTrail}
          setListAuditTrail={setListAuditTrail}
        />
      </div>
    </div>
  );
}

ViewAuditTrail.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ViewAuditTrail;
