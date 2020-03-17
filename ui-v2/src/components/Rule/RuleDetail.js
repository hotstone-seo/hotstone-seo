import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { Descriptions, message } from 'antd';
import moment from 'moment';
import { getDataSource } from 'api/datasource';

const formatDate = (dateString) => moment(dateString).format('DD MMM YYYY HH:mm:ss');

function RuleDetail({ rule }) {
  const {
    name,
    url_pattern: urlPattern,
    data_source_id: dataSourceID,
    created_at: createdAt,
    updated_at: updatedAt,
  } = rule;
  const [dataSource, setDataSource] = useState(null);

  useEffect(() => {
    if (dataSourceID) {
      getDataSource(dataSourceID)
        .then((retrievedValue) => {
          setDataSource(retrievedValue);
        })
        .catch((error) => {
          message.error(error.message);
        });
    }
  }, [dataSourceID]);

  return (
    <Descriptions>
      <Descriptions.Item key="name" label="Name">{name}</Descriptions.Item>
      <Descriptions.Item key="urlPattern" label="URL Pattern">{urlPattern}</Descriptions.Item>
      {dataSource && (
        <Descriptions.Item data-testid="lbl-data-source" key="dataSource" label="Data Source">
          <Link to={`/datasources/${dataSource.id}`}>
            {dataSource.name}
          </Link>
        </Descriptions.Item>
      )}
      <Descriptions.Item key="createdAt" label="Created at">{formatDate(createdAt)}</Descriptions.Item>
      <Descriptions.Item key="updatedAt" label="Updated at">{formatDate(updatedAt)}</Descriptions.Item>
    </Descriptions>
  );
}

RuleDetail.propTypes = {
  rule: PropTypes.shape({
    name: PropTypes.string,
    url_pattern: PropTypes.string,
    data_source_id: PropTypes.number,
    created_at: PropTypes.string,
    updated_at: PropTypes.string,
  }).isRequired,
};

export default RuleDetail;
